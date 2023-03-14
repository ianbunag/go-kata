package implement_trie

type Trie struct {
	nextCharacters [26]*Trie
	isWord         bool
}

func Constructor() Trie {
	return Trie{}
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func (trie *Trie) Insert(word string) {
	if word == "" {
		trie.isWord = true
		return
	}

	character, nextWord := trie.split(word)

	if !trie.hasNextCharacter(character) {
		trie.setNextCharacter(character)
	}

	trie.getNextCharacter(character).Insert(nextWord)
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func (trie *Trie) Search(word string) bool {
	if word == "" {
		return trie.isWord
	}

	character, nextWord := trie.split(word)

	if trie.hasNextCharacter(character) {
		return trie.getNextCharacter(character).Search(nextWord)
	}

	return false
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func (trie *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}

	character, nextWord := trie.split(prefix)

	if trie.hasNextCharacter(character) {
		return trie.getNextCharacter(character).StartsWith(nextWord)
	}

	return false
}

func (trie *Trie) hasNextCharacter(character byte) bool {
	return trie.getNextCharacter(character) != nil
}

func (trie *Trie) getNextCharacter(character byte) *Trie {
	return trie.nextCharacters[character-'a']
}

func (trie *Trie) setNextCharacter(character byte) {
	trie.nextCharacters[character-'a'] = &Trie{}
}

func (trie Trie) split(word string) (byte, string) {
	return word[0], word[1:]
}
