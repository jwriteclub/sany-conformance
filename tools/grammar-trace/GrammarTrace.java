// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

import tla2sany.output.SilentSanyOutput;
import tla2sany.parser.T;
import tla2sany.parser.TLAplusParser;

/**
 * Parses each listed file with the patched parser in one JVM and prints, per
 * event: event, number of files, smallest file, its size. With -each it prints
 * every file's events instead. The argument is a file holding one path a line.
 */
public final class GrammarTrace {
  private GrammarTrace() {}

  public static void main(final String[] args) throws Exception {
    final boolean each = args[0].equals("-each");
    final List<String> files = Files.readAllLines(Path.of(args[each ? 1 : 0]));
    final Map<String, Integer> count = new TreeMap<>();
    final Map<String, String> smallest = new TreeMap<>();
    final Map<String, Long> size = new TreeMap<>();
    for (final String f : files) {
      final byte[] b = Files.readAllBytes(Path.of(f));
      T.EVENTS.clear();
      String end = "ok";
      try {
        final TLAplusParser p = new TLAplusParser(new SilentSanyOutput(), b);
        end = p.parse() ? "accept" : "reject";
      } catch (final Throwable t) {
        end = t.getClass().getSimpleName();
      }
      if (each) {
        System.out.println(f + "\t" + end + "\t" + String.join(" ", T.EVENTS));
      }
      for (final String e : T.EVENTS) {
        count.merge(e, 1, Integer::sum);
        if (!size.containsKey(e) || b.length < size.get(e)) {
          size.put(e, (long) b.length);
          smallest.put(e, f);
        }
      }
    }
    if (!each) {
      for (final String e : count.keySet()) {
        System.out.println(e + "\t" + count.get(e) + "\t" + smallest.get(e) + "\t" + size.get(e));
      }
    }
  }
}
