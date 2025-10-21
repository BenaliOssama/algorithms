use graphs::Vertex;
use std::collections::HashMap;

fn main() {
    // Create vertices
    let a = Vertex::new("A");
    let b = Vertex::new("B");
    let c = Vertex::new("C");
    let d = Vertex::new("D");

    // Connect vertices (A->B, A->C, B->D, C->D, D->A to form a cycle)
    Vertex::add_adjacent_vertex(&a, &b);
    Vertex::add_adjacent_vertex(&a, &c);
    Vertex::add_adjacent_vertex(&b, &d);
    Vertex::add_adjacent_vertex(&c, &d);
    Vertex::add_adjacent_vertex(&d, &a); // cycle

    // DFS traversal
    let mut visited = HashMap::new();
    Vertex::dfs_traverse(&a, &mut visited);
    println!("{visited:?}");
}
