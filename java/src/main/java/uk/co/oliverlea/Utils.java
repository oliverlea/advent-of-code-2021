package uk.co.oliverlea;

import java.io.IOException;
import java.net.URISyntaxException;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.List;

public class Utils {

    private ClassLoader classLoader;

    public static List<String> readLines(String filename) throws IOException, URISyntaxException {
        URL fileUrl = Utils.class.getClassLoader().getResource(filename);
        Path filePath = Paths.get(fileUrl.toURI());
        return Files.readAllLines(filePath);
    }

    public record Pair<T, V>(T first, V second) {}
}
