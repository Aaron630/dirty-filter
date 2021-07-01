package dirtyFilter

import (
	"fmt"
	"testing"
)

func TestAddDirtyWord(t *testing.T) {
	g := &trieTree{root: &node{child: make(map[string]*node)}}
	g.AddDirtyWord([]string{"你妈逼的", "你妈", "狗日"})
	text := "你文明用语你&* 妈, 逼的你这个狗 日的，怎么这么傻啊。我也是服了，我日,这些话我都说不出口"
	g.SearchDirtyWord(text)
}

func ExampleAddDirtyWord() {
	g := &trieTree{root: &node{child: make(map[string]*node)}}
	g.AddDirtyWord([]string{"你大爷", "你大妈"})
	fmt.Println("sssssssssss")
	fmt.Println(g.root.child["你"])
}
