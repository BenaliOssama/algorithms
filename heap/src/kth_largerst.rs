// this is a solution to this leetcode problem:  https://leetcode.com/problems/kth-largest-element-in-a-stream/

// this example usage of min-heap


struct KthLargest {
    k: usize,
    heap: Heap<i32>, // min-heap
}

impl KthLargest {
    fn new(k: i32, nums: Vec<i32>) -> Self {
        let mut heap = Heap::new();
        let k = k as usize;

        for num in nums {
            if heap.data.len() < k {
                heap.insert(num);
            } else if num > *heap.root_node().unwrap() {
                heap.delete();   
                heap.insert(num); 
            }
        }

        Self { k, heap }
    }

    fn add(&mut self, val: i32) -> i32 {
        if self.heap.data.len() < self.k {
            self.heap.insert(val);
        } else if val > *self.heap.root_node().unwrap() {
            self.heap.delete();
            self.heap.insert(val);
        }

        *self.heap.root_node().unwrap()
    }
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * let obj = KthLargest::new(k, nums);
 * let ret_1: i32 = obj.add(val);
 */

#[derive(Debug, Clone)]
pub struct Heap<T> {
    pub data: Vec<T>,
}

impl<T: Ord> Heap<T> {
    pub fn new() -> Self {
        Self { data: Vec::new() }
    }

    pub fn root_node(&self) -> Option<&T> {
        self.data.get(0)
    }

    pub fn left_child_index(&self, index: usize) -> Option<usize> {
        let left = 2 * index + 1;
        if left < self.data.len() { Some(left) } else { None }
    }

    pub fn right_child_index(&self, index: usize) -> Option<usize> {
        let right = 2 * index + 2;
        if right < self.data.len() { Some(right) } else { None }
    }

    pub fn parent_index(&self, index: usize) -> Option<usize> {
        if index == 0 { None } else { Some((index - 1) / 2) }
    }

    // **Insert for min-heap**
    pub fn insert(&mut self, value: T) {
        self.data.push(value);
        let mut index = self.data.len() - 1;

        while let Some(parent) = self.parent_index(index) {
            if self.data[index] < self.data[parent] { 
                self.data.swap(index, parent);
                index = parent;
            } else {
                break;
            }
        }
    }

    pub fn delete(&mut self) {
        if self.data.is_empty() { return; }

        let last = self.data.pop().unwrap();
        if !self.data.is_empty() {
            self.data[0] = last;
            self.trickle_down(0);
        }
    }

    fn trickle_down(&mut self, mut index: usize) {
        while let Some(smaller_child_index) = self.smaller_child_index(index) {
            if self.data[smaller_child_index] < self.data[index] { 
                self.data.swap(index, smaller_child_index);
                index = smaller_child_index;
            } else {
                break;
            }
        }
    }

    fn smaller_child_index(&self, index: usize) -> Option<usize> {
        let left = self.left_child_index(index);
        let right = self.right_child_index(index);

        match (left, right) {
            (Some(l), Some(r)) => {
                if self.data[l] < self.data[r] { Some(l) } else { Some(r) }
            }
            (Some(l), None) => Some(l),
            _ => None,
        }
    }
}

