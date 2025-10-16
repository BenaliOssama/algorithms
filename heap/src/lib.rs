#[derive(Debug)]
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
        if left < self.data.len() {
            Some(left)
        } else {
            None
        }
    }

    pub fn right_child_index(&self, index: usize) -> Option<usize> {
        let right = 2 * index + 2;
        if right < self.data.len() {
            Some(right)
        } else {
            None
        }
    }

    pub fn parent_index(&self, index: usize) -> Option<usize> {
        if index == 0 {
            None
        } else {
            Some((index - 1) / 2)
        }
    }

    pub fn insert(&mut self, value: T) {
        self.data.push(value);
        let mut index = self.data.len() - 1;

        while let Some(parent) = self.parent_index(index) {
            if self.data[index] > self.data[parent] {
                self.data.swap(index, parent);
                index = parent;
            } else {
                break;
            }
        }
    }

    pub fn delete(&mut self) {
        if self.data.is_empty() {
            return;
        }

        // Replace root with the last element
        let last = self.data.pop().unwrap();
        if !self.data.is_empty() {
            self.data[0] = last;
            self.trickle_down(0);
        }
    }

    fn trickle_down(&mut self, mut index: usize) {
        while let Some(larger_child_index) = self.larger_child_index(index) {
            if self.data[larger_child_index] > self.data[index] {
                self.data.swap(index, larger_child_index);
                index = larger_child_index;
            } else {
                break;
            }
        }
    }

    fn larger_child_index(&self, index: usize) -> Option<usize> {
        let left = self.left_child_index(index);
        let right = self.right_child_index(index);

        match (left, right) {
            (Some(l), Some(r)) => {
                if self.data[l] > self.data[r] {
                    Some(l)
                } else {
                    Some(r)
                }
            }
            (Some(l), None) => Some(l),
            _ => None,
        }
    }
}

