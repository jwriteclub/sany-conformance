// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package tla2sany.parser;

import java.util.Set;
import java.util.TreeSet;

/**
 * Events of a parser patched by scripts/grammar-trace/patch.py. J and P are the
 * two operands of a LOOKAHEAD decision, E is a syntax-error arm's
 * jj_consume_token(-1). A decision with no predicate only has J.
 */
public final class T {
  public static final Set<String> EVENTS = new TreeSet<>();

  private T() {}

  public static boolean j(final int line, final boolean v) {
    EVENTS.add("D:" + line + ":J" + (v ? "1" : "0"));
    return v;
  }

  public static boolean p(final int line, final boolean v) {
    EVENTS.add("D:" + line + ":P" + (v ? "1" : "0"));
    return v;
  }

  /** OpenStart's two sites, per token kind: did the scan match, did the predicate hold. */
  public static boolean o(final int line, final int kind, final boolean v) {
    EVENTS.add("D:" + line + ":J" + (v ? "1" : "0"));
    if (v) {
      EVENTS.add("O:" + line + ":" + TLAplusParserConstants.tokenImage[kind] + ":scan");
    }
    return v;
  }

  public static boolean q(final int line, final int kind, final boolean v) {
    EVENTS.add("D:" + line + ":P" + (v ? "1" : "0"));
    EVENTS.add("O:" + line + ":" + TLAplusParserConstants.tokenImage[kind] + ":above=" + v);
    return v;
  }

  public static void e(final int line) {
    EVENTS.add("E:" + line);
  }
}
