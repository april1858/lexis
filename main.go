package main

import (
	"fmt"
	"regexp"
)

type Node struct {
	Parent *Node
	Lnode  *Node
	Rnode  *Node
	Value  string
}

func NewNode(v string) *Node {
	n := Node{
		Value: v,
	}
	return &n
}

func main() {
	root := NewNode("")

	current := root

	slStr := []string{"a", "and", "b", "or"}

	re, err := regexp.Compile(`and|not|or|xor`)
	if err != nil {
		fmt.Println("error re - ", err)
	}

	for i := 0; i < len(slStr); i++ {
		if !re.MatchString(slStr[i]) {

			if current.Lnode == nil {
				current.Lnode = NewNode(slStr[i])
				current.Lnode.Parent = current
			} else if current.Rnode == nil {
				current.Rnode = NewNode(slStr[i])
				current.Rnode.Parent = current
			}
		} else {
			if current.Value == "" {
				current.Value = slStr[i]
			} else {
				current.Parent = NewNode(slStr[i])
				current.Parent.Lnode = current
				current = current.Parent
				root = current
			}
		}
	}

	fmt.Println(root)
	fmt.Println(root.Lnode)
	fmt.Println(root.Rnode)

}
