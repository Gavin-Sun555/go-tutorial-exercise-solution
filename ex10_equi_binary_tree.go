package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk_Helper(t *tree.Tree, ch chan int){
	if (t != nil){
		Walk_Helper(t.Left, ch)
		ch <- t.Value
		Walk_Helper(t.Right, ch)
	}
	return
}

func Walk(t *tree.Tree, ch chan int){
	Walk_Helper(t,ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch)
	go Walk(t2, ch2)
	for{
		a,ok1 := <-ch
		b,ok2 := <-ch2
		if(!ok1 && !ok2){
            //if both tree traverse to the end
			return true
		}else if(!ok1 || !ok2){
            //if either tree traverse to the end but other not(meaning two trees are different)
			return false
		}else if(a != b){
            //if either number is different
			return false
		}
	}
	
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
