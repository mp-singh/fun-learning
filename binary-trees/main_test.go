package main

import (
	"reflect"
	"testing"
)

var tree = &Node{
	data: 6,
	left: &Node{
		data:  4,
		left:  &Node{data: 2},
		right: &Node{data: 5},
	},
	right: &Node{
		data: 7,
		left: nil,
		right: &Node{
			data:  9,
			left:  &Node{data: 8},
			right: &Node{data: 11},
		},
	},
}

func Test_maxDepth(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test max depth of something normal", args: args{tree}, want: 4},
		{name: "empty tree", args: args{}, want: 0},
		{name: "nil tree", args: args{nil}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.node); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_size(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "size of something normal", args: args{tree}, want: 8},
		{name: "empty", args: args{}, want: 0},
		{name: "empty", args: args{nil}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := size(tt.args.node); got != tt.want {
				t.Errorf("size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasPathSum(t *testing.T) {
	type args struct {
		node *Node
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "path 31", args: args{tree, 31}, want: false},
		{name: "path 1", args: args{tree, 1}, want: false},
		{name: "path 2", args: args{tree, 2}, want: false},
		{name: "path 8", args: args{tree, 8}, want: false},
		{name: "path 12", args: args{tree, 12}, want: true},
		{name: "path 30", args: args{tree, 30}, want: true},
		{name: "empty tree", args: args{&Node{}, 0}, want: true},
		{name: "nil tree", args: args{nil, 0}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.node, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchBST(t *testing.T) {
	type args struct {
		node *Node
		val  int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "search for 10", args: args{tree, 10}, want: nil},
		{name: "search for 9", args: args{tree, 9}, want: &Node{data: 9, left: &Node{data: 8}, right: &Node{data: 11}}},
		{name: "search for 8", args: args{tree, 8}, want: &Node{data: 8}},
		{name: "search for 20", args: args{tree, 20}, want: nil},
		{name: "search empty", args: args{&Node{}, 100}, want: nil},
		{name: "search nil", args: args{nil, 100}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.node, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchBSTLoop(t *testing.T) {
	type args struct {
		node *Node
		val  int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "search for 10", args: args{tree, 10}, want: nil},
		{name: "search for 9", args: args{tree, 9}, want: &Node{data: 9, left: &Node{data: 8}, right: &Node{data: 11}}},
		{name: "search for 8", args: args{tree, 8}, want: &Node{data: 8}},
		{name: "search for 20", args: args{tree, 20}, want: nil},
		{name: "search empty", args: args{&Node{}, 100}, want: nil},
		{name: "search nil", args: args{nil, 100}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBSTLoop(tt.args.node, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBSTLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMax(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "find max true", args: args{tree}, want: &Node{data: 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMax(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxLoop(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "find max true", args: args{tree}, want: &Node{data: 11}},
		{name: "empty", args: args{&Node{}}, want: &Node{}},
		{name: "nil", args: args{nil}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLoop(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMaxLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "find min true", args: args{tree}, want: &Node{data: 2}},
		{name: "empty", args: args{&Node{}}, want: &Node{}},
		{name: "nil", args: args{nil}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinLoop(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "find min true", args: args{tree}, want: &Node{data: 2}},
		{name: "empty", args: args{&Node{}}, want: &Node{}},
		{name: "nil", args: args{nil}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinLoop(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "inorder", args: args{tree}, want: []int{2, 4, 5, 6, 7, 8, 9, 11}},
		{name: "empty", args: args{}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "postorder", args: args{tree}, want: []int{2, 5, 4, 8, 11, 9, 7, 6}},
		{name: "empty", args: args{}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "preorder", args: args{tree}, want: []int{6, 4, 2, 5, 7, 9, 8, 11}},
		{name: "empty", args: args{}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_find(t *testing.T) {
	type args struct {
		node *Node
		val  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "search for 10", args: args{tree, 10}, want: false},
		{name: "search for 9", args: args{tree, 9}, want: true},
		{name: "search for 8", args: args{tree, 8}, want: true},
		{name: "search for 20", args: args{tree, 20}, want: false},
		{name: "search in empty tree", args: args{&Node{}, 9}, want: false},
		{name: "search in nil tree", args: args{nil, 9}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find(tt.args.node, tt.args.val); got != tt.want {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	type args struct {
		root *Node
		node *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "insert 20", args: args{tree, &Node{data: 20}}, want: true},
		{name: "insert 1", args: args{tree, &Node{data: 1}}, want: true},
		{name: "val already exists", args: args{tree, &Node{data: 2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.root, tt.args.node); got != tt.want {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
	//assert.Equal(t, inorderTraversal(tree), []int{1, 2, 4, 5, 6, 7, 8, 9, 11, 20})
}

// bench
func BenchmarkSearchBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchBSTLoop(tree, 9)
	}
}

func BenchmarkSearchBSTLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchBST(tree, 9)
	}
}

func BenchmarkFindMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMax(tree)
	}
}

func BenchmarkFindMaxLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMaxLoop(tree)
	}
}

func BenchmarkFindMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMin(tree)
	}
}

func BenchmarkFindMinLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMinLoop(tree)
	}
}
