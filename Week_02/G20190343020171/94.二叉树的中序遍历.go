package G20190343020171
/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (69.13%)
 * Likes:    341
 * Dislikes: 0
 * Total Accepted:    84.9K
 * Total Submissions: 122.7K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 * 
 * 示例:
 * 
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * 输出: [1,3,2]
 * 
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 * 
 */

// @lc code=start

// TreeNode efinition for a binary tree node.
type TreeNode struct {
    Val int
     Left *TreeNode
     Right *TreeNode
}
 
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	return append(append(left, root.Val),right...)
}
// @lc code=end

