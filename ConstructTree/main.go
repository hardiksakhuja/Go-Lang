/*
Asking the user to enter inorder and postorder
Printing Level Order traversal as the final output
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	key   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the inorder traversal of Tree  *space separated intergers*:\n ")
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	inorder := convertToIntArr(line)
	fmt.Print("Enter the postorder traversal of Tree  *space separated intergers*:\n ")
	line, err = in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	postorder := convertToIntArr(line)
	rootNode := buildTree(inorder, postorder)

	fmt.Println("The level order Traversal of the constructed binary tree is:")
	levelOrder(rootNode)
}

func buildTree(inorder, postorder []int) *Node {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &Node{
		key: rootVal,
	}
	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}

	root.left = buildTree(inorder[:i], postorder[:i])
	root.right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])

	return root
}

func convertToIntArr(str string) []int {
	arr := make([]int, 0)
	for _, val := range strings.Fields(str) {
		integer, _ := strconv.Atoi(val)
		arr = append(arr, integer)
	}
	return arr
}

func levelOrder(node *Node) {
	h := height(node)
	for i := 1; i <= h; i++ {
		printLevel(node, i)
	}
}

func printLevel(node *Node, level int) {
	if node == nil {
		return
	}
	if level == 1 {
		fmt.Printf("%d  ", node.key)
	} else if level > 1 {
		printLevel(node.left, level-1)
		printLevel(node.right, level-1)
	}
}

func height(node *Node) int {
	if node == nil {
		return 0
	} else {
		//Compute the height of each subtree
		lheight := height(node.left)
		rheight := height(node.right)

		//Use the larger one
		if lheight > rheight {
			return lheight + 1
		} else {
			return rheight + 1
		}

	}
}
