use std::collections::HashMap;

fn main() {
    println!("Hello, world!");
}

mod trie_node {
    use std::collections::HashMap;

    pub struct TrieNode {
        pub children: HashMap<char, TrieNode>,
    }

    impl TrieNode {
        pub fn new() -> Self {
            TrieNode {
                children: HashMap::new(),
            }
        }
    }
}

mod trie_class {
    use crate::trie_node::TrieNode;

    pub struct Trie {
        root: TrieNode,
    }

    impl Trie {
        fn new() -> Self {
            Trie {
                root: TrieNode::new(),
            }
        }
    }
}
