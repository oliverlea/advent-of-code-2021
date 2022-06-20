package uk.co.oliverlea;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.*;
import java.util.stream.Collectors;

public class Day04 {

    private static List<Integer> parseCalledNumbers(List<String> lines) {
        return Arrays.stream(lines.get(0).split(","))
                .mapToInt(Integer::parseInt)
                .boxed().toList();
    }

    private static List<Board> parseBoards(List<String> lines) {
        // Skip first line which has the called numbers
        lines = lines.subList(1, lines.size());

        List<List<List<Integer>>> boardGrids = new ArrayList<>();
        List<List<Integer>> curBoard = new ArrayList<>();
        for (String line : lines) {
            if (line.isBlank()) {
                if (!curBoard.isEmpty()) {
                    boardGrids.add(curBoard);
                    curBoard = new ArrayList<>();
                }
                continue;
            }
            List<Integer> curLine = Arrays.stream(line.strip().split("\s+"))
                    .mapToInt(Integer::parseInt)
                    .boxed().collect(Collectors.toList());
            curBoard.add(curLine);
        }
        if (!curBoard.isEmpty()) {
            boardGrids.add(curBoard);
        }

        return boardGrids.stream().map(Board::new).toList();
    }

    public static long problem1(List<Integer> calledNumbers, List<Board> boards) {
        return boards.stream()
                .map(b -> b.score(calledNumbers))
                .min(Comparator.comparingInt(BoardResult::turns))
                .get().score();
    }

    public static long problem2(List<Integer> calledNumbers, List<Board> boards) {
        return boards.stream()
                .map(b -> b.score(calledNumbers))
                .max(Comparator.comparing(BoardResult::turns))
                .get().score();
    }

    public static void main(String... args) throws IOException, URISyntaxException {
        List<String> lines = Utils.readLines("input04");
        List<Integer> calledNumbers = parseCalledNumbers(lines);

        List<Board> boards = parseBoards(lines);
        System.out.println(problem1(calledNumbers, boards));
        boards = parseBoards(lines);
        System.out.println(problem2(calledNumbers, boards));
    }

    private record BoardResult(int turns, long score) {};

    private static class Board {

        private final List<List<Integer>> grid;
        private long sumOfUnmarked;

        public Board(List<List<Integer>> grid) {
            this.grid = grid;
            this.sumOfUnmarked = grid.stream()
                    .mapToLong(row -> row.stream().mapToInt(x -> x).sum())
                    .sum();
        }

        private boolean rowSet(int rowIndex) {
            for (Integer rowValue : this.grid.get(rowIndex)) {
                if (rowValue != null) {
                    return false;
                }
            }
            return true;
        }

        private boolean colSet(int colIndex) {
            for (List<Integer> integers : this.grid) {
                if (integers.get(colIndex) != null) {
                    return false;
                }
            }
            return true;
        }

        public BoardResult score(List<Integer> calls) {
            for (int i = 0; i < calls.size(); i++) {
                int call = calls.get(i);

                for (int y = 0; y < this.grid.size(); y++) {
                    for (int x = 0; x < this.grid.get(0).size(); x++) {
                        Integer valueAtLoc = this.grid.get(y).get(x);
                        if (valueAtLoc == null) {
                            continue;
                        }
                        if (valueAtLoc == call) {
                            this.sumOfUnmarked -= valueAtLoc;
                            this.grid.get(y).set(x, null);
                            if (this.rowSet(y) || this.colSet(x)) {
                                return new BoardResult(i, this.sumOfUnmarked * call);
                            }
                        }
                    }
                }
            }
            return new BoardResult(Integer.MAX_VALUE, Long.MAX_VALUE);
        }
    }
}
