package uk.co.oliverlea;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.List;

public class Day2 {

    private record Movement(String direction, int amount) {}

    public static int problem1(List<Movement> movements) {
        int x = 0;
        int y = 0;
        for (Movement move : movements) {
            switch (move.direction) {
                case "up": {
                    y -= move.amount;
                    break;
                }
                case "down": {
                    y += move.amount;
                    break;
                }
                case "forward": {
                    x += move.amount;
                    break;
                }
                case "backward": {
                    x -= move.amount;
                    break;
                }
                default: throw new IllegalStateException("Unknown direction: " + move.direction);
            }
        }
        return x * y;
    }

    public static int problem2(List<Movement> movements) {
        int x = 0;
        int y = 0;
        int aim = 0;
        for (Movement move : movements) {
            switch (move.direction) {
                case "up": {
                    aim -= move.amount;
                    break;
                }
                case "down": {
                    aim += move.amount;
                    break;
                }
                case "forward": {
                    y += move.amount;
                    x += aim * move.amount;
                    break;
                }
                case "backward": {
                    y -= move.amount;
                    x -= aim * move.amount;
                    break;
                }
                default: throw new IllegalStateException("Unknown direction: " + move.direction);
            }
        }
        return x * y;
    }

    public static void main(String... args) throws IOException, URISyntaxException {
        List<Movement> movements = Utils.readLines("input2").stream()
                        .map(l -> {
                            String[] lexed = l.split(" ");
                            return new Movement(lexed[0], Integer.parseInt(lexed[1]));
                        })
                        .toList();

        System.out.println(problem1(movements));
        System.out.println(problem2(movements));
    }
}
