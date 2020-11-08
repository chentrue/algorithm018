package main

//根据一棵树的前序遍历与中序遍历构造二叉树。
//
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
	}


func buildTree(preorder []int, inorder []int) *TreeNode{
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 中顺序列找根结点
	var root int
	for k, v := range  inorder {
		if v == preorder[0] {
			root = k
			break
		}
	}
	// 左右子树归类
	// pre_left, pre_right := preorder[1: root+1], preorder[root+1:]
	// in_left, in_right := inorder[0:root], inorder[root+1:]
	return &TreeNode{
		Val: preorder[0],
		Left: buildTree(preorder[1: root+1], inorder[0: root]),
		Right: buildTree(preorder[root+1:],inorder[root+1:]),
	}
}

