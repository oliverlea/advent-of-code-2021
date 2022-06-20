package uk.co.oliverlea;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.IntStream;

public class Day05 {

    private static void addLine(List<List<Integer>> grid, Utils.Pair<Coord, Coord> lineSpec, boolean skipDiagonal) {
        if (lineSpec.first().x != lineSpec.second().x && lineSpec.first().y != lineSpec.second().y) {
            // Diagonal
            if (skipDiagonal) {
                return;
            }

            if (lineSpec.first().x > lineSpec.second().x) {
                // Flip around so we're always working left to right
                lineSpec = new Utils.Pair<>(lineSpec.second(), lineSpec.first());
            }
            boolean down = lineSpec.first().y < lineSpec.second().y;
            if (down) {
                for (int curX = lineSpec.first().x; curX <= lineSpec.second().x; curX++) {
                    int curY = lineSpec.first().y + (curX - lineSpec.first().x);
                    grid.get(curY).set(curX, grid.get(curY).get(curX) + 1);
                }
            } else {
                for (int curX = lineSpec.first().x; curX <= lineSpec.second().x; curX++) {
                    int curY = lineSpec.first().y - (curX - lineSpec.first().x);
                    grid.get(curY).set(curX, grid.get(curY).get(curX) + 1);
                }
            }
        } else if (lineSpec.first().x != lineSpec.second().x) {
            // Horizontal
            int curY = lineSpec.first().y;
            IntStream.rangeClosed(
                    Math.min(lineSpec.first().x, lineSpec.second().x),
                    Math.max(lineSpec.first().x, lineSpec.second().x)
            ).forEach(curX -> grid.get(curY).set(curX, grid.get(curY).get(curX) + 1));
        } else {
            // Vertical
            int curX = lineSpec.first().x;
            IntStream.rangeClosed(
                    Math.min(lineSpec.first().y, lineSpec.second().y),
                    Math.max(lineSpec.first().y, lineSpec.second().y)
            ).forEach(curY -> grid.get(curY).set(curX, grid.get(curY).get(curX) + 1));
        }
    }

    private static List<List<Integer>> emptyGrid(int width, int height) {
        List<List<Integer>> grid = new ArrayList<>();
        for (int y = 0; y < height; y++) {
            List<Integer> row = new ArrayList<>(width);
            for (int x = 0; x < width; x++) {
                row.add(0);
            }
            grid.add(row);
        }
        return grid;
    }

    private static List<List<Integer>> constructGrid(List<Utils.Pair<Coord, Coord>> lineSpecs) {
        int largestX = lineSpecs.stream().mapToInt(ls -> Math.max(ls.first().x, ls.second().x)).max().getAsInt();
        int largestY = lineSpecs.stream().mapToInt(ls -> Math.max(ls.first().y, ls.second().y)).max().getAsInt();
        return emptyGrid(largestX + 1, largestY + 1);
    }

    public static long problem1(List<List<Integer>> grid, List<Utils.Pair<Coord, Coord>> lineSpecs) {
        lineSpecs.forEach(ls -> addLine(grid, ls, true));
        return grid.stream().mapToLong(row -> row.stream().filter(x -> x > 1).count()).sum();
    }

    public static long problem2(List<List<Integer>> grid, List<Utils.Pair<Coord, Coord>> lineSpecs) {
        lineSpecs.forEach(ls -> addLine(grid, ls, false));
        return grid.stream().mapToLong(row -> row.stream().filter(x -> x > 1).count()).sum();
    }

    public static void main(String... args) throws IOException, URISyntaxException {
        List<String> lines = Utils.readLines("input05");

        List<Utils.Pair<Coord, Coord>> lineSpecs = new ArrayList<>(lines.size());
        for (String line : lines) {
            String[] startAndEnd = line.split(" -> ");
            lineSpecs.add(new Utils.Pair<>(Coord.parse(startAndEnd[0]), Coord.parse(startAndEnd[1])));
        }

        List<List<Integer>> grid = constructGrid(lineSpecs);
        System.out.println(problem1(grid, lineSpecs));

        grid = constructGrid(lineSpecs);
        System.out.println(problem2(grid, lineSpecs));

    }


    private record Coord(int x, int y) {

        public static Coord parse(String input) {
            String[] splitInput = input.split(",");
            return new Coord(Integer.parseInt(splitInput[0]), Integer.parseInt(splitInput[1]));
        }
    }
}
