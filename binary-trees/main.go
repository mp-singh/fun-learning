package main

import (
	"fmt"
	"github.com/phf/go-queue/queue"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

//size - Returns size of tree
func size(node *Node) int {
	if node == nil {
		return 0
	}
	return size(node.right) + size(node.left) + 1
}

//maxDepth - Returns max depth of tree
func maxDepth(node *Node) int {
	if node == nil {
		return 0
	}
	leftMax := maxDepth(node.left)
	rightMax := maxDepth(node.right)
	if leftMax > rightMax {
		return leftMax + 1
	} else {
		return rightMax + 1
	}
}

//preorderTraversal - Returns an array with nodes in preorder
func preorderTraversal(node *Node) []int {
	result := make([]int, size(node))
	index := 0
	preorderTraversalHelper(node, result, &index)
	return result
}

//preorderTraversalHelper - Preorder traversal helper
func preorderTraversalHelper(node *Node, result []int, index *int) {
	if node == nil {
		return
	}
	result[*index] = node.data // n
	*index++
	preorderTraversalHelper(node.left, result, index)  // l
	preorderTraversalHelper(node.right, result, index) // r
}

//inorderTraversal - Returns an array with nodes in inorder
func inorderTraversal(node *Node) []int {
	result := make([]int, size(node))
	index := 0
	inorderTraversalHelper(node, result, &index)
	return result
}

//inorderTraversalHelper - Inorder traversal helper
func inorderTraversalHelper(node *Node, result []int, index *int) {
	if node == nil {
		return
	}

	inorderTraversalHelper(node.left, result, index) // l
	// n
	result[*index] = node.data
	*index++
	inorderTraversalHelper(node.right, result, index) // r
}

//postorderTraversal - Returns an array with nodes in postorder
func postorderTraversal(node *Node) []int {
	result := make([]int, size(node))
	index := 0
	postorderTraversalHelper(node, result, &index)
	return result
}

//postorderTraversalHelper - Postorder traversal helper
func postorderTraversalHelper(node *Node, result []int, index *int) {
	if node == nil {
		return
	}

	postorderTraversalHelper(node.left, result, index)  // l
	postorderTraversalHelper(node.right, result, index) // r
	// n
	result[*index] = node.data
	*index++
}

//mirror - Return a tree such that the roles of the left and right pointers are swapped at every node
func mirror(node *Node) {
	if node == nil {
		return
	}

	mirror(node.left)
	mirror(node.right)

	//swap left and right node
	node.right, node.left = node.left, node.right
}

//find - Returns true or false depending on if the val is found in the tree
func find(node *Node, val int) bool {
	if node == nil {
		return false
	}
	if node.data == val {
		return true
	}
	return find(node.left, val) || find(node.right, val)
}

//searchBST - Return node in a tree if exists
func searchBST(node *Node, val int) *Node {
	if node == nil {
		return nil
	}
	if val == node.data {
		return node
	} else if val > node.data {
		return searchBST(node.right, val)
	} else if val < node.data {
		return searchBST(node.left, val)
	}
	fmt.Printf("unable to find [%d] in the tree.", val)
	return nil
}

//searchBSTLoop - Return node in a tree if exists using a loop
func searchBSTLoop(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	for node != nil {
		if val == node.data {
			return node
		} else if val < node.data {
			node = node.left
		} else if val > node.data {
			node = node.right
		}
	}
	return nil
}

//hasPathSum - return true or false if there exists a path such that root-to-leaf nodes add up to the sum
func hasPathSum(node *Node, sum int) bool {
	if node == nil {
		return sum == 0
	}

	return hasPathSum(node.left, sum-node.data) || hasPathSum(node.right, sum-node.data)
}

//printPaths - print out all of its root-to-leaf paths using a helper method
func printPaths(node *Node) {
	result := make([]int, 100)
	pathLength := 0
	printPathsHelper(node, result, &pathLength)
	fmt.Println(result)
}

//printPathsHelper - printPaths helper
func printPathsHelper(node *Node, array []int, pathLength *int) {
	if node == nil {
		return
	}
	array[*pathLength] = node.data
	*pathLength++
	if node.left == nil && node.right == nil {
		for i := 0; i < *pathLength; i++ {
			fmt.Printf("%d ", array[i])
		}
		fmt.Println("")
	} else {
		printPathsHelper(node.left, array, pathLength)
		printPathsHelper(node.right, array, pathLength)
	}
}

//insert - inserts a new node in the given tree
func insert(root *Node, node *Node) bool {
	if root.left == nil && node.data < root.data {
		root.left = node
		return true
	} else if root.right == nil && node.data > root.data {
		root.right = node
		return true
	}

	if node.data < root.data {
		return insert(root.left, node)
	} else if node.data > root.data {
		return insert(root.right, node)
	}

	//fmt.Printf("The node [%d] already exists in tree\n", node.data)
	return false
}

//findMax - Returns the maximum node in the tree
func findMax(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return findMax(node.right)
}

//findMaxLoop - Returns the maximum node in the tree using a loop
func findMaxLoop(node *Node) *Node {
	if node == nil {
		return nil
	}

	for node.right != nil {
		node = node.right
	}

	return node
}

//findMin - Return the minimum value in the tree
func findMin(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.left == nil {
		return node
	}

	return findMin(node.left)
}

//findMinLoop - Return the minimum value in the tree using a loop
func findMinLoop(node *Node) *Node {
	if node == nil {
		return nil
	}

	for node.left != nil {
		node = node.left
	}

	return node
}

func printLevelOrder(node *Node) {
	for i := 1; i <= maxDepth(node); i++ {
		printGivenLevel(node, i)
	}
}

func printGivenLevel(node *Node, level int) {
	if node == nil {
		return
	}
	if level == 1 {
		fmt.Printf("%d ", node.data)
	} else {
		printGivenLevel(node.left, level-1)
		printGivenLevel(node.right, level-1)
	}
}

func printLevelOrderWithQueue(node *Node) {
	q := queue.New()
	temp := node

	for temp != nil {
		fmt.Printf("%d ", temp.data)
		if temp.left != nil {
			q.PushBack(temp.left)
		}
		if temp.right != nil {
			q.PushBack(temp.right)
		}
		if q.Len() > 0 {
			temp = q.PopFront().(*Node)
			continue
		}
		temp = nil
	}
}
