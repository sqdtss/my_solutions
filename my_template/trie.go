package my_template

// for using:
// t := CreateTrie()
// t.Insert("hello")
// t.Insert("helloworld")
// t.Insert("hellots")
// fmt.Println(t.StartsWith("hello"))
// fmt.Println(t.Search("helloa"))

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func CreateTrie() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	for _, c := range word {
		index := c - 'a'
		if this.children[index] == nil {
			this.children[index] = &Trie{}
		}
		this = this.children[index]
	}
	this.isEnd = true
}

func (this *Trie) searchPrefix(prefix string) *Trie {
	for _, c := range prefix {
		index := c - 'a'
		if this.children[index] == nil {
			return nil
		}
		this = this.children[index]
	}
	return this
}

func (this *Trie) Search(word string) bool {
	if res := this.searchPrefix(word); res != nil && res.isEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}
