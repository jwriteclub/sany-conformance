// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

//! tlxcheck: parse and check one TLA+ file with tlx-sema, nothing else.
//!
//! usage: tlxcheck [--parse] [-I dir]... File.tla
//!
//! `--parse` runs only tlx's front end (`tlx_front::parse::parse_modules`, the
//! call behind `tlx parse`) on the root file: no import resolution, no names.
//!
//! It calls `tlx_sema::Program::load`, the same call as tlx's own `check`
//! example, so no model checking happens. One JSON object on stdout:
//! `{"ok":bool,"crash":string|null,"errors":[{"file","line","col","message"}]}`.
//! Exit status 0 when the call returned, 2 on bad usage, 101 on a panic
//! (still reported as JSON). No tlx source is changed.

use std::path::PathBuf;

use tlx_sema::Program;

const STACK: usize = 512 << 20; // the stack size tlx itself uses to evaluate

fn esc(s: &str) -> String {
    let mut o = String::with_capacity(s.len() + 2);
    o.push('"');
    for c in s.chars() {
        match c {
            '"' => o.push_str("\\\""),
            '\\' => o.push_str("\\\\"),
            '\n' => o.push_str("\\n"),
            '\r' => o.push_str("\\r"),
            '\t' => o.push_str("\\t"),
            c if (c as u32) < 0x20 => o.push_str(&format!("\\u{:04x}", c as u32)),
            c => o.push(c),
        }
    }
    o.push('"');
    o
}

fn run_parse(file: PathBuf) -> String {
    let src = match std::fs::read_to_string(&file) {
        Ok(s) => s,
        Err(e) => return format!("{{\"ok\":false,\"crash\":null,\"errors\":[{{\"file\":{},\"line\":0,\"col\":0,\"message\":{}}}]}}", esc(&file.display().to_string()), esc(&e.to_string())),
    };
    match tlx_front::parse::parse_modules(&src, tlx_front::ast::FileId(0)) {
        Ok(_) => "{\"ok\":true,\"crash\":null,\"errors\":[]}".to_string(),
        Err(errors) => {
            let items: Vec<String> = errors
                .iter()
                .map(|e| {
                    format!(
                        "{{\"file\":{},\"line\":{},\"col\":{},\"message\":{}}}",
                        esc(&file.display().to_string()),
                        e.span.line,
                        e.span.col,
                        esc(&e.message)
                    )
                })
                .collect();
            format!("{{\"ok\":false,\"crash\":null,\"errors\":[{}]}}", items.join(","))
        }
    }
}

fn run(file: PathBuf, includes: Vec<PathBuf>) -> String {
    match Program::load(&file, &includes) {
        Ok(_) => "{\"ok\":true,\"crash\":null,\"errors\":[]}".to_string(),
        Err(errors) => {
            let items: Vec<String> = errors
                .iter()
                .map(|e| {
                    let f = e.file.as_ref().map(|p| p.display().to_string()).unwrap_or_default();
                    format!(
                        "{{\"file\":{},\"line\":{},\"col\":{},\"message\":{}}}",
                        esc(&f),
                        e.span.line,
                        e.span.col,
                        esc(&e.message)
                    )
                })
                .collect();
            format!("{{\"ok\":false,\"crash\":null,\"errors\":[{}]}}", items.join(","))
        }
    }
}

fn main() {
    let mut includes = Vec::new();
    let mut files = Vec::new();
    let mut parse_only = false;
    let mut args = std::env::args().skip(1);
    while let Some(a) = args.next() {
        if a == "--parse" {
            parse_only = true;
        } else if a == "-I" {
            match args.next() {
                Some(d) => includes.push(PathBuf::from(d)),
                None => {
                    eprintln!("-I needs a directory");
                    std::process::exit(2);
                }
            }
        } else {
            files.push(PathBuf::from(a));
        }
    }
    let [file] = files.as_slice() else {
        eprintln!("usage: tlxcheck [--parse] [-I dir]... File.tla");
        std::process::exit(2);
    };
    let file = file.clone();
    std::panic::set_hook(Box::new(|_| {}));
    let handle = std::thread::Builder::new()
        .stack_size(STACK)
        .spawn(move || std::panic::catch_unwind(|| if parse_only { run_parse(file) } else { run(file, includes) }))
        .expect("spawn");
    match handle.join() {
        Ok(Ok(json)) => println!("{json}"),
        Ok(Err(p)) | Err(p) => {
            let msg = p
                .downcast_ref::<String>()
                .cloned()
                .or_else(|| p.downcast_ref::<&str>().map(|s| s.to_string()))
                .unwrap_or_else(|| "panic".to_string());
            println!("{{\"ok\":false,\"crash\":{},\"errors\":[]}}", esc(&msg));
            std::process::exit(101);
        }
    }
}
