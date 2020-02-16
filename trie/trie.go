package trie

type node struct {
	next     map[rune]*node
	terminal bool
}

func newNode() *node {
	return &node{
		next: make(map[rune]*node),
	}
}

func newTerminalNode() *node {
	return &node{
		terminal: true,
		next:     make(map[rune]*node),
	}
}

// Trie is a trie for runes
type Trie struct {
	root *node
}

// NewTrie is the correct way to init a Trie
func NewTrie() Trie {
	return Trie{
		root: newNode(),
	}
}

// Load adds a []rune to the Trie
func (t *Trie) Load(emoji []rune) {
	if len(emoji) == 1 {
		t.root.next[emoji[0]] = newTerminalNode()
		return
	}
	current := t.root
	for i, r := range emoji {
		if i == len(emoji)-1 {
			// last element
			if current.next[r] == nil {
				current.next[r] = newTerminalNode()
			} else {
				// this may never happen
				current.next[r].terminal = true
			}
		} else {
			if current.next[r] == nil {
				current.next[r] = newNode()
			}
			current = current.next[r]
		}
	}
}

// SearchCallback is a function called when a match is found
// return false to keep searching
type SearchCallback func(hit []rune, at int) (done bool)

// Search searches s for any matches in the trie, returns the first n matches
func (t *Trie) Search(s string, n int) [][]rune {
	var found [][]rune
	t.SearchF(s, func(hit []rune, _ int) bool {
		found = append(found, hit)
		if n > 0 && len(found) >= n {
			return true
		}
		return false
	})
	return found
}

// SearchF searches s for trie matches and calls cb when it hits one,
// it will continue until cb returns true
func (t *Trie) SearchF(s string, cb SearchCallback) {
	input := []rune(s)
	var done bool
	for i := 0; i < len(input) && !done; i++ {
		r := input[i]
		if t.root.next[r] != nil {
			breadcrumbs := []rune{}
			var checkpoint []rune
			current := t.root
			for iter := i; iter < len(input); iter++ {
				// get the current rune
				newRune := input[iter]
				next := current.next[newRune]
				if next == nil {
					break
				}
				breadcrumbs = append(breadcrumbs, newRune)
				if next.terminal {
					checkpoint = breadcrumbs
				}
				current = next
			}
			if len(checkpoint) > 0 {
				done = cb(checkpoint, i)
				// advance our progress pass the emoji
				i += len(checkpoint) - 1
			}
		}
	}
}
