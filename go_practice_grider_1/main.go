package main

func main() {
	newCards := newDeckFromFile("funfile.txt")

	newCards.altShuffle()
	newCards.print()
}
