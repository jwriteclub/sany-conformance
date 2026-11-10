#!/bin/sh
# Copyright 2026 Christopher O'Connell
# SPDX-License-Identifier: Apache-2.0
#
# Run prudence on the cases in this repository, compare its answers with SANY's, and
# write a report.
#
#   scripts/compare.sh [--set NAME] [--limit N] [--ids FILE] [--stage S] [--jobs N] [--out DIR]
#                      [--java --jar tla2tools.jar [--javabin java]] [--deep-cuts]
#                      [--timeout DUR] [--info] [--no-3d] [--pull]
#
#   --set NAME     one set, e.g. tlaplus-syntax-corpus (default: all sets)
#   --limit N      the first N cases of each set, in manifest order (default: all)
#   --ids FILE     only the case ids listed in FILE, one per line
#   --stage S      prudence stage: semantic or syntax (default: full, semantic and level)
#   --jobs N       cases in parallel (default 1)
#   --out DIR      where report.md and results.jsonl go (default compare-out)
#   --java --jar J run stock SANY with jar J instead of reading the recorded output
#   --javabin JAVA the java binary for --java (default: java on PATH)
#   --deep-cuts    include deep-cuts/ as well as sets/
#   --timeout DUR  per-case timeout for each side, a Go duration (default 300s)
#   --info         keep prudence's non-error diagnostics (the incomplete reason) in results.jsonl
#   --no-3d        turn normalisation rule 3d off (the published numbers use it)
#   --pull         git pull --ff-only first, to get the newest corpus
#
# Needs Go, and a clean checkout of github.com/jwriteclub/prudence at the commit in
# cmd/compare/prudence.pin (PRUDENCE_DIR, default ../prudence). Nothing is downloaded.
# The scoring rule is Compare in cmd/compare/compare.go; start there. The exit status
# is nonzero only if the harness fails, never because cases differ.
set -eu
cd "$(dirname "$0")/.."
REPO=$(pwd)
PRUDENCE_DIR=${PRUDENCE_DIR:-$REPO/../prudence}
[ -f "$PRUDENCE_DIR/go.mod" ] || { echo "no prudence checkout at $PRUDENCE_DIR (set PRUDENCE_DIR)" >&2; exit 2; }
PRUDENCE_DIR=$(cd "$PRUDENCE_DIR" && pwd)
# cmd/compare builds against one immutable prudence commit (cmd/compare/prudence.pin).
PIN=$(grep -v '^#' "$REPO/cmd/compare/prudence.pin" | tr -d '[:space:]')
HEAD=$(git -C "$PRUDENCE_DIR" rev-parse HEAD)
[ "$HEAD" = "$PIN" ] || { echo "prudence at $PRUDENCE_DIR is $HEAD, cmd/compare/prudence.pin wants $PIN" >&2; exit 2; }
[ -z "$(git -C "$PRUDENCE_DIR" status --porcelain --untracked-files=no)" ] || { echo "prudence at $PRUDENCE_DIR has local changes" >&2; exit 2; }

args=""
out=compare-out
while [ $# -gt 0 ]; do
  case "$1" in
    --pull) git pull --ff-only; shift ;;
    --out) out=$2; shift 2 ;;
    --set) args="$args -set $2"; shift 2 ;;
    --limit) args="$args -limit $2"; shift 2 ;;
    --ids) args="$args -ids $2"; shift 2 ;;
    --stage) args="$args -stage $2"; shift 2 ;;
    --jobs) args="$args -jobs $2"; shift 2 ;;
    --jar) args="$args -jar $2"; shift 2 ;;
    --javabin) args="$args -javabin $2"; shift 2 ;;
    --java) args="$args -java"; shift ;;
    --deep-cuts) args="$args -deep-cuts"; shift ;;
    --timeout) args="$args -timeout $2"; shift 2 ;;
    --info) args="$args -info"; shift ;;
    --no-3d) args="$args -no-3d"; shift ;;
    *) echo "unknown argument $1" >&2; exit 2 ;;
  esac
done

mkdir -p "$out"
out=$(cd "$out" && pwd)
# A workspace joins cmd/compare to the local prudence checkout without editing go.mod.
cat > "$out/go.work" <<W
go 1.23
use $REPO/cmd/compare
use $PRUDENCE_DIR
W
(cd "$REPO/cmd/compare" && GOWORK="$out/go.work" GOPROXY=off go build -o "$out/compare" .)
# shellcheck disable=SC2086
"$out/compare" -repo "$REPO" -out "$out" $args
echo "report: $out/report.md"
sed -n '/^## Summary/,/^Mismatches by kind/p' "$out/report.md"
