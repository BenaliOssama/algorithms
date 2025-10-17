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

    pub fn autocomplete(&self, prefix: &str) -> Option<Vec<String>> {
        let current_node = self.search(prefix)?;
        
        let mut word = String::from(prefix);
        let mut words: Vec<String> = vec![];

        self.collect_all_words(Some(current_node), &mut word, &mut words);

        Some(words)
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
    
    pub fn collect_all_words(
        &self,
        node: Option<&TrieNode>,
        word: &mut String,
        words: &mut Vec<String>,
    ) {
        let current_node = node.unwrap_or(&self.root);

        for (key, child_node) in &current_node.children {
            word.push(*key);

            if child_node.is_end_of_word {
                words.push(word.clone());
            }

           self.collect_all_words(Some(child_node), word, words);

           word.pop(); // one step back, loop goes to the next key {backtrack}
        }
    }

}




