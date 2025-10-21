use std::rc::{Rc};
use std::cell::RefCell;
use std::collections::HashMap;

type VertexRef<T> = Rc<RefCell<Vertex<T>>>;

#[derive(Debug)]
pub struct Vertex<T> {
    pub value: T,
    pub adjacent_vertices:  Vec<VertexRef<T>>,
}

impl<T> Vertex<T> {
    // Return a shared, mutable vertex
    pub fn new(value: T) -> VertexRef<T> {
        Rc::new(RefCell::new(Vertex{
            value,
            adjacent_vertices: Vec::new(), 
        }))
    }

    // Add an existing vertex as a neighbor
    pub fn add_adjacent_vertex( this: &VertexRef<T>, vertex: &VertexRef<T>) {
        // the . operator in Rust automatically dereferences through smart pointers
        // (via the Deref trait
        // https://doc.rust-lang.org/std/ops/trait.Deref.html
        //(*this).borrow_mut().adjacent_vertices.push(Rc::clone(vertex));
        this.borrow_mut().adjacent_vertices.push(Rc::clone(vertex));
    }
}

impl<T: std::cmp::PartialEq + Clone> Vertex<T> {
    pub fn dfs_traverse(
        vertex: &VertexRef<T>, 
        visited_vertices: &mut HashMap<*const RefCell<Vertex<T>>, bool>
    ) {
        // Use raw pointer as key because Rc does not implement Hash/Eq
        let ptr = Rc::as_ptr(vertex);
        if visited_vertices.get(&ptr).is_some() {
            return; // already visited
        }

        visited_vertices.insert(ptr, true);

        // Iterate over adjacent vertices
        for adj in &vertex.borrow().adjacent_vertices {
            Self::dfs_traverse(adj, visited_vertices);
        }
    }


    pub fn dfs(
        vertex: &VertexRef<T>, 
        search_value: T,
        visited_vertices: &mut HashMap<*const RefCell<Vertex<T>>, bool>
    ) -> Option<VertexRef<T>> {
        if vertex.borrow().value == search_value {
            return Some(Rc::clone(vertex));
        }

        let ptr = Rc::as_ptr(vertex);
        if visited_vertices.contains_key(&ptr) {
            return None;
        }

        visited_vertices.insert(ptr, true);

        for adj in &vertex.borrow().adjacent_vertices {
            if let Some(found) = Self::dfs(adj, search_value.clone(), visited_vertices) {
                return Some(found);
            }
        }

        None
    }
}

