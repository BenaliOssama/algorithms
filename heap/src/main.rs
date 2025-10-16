use heap::Heap;

fn main() {
    let mut heap = Heap::new();
    heap.insert(10);
    heap.insert(20);
    heap.insert(5);
    println!("{:?}", heap); // Should show max-heap structure

    heap.delete();
    println!("{:?}", heap); // Root (20) removed
}

