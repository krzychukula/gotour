package main

import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkTree(t, ch)
	close(ch)
}

func walkTree(t *tree.Tree, ch chan int) {

	if t.Left != nil {
		walkTree(t.Left, ch)
	}
	ch <- t.Value

	if t.Right != nil {
		walkTree(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			fmt.Println(v1, v2, ok1, ok2)
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(2), ch)

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("should true: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("should false: ", Same(tree.New(1), tree.New(2)))

	//var input string
	//fmt.Scanln(&input)
}
