// https://github.com/dghubble/trie
// Trie (префиксное дерево) — это структура данных, которая помогает эффективно хранить
//  и искать строки. В Go она полезна, когда нужно быстро проверять префиксы, делать
// автодополнение или поддерживать поиск с подстановочными символами.
// Когда стоит применять Trie
// • Когда у вас большой словарь или список слов, требующих быстрого поиска
// • Для функций startsWith(prefix), когда нужно быстро проверить, есть ли слова
// с заданным префиксом
// • Когда нужны частичные совпадения с wildcard-символами — например, поиск по шаблонам с точками или звездочками
// • Для реализации игр со словами, таких как Boggle, Scrabble или Wordle
// • В автодополнении и DNS-lookup
// • Для любых задач, где важна быстрая проверка строк по префиксу

package trie

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool        // Флаг конца слова
	value    interface{} // Опционально: Значение, связанное со словом
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[char]
	}

	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root

	for _, char := range word {
		if child, exists := node.children[char]; exists {
			node = child
		} else {
			return false
		}
	}

	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) []string {
	node := t.root

	// Находим узел префикса
	for _, char := range prefix {
		if child, exists := node.children[char]; exists {
			node = child
		} else {
			return []string{} // Префикс не найден
		}
	}

	// Собираем все слова с этим префиксом
	var results []string
	t.dfs(node, prefix, &results)
	return results
}

func (t *Trie) dfs(node *TrieNode, current string, results *[]string) {
	if node.isEnd {
		*results = append(*results, current)
	}

	for char, child := range node.children {
		t.dfs(child, current+string(char), results)
	}
}
