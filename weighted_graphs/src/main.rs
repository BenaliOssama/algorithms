use std::collections::HashMap;
use weighted_graphs::*;

fn main() {
    let mut a = City::new("A");
    let mut b = City::new("B");
    let mut c = City::new("C");

    a.connect(&mut b, 5);
    b.connect(&mut c, 2);
    a.connect(&mut c, 9);

    let cities = HashMap::from([
        (a.name.clone(), a),
        (b.name.clone(), b),
        (c.name.clone(), c),
    ]);

    let (dist, path) = dijkstra_shortest_path(&cities, "A", "C").unwrap();
    println!("Shortest distance: {}", dist);
    println!("Path: {:?}", path);
}
