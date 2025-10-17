use tries::{Trie, TrieNode};

fn main(){
    println!("hello world");
    
    let trie = TrieNode::new();

    println!("{:?}", trie);

    let mut trie = Trie::new();
    println!("{:?}", trie);
    trie.insert("ace");
    trie.insert("act");
    trie.insert("bad");
    trie.insert("bat");
    trie.insert("bake");
    trie.insert("batter");
    trie.insert("cab");
    trie.insert("cat");
    trie.insert("catnap");
    trie.insert("catnip");
    println!("{:?}", trie);
}
