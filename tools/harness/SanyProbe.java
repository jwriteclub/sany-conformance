// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

import java.io.FileOutputStream;
import java.io.OutputStreamWriter;
import java.io.PrintWriter;
import java.io.Writer;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import tla2sany.drivers.FrontEndException;
import tla2sany.drivers.SANY;
import tla2sany.drivers.SanyExitCode;
import tla2sany.drivers.SanySettings;
import tla2sany.modanalyzer.ParseUnit;
import tla2sany.modanalyzer.SpecObj;
import tla2sany.output.LogLevel;
import tla2sany.output.OutErrSanyOutput;
import tla2sany.output.SanyOutput;
import tla2sany.parser.SyntaxTreeNode;
import tla2sany.semantic.Errors.ErrorDetails;
import tla2sany.st.Location;
import tla2sany.st.TreeNode;
import util.FileUtil;
import util.ToolIO;

/**
 * Runs one SANY front-end pass the way the command line driver does, then
 * writes what the command line does not expose: the standardized error codes
 * with locations, and the syntax trees. Log text goes to stdout unchanged, so
 * stdout is comparable with the plain CLI.
 *
 * Usage: SanyProbe result.json tree.txt Root.tla
 */
public final class SanyProbe {

  private SanyProbe() {}

  public static void main(final String[] args) throws Exception {
    final String resultPath = args[0];
    final String treePath = args[1];
    final String file = args[2];

    final SanyOutput out =
        new OutErrSanyOutput(ToolIO.out, ToolIO.out, LogLevel.INFO, LogLevel.ERROR);
    out.log(LogLevel.INFO, "\n****** SANY2 %s\n", SANY.version);

    final StringBuilder json = new StringBuilder("{");
    final SpecObj spec = new SpecObj(file, null);
    String crash = null;
    int exit = 0;
    if (null == FileUtil.createNamedInputStream(file, spec.getResolver())) {
      crash = "cannot find file " + file;
    } else {
      try {
        final SanyExitCode code = SANY.parse(spec, file, out, SanySettings.defaultSettings());
        exit = code.code();
      } catch (final FrontEndException fe) {
        crash = "FrontEndException: " + fe;
      } catch (final Throwable t) {
        crash = t.getClass().getName() + ": " + t.getMessage();
      }
    }

    field(json, "exit", Integer.toString(exit));
    field(json, "error_level", Integer.toString(spec.getErrorLevel()));
    field(json, "syntax_ok", Boolean.toString(
        spec.getParseErrors().isSuccess()
            && SanyExitCode.SYNTAX_PARSING_FAILURE.code() != spec.getErrorLevel()));
    field(json, "crash", null == crash ? "null" : quote(crash));
    field(json, "parse_messages", messages(spec.getParseErrors().getMessages()));
    field(json, "semantic_messages", messages(spec.getSemanticErrors().getMessages()));
    field(json, "root_module", null == spec.getName() ? "null" : quote(spec.getName()));

    String treeError = null;
    boolean treeWritten = false;
    try {
      treeWritten = writeTrees(spec, treePath);
    } catch (final Throwable t) {
      treeError = t.getClass().getName() + ": " + t.getMessage();
    }
    field(json, "tree_written", Boolean.toString(treeWritten));
    field(json, "tree_error", null == treeError ? "null" : quote(treeError));

    json.append("\"end\":true}\n");

    ToolIO.out.flush();
    try (Writer w = new OutputStreamWriter(new FileOutputStream(resultPath), StandardCharsets.UTF_8)) {
      w.write(json.toString());
    }
    System.exit(0);
  }

  private static void field(final StringBuilder sb, final String name, final String value) {
    sb.append('"').append(name).append("\":").append(value).append(',');
  }

  private static String messages(final List<ErrorDetails> list) {
    final StringBuilder sb = new StringBuilder("[");
    boolean first = true;
    for (final ErrorDetails d : list) {
      if (!first) {
        sb.append(',');
      }
      first = false;
      final Location l = d.getLocation();
      sb.append("{\"code\":").append(d.getCode().getStandardValue());
      sb.append(",\"name\":").append(quote(d.getCode().name()));
      sb.append(",\"severity\":").append(quote(d.getCode().getSeverityLevel().name()));
      sb.append(",\"module\":").append(null == l.source() ? "null" : quote(l.source()));
      sb.append(",\"bl\":").append(l.beginLine());
      sb.append(",\"bc\":").append(l.beginColumn());
      sb.append(",\"el\":").append(l.endLine());
      sb.append(",\"ec\":").append(l.endColumn());
      sb.append(",\"text\":").append(quote(d.getMessage()));
      sb.append('}');
    }
    return sb.append(']').toString();
  }

  /** Dumps the concrete syntax tree of every non-standard module, sorted by module name. */
  private static boolean writeTrees(final SpecObj spec, final String path) throws Exception {
    final List<String> names = new ArrayList<>(spec.parseUnitContext.keySet());
    Collections.sort(names);
    boolean any = false;
    try (PrintWriter w = new PrintWriter(new OutputStreamWriter(new FileOutputStream(path), StandardCharsets.UTF_8))) {
      for (final String name : names) {
        if (spec.getResolver().isStandardModule(name)) {
          continue;
        }
        final ParseUnit unit = spec.parseUnitContext.get(name);
        final TreeNode root = unit.getParseTree();
        if (null == root) {
          continue;
        }
        any = true;
        w.print("module ");
        w.print(name);
        w.print('\n');
        dump(w, root, 0);
      }
    }
    return any;
  }

  private static void dump(final PrintWriter w, final TreeNode node, final int depth) {
    final SyntaxTreeNode n = (SyntaxTreeNode) node;
    final Location l = n.getLocation();
    for (int i = 0; i < depth; i++) {
      w.print(' ');
    }
    w.print(n.getKind());
    w.print(' ');
    w.print(quote(n.getImage()));
    w.print(' ');
    w.print(l.beginLine() + ":" + l.beginColumn() + "-" + l.endLine() + ":" + l.endColumn());
    final String[] pre = n.getPreComments();
    if (null != pre) {
      for (final String c : pre) {
        w.print(" pre=");
        w.print(quote(c));
      }
    }
    w.print('\n');
    final TreeNode[] heirs = n.heirs();
    if (null != heirs) {
      for (final TreeNode h : heirs) {
        dump(w, h, depth + 1);
      }
    }
  }

  private static String quote(final String s) {
    if (null == s) {
      return "null";
    }
    final StringBuilder sb = new StringBuilder(s.length() + 2);
    sb.append('"');
    for (int i = 0; i < s.length(); i++) {
      final char c = s.charAt(i);
      switch (c) {
        case '"':
          sb.append("\\\"");
          break;
        case '\\':
          sb.append("\\\\");
          break;
        case '\n':
          sb.append("\\n");
          break;
        case '\r':
          sb.append("\\r");
          break;
        case '\t':
          sb.append("\\t");
          break;
        default:
          if (c < 0x20 || (c >= 0xD800 && c <= 0xDFFF) || c == 0xFFFE || c == 0xFFFF) {
            sb.append(String.format("\\u%04x", (int) c));
          } else {
            sb.append(c);
          }
      }
    }
    return sb.append('"').toString();
  }
}
