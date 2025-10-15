pub struct Heap<T> {
    pub data: Vec<T>,
}

impl<T> Heap<T> {
    pub fn root_node(&self) -> Option<&T> {
        self.data.get(0)
    }

    pub fn last_node(&self) -> Option<&T> {
        self.data.last()
    }

    pub fn left_child(&self, index: usize) -> Option<&T> {
        let left_index = 2 * index + 1;
        self.data.get(left_index)
    }

    pub fn right_child(&self, index: usize) -> Option<&T> {
        let right_index = 2 * index + 2;
        self.data.get(right_index)
    }

    pub fn parent_index(&self, index: usize) -> Option<&T> {
        let parent_index = (index - 1) / 2 ;
        self.data.get(parent_index)
    }
}

