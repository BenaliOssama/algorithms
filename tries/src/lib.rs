use std::collections::HashMap;

#[derive(Debug)]
pub struct TrieNode {
    children: HashMap<char, TrieNode>,
    is_end_of_word: bool,
}

impl TrieNode {
    pub fn new() -> Self {
        Self {
            children: HashMap::new(),
            is_end_of_word: false,
        }
    }
}


#[derive(Debug)]
pub struct Trie {
    root: TrieNode,
}

impl Trie{
    pub fn new()-> Self{
        Self{
            root : TrieNode::new(),
        }
    }


    pub fn search(&self, word: &str) -> Option<&TrieNode> {
        let mut current_node = &self.root;

        for ch in word.chars() {
            if let Some(node) = current_node.children.get(&ch) {
                current_node = node;
            } else {
                return None;
            }
        }

        Some(current_node)
    }

    pub fn contains(&self, word: &str) -> bool {
        if let Some(node) = self.search(word) {
            node.is_end_of_word
        } else {
            false
        }
    }
    
    pub fn insert(&mut self, word: &str) {
        let mut current_node = &mut self.root;

        for ch in word.chars() {
                current_node = current_node.children
                    .entry(ch).or_insert(TrieNode::new());
            }

            current_node.is_end_of_word = true;
    }
    
    pub fn collect_all_words(&self, node: &TrieNode, word: &str, words: &Vec<String>){


    }
}




