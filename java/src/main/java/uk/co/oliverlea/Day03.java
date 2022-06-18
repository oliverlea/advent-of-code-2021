package uk.co.oliverlea;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.*;
import java.util.function.BiFunction;
import java.util.function.BiPredicate;
import java.util.function.Function;
import java.util.function.Predicate;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class Day03 {

    public static int problem1(List<List<Character>> lines) {
        StringBuilder mostCommon = new StringBuilder();
        for (int i = 0; i < lines.get(0).size(); i++) {
            int finalI = i;
            long trues = lines.stream().map(l -> l.get(finalI)).filter(cur -> cur == '1').count();
            mostCommon.append(trues > lines.size() / 2 ? "1" : "0");
        }
        StringBuilder leastCommon = new StringBuilder();
        for (Character c : mostCommon.toString().toCharArray()) {
            leastCommon.append(c == '0' ? '1' : '0');
        }

        int gamma = Integer.parseInt(mostCommon.toString(), 2);
        int epsilon = Integer.parseInt(leastCommon.toString(), 2);
        return gamma * epsilon;
    }

    private static int reduceMatching(
            List<List<Character>> lines,
            BiFunction<Long, List<List<Character>>, Character> reducer
    ) {
        Set<Integer> remainingMatches = IntStream.range(0, lines.size()).boxed().collect(Collectors.toSet());
        int lineLength = lines.get(0).size();

        for (int i = 0; i < lineLength; i++) {
            List<List<Character>> remainingLines = remainingMatches.stream().map(lines::get).toList();
            int indexCopy = i;

            long oneCount = remainingLines.stream()
                    .map(l -> l.get(indexCopy))
                    .filter(c -> c == '1')
                    .count();
            char mostCommon = reducer.apply(oneCount, remainingLines);

            Set<Integer> keptMatches = new HashSet<>(remainingMatches);
            for (Integer index : remainingMatches) {
                if (lines.get(index).get(i) != mostCommon) {
                    keptMatches.remove(index);
                }
            }
            remainingMatches = keptMatches;

            if (remainingMatches.size() == 1) {
                break;
            }
        }

        if (remainingMatches.size() != 1) {
            throw new IllegalStateException("Expected one result remaining, found: " + remainingMatches.size());
        }
        String binaryString = lines.get(remainingMatches.iterator().next()).stream()
                .map(String::valueOf)
                .collect(Collectors.joining());
        return Integer.valueOf(binaryString, 2);
    }

    public static int problem2(List<List<Character>> lines) {
        int oxygen = reduceMatching(
                lines,
                (oneCount, remaining) -> oneCount >= remaining.size() - oneCount ? '1' : '0'
        );
        int co2 = reduceMatching(
                lines,
                (oneCount, remaining) -> oneCount < remaining.size() - oneCount ? '1' : '0'
        );
        return oxygen * co2;
    }

    public static void main(String... args) throws IOException, URISyntaxException {
        List<List<Character>> bits = Utils.readLines("input03").stream()
                .map(l ->
                        l.chars().mapToObj(c -> (char) c).collect(Collectors.toList())
                )
                .toList();

        System.out.println(problem1(bits));
        System.out.println(problem2(bits));
    }
}
