from typing import Any


class TrieNode:

    def __init__(self) -> None:
        self.children: dict[str, Any] = {}
        self.isEnd = False



class Trie:

    def __init__(self) -> None:
        self.root = TrieNode()


    def insert(self, word: str):
        node = self.root
        for char in word.lower():
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        node.isEnd = True

    
    def show_sugggestions(self, node: TrieNode, word: str, results: list[str]):
        if node.isEnd:
            results.append(word)

        for char, nextNode in node.children.items():
            self.show_sugggestions(nextNode, word+char, results)

        
    def search(self, word: str) -> list[str]:
        node = self.root
        for char in word.lower():
            if char not in node.children:
                return []
            node = node.children[char]

        res: list[str] = []
        self.show_sugggestions(node, word.lower(), res)
        return res
