use std::collections::HashMap;

#[derive(Debug)]
pub struct TrieNode {
    children: HashMap<char, TrieNode>, 
    is_end_of_word: bool,
}

impl TrieNode {
    pub fn new()-> Self{
        Self{
            children : HashMap::new(),
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
}
