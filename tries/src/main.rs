use tries::Trie;

fn main(){
    let mut trie = Trie::new();
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
    let words = trie.autocomplete("ca");
    println!("completing ca -> {:?}", words);
}
