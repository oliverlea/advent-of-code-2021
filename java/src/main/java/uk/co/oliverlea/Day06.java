package uk.co.oliverlea;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.function.Function;
import java.util.stream.Collectors;

public class Day06 {

    private static final int START_AGE = 8;
    private static final int RESTART_AGE = 6;

    public static long solution1(List<Integer> fishAges, int targetDays) {
        Map<Integer, Long> currentGen = fishAges.stream()
                .collect(Collectors.groupingBy(Function.identity(), Collectors.counting())
        );
        for (int i = 0; i < targetDays; i++) {
            Map<Integer, Long> nextGen = new HashMap<>();
            for (int fishAge = 0; fishAge <= START_AGE; fishAge++) {
                long count = currentGen.getOrDefault(fishAge, 0L);
                if (fishAge == 0) {
                    nextGen.compute(START_AGE, (k, v) -> v == null ? count :  v + count);
                    nextGen.compute(RESTART_AGE, (k, v) -> v == null ? count :  v + count);
                } else {
                    nextGen.compute(fishAge - 1, (k, v) -> v == null ? count :  v + count);
                }
            }
            currentGen = nextGen;
        }
        return currentGen.values().stream().mapToLong(x -> x).sum();
    }

    public static void main(String... args) throws IOException, URISyntaxException {
        List<String> lines = Utils.readLines("input06");
        List<Integer> fishDays = Arrays.stream(lines.get(0).split(","))
                .map(Integer::parseInt).toList();

        System.out.println(solution1(fishDays, 80));
        System.out.println(solution1(fishDays, 256));
    }
}
