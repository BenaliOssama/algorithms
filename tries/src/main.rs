use tries::{Trie, TrieNode};

fn main(){
    println!("hello world");
    
    let trie = TrieNode::new();

    println!("{:?}", trie);

    let trie = Trie::new();

    println!("{:?}", trie);
}
