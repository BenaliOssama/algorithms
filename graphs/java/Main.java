import java.util.*;

class City {
    private String name;
    private Map<City, Integer> routes;

    public City(String name) {
        this.name = name;
        this.routes = new HashMap<>();
    }

    public String getName() {
        return name;
    }

    public Map<City, Integer> getRoutes() {
        return routes;
    }

    public void addRoute(City city, int price) {
        routes.put(city, price);
    }

    @Override
    public String toString() {
        return name;
    }
}

class Dijkstra {
    public static List<String> shortestPath(City start, City destination) {
        Map<String, Integer> cheapestPrices = new HashMap<>();
        Map<String, String> previousCity = new HashMap<>();
        Set<City> visited = new HashSet<>();
        List<City> unvisited = new ArrayList<>();

        cheapestPrices.put(start.getName(), 0);
        City current = start;

        while (current != null) {
            visited.add(current);
            unvisited.remove(current);

            for (Map.Entry<City, Integer> entry : current.getRoutes().entrySet()) {
                City adj = entry.getKey();
                int price = entry.getValue();

                if (!visited.contains(adj) && !unvisited.contains(adj)) {
                    unvisited.add(adj);
                }

                int newPrice = cheapestPrices.getOrDefault(current.getName(), Integer.MAX_VALUE) + price;
                if (newPrice < cheapestPrices.getOrDefault(adj.getName(), Integer.MAX_VALUE)) {
                    cheapestPrices.put(adj.getName(), newPrice);
                    previousCity.put(adj.getName(), current.getName());
                }
            }

            // Pick the unvisited city with the smallest known price
            current = unvisited.stream()
                    .min(Comparator.comparingInt(c -> cheapestPrices.getOrDefault(c.getName(), Integer.MAX_VALUE)))
                    .orElse(null);
        }

        // Reconstruct shortest path
        List<String> path = new ArrayList<>();
        String currentName = destination.getName();

        while (!currentName.equals(start.getName())) {
            path.add(currentName);
            currentName = previousCity.get(currentName);
            if (currentName == null) { // no path found
                return Collections.emptyList();
            }
        }

        path.add(start.getName());
        Collections.reverse(path);
        return path;
    }
}

public class Main {
    public static void main(String[] args) {
        City a = new City("A");
        City b = new City("B");
        City c = new City("C");
        City d = new City("D");

        a.addRoute(b, 1);
        a.addRoute(c, 4);
        b.addRoute(c, 2);
        b.addRoute(d, 5);
        c.addRoute(d, 1);

        List<String> path = Dijkstra.shortestPath(a, d);
        System.out.println("Shortest path from A to D: " + path);
    }
}

