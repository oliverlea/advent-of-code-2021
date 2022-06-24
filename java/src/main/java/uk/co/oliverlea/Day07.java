package uk.co.oliverlea;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class Day07 {

    private static int median(List<Integer> xs) {
        if (xs.isEmpty()) {
            throw new IllegalArgumentException();
        }
        Collections.sort(xs);
        return xs.get(xs.size() / 2);
    }

    private static int mean(List<Integer> xs) {
        if (xs.isEmpty()) {
            throw new IllegalArgumentException();
        }
        float mean = ((float) xs.stream().mapToInt(x -> x).sum()) / (float) xs.size();
        return Math.round(mean);
    }

    private static int sumBetween(int a, int b) {
        return ((b - a + 1) * (a + b)) / 2;
    }

    public static int solution1(List<Integer> crabPositions) {
        int median = median(crabPositions);
        return crabPositions.stream().mapToInt(x -> Math.abs(x - median)).sum();
    }

    public static int solution2(List<Integer> crabPositions) {
        int mean = mean(crabPositions);
        return crabPositions.stream()
                .mapToInt(x -> Math.abs(x - mean))
                .map(x -> sumBetween(0, x))
                .sum();
    }

    public static void main(String... args) throws IOException, URISyntaxException {
        List<String> lines = Utils.readLines("input07");
        List<Integer> crabPositions = Arrays.stream(lines.get(0).split(",")).map(Integer::parseInt).toList();

        System.out.println(solution1(new ArrayList<>(crabPositions)));
        System.out.println(solution2(new ArrayList<>(crabPositions)));
    }
}
