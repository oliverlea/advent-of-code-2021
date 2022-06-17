package uk.co.oliverlea;

import java.io.IOException;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

public class Day1 {

    public static int problem1(List<Integer> input) {
        int increases = 0;
        int prev = Integer.MAX_VALUE;
        for (Integer cur : input) {
            if (cur > prev) {
                increases++;
            }
            prev = cur;
        }
        return increases;
    }

    public static int problem2(List<Integer> input) {
        int prevSum = Integer.MAX_VALUE;
        int increases = 0;
        for (int i = 0; i < input.size() - 2; i++) {
            int curSum = input.get(i) + input.get(i + 1) + input.get(i + 2);
            if (curSum > prevSum) {
                increases++;
            }
            prevSum = curSum;
        }
        return increases;
    }

    public static void main(String[] args) throws IOException, URISyntaxException {
        List<Integer> input = Utils.readLines("input1").stream()
                .map(Integer::valueOf)
                .toList();

        System.out.println(problem1(input));
        System.out.println(problem2(input));
    }

}
