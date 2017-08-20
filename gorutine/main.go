package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkRecur(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkRecur(t.Left, ch)
	ch <- t.Value
	walkRecur(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecur(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	if ok := Same(tree.New(1), tree.New(1)); ok {
		fmt.Println("Одинаковые")
	} else {
		fmt.Println("Разные")
	}
}
