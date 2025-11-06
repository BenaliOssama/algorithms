use graphs::*;
use std::rc::{Rc};
use std::cell::RefCell;

fn main() {
    // Create vertices
    let a = Rc::new(RefCell::new(Vertex { value: "A", adjacent_vertices: vec![] }));
    let b = Rc::new(RefCell::new(Vertex { value: "B", adjacent_vertices: vec![] }));
    let c = Rc::new(RefCell::new(Vertex { value: "C", adjacent_vertices: vec![] }));
    let d = Rc::new(RefCell::new(Vertex { value: "D", adjacent_vertices: vec![] }));

    // Build connections (graph)
    a.borrow_mut().adjacent_vertices.push(b.clone());
    a.borrow_mut().adjacent_vertices.push(c.clone());
    b.borrow_mut().adjacent_vertices.push(d.clone());
    c.borrow_mut().adjacent_vertices.push(d.clone());

    println!("BFS Traversal:");
    Vertex::bfs_traverse(a.clone());

    println!("\nSearching for vertex 'D':");
    match Vertex::bfs(a.clone(), &"D") {
        Some(found) => println!("Found: {}", found.borrow().value),
        None => println!("Not found"),
    }

    println!("\nSearching for vertex 'X':");
    match Vertex::bfs(a.clone(), &"X") {
        Some(found) => println!("Found: {}", found.borrow().value),
        None => println!("Not found"),
    }
}