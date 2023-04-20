package main

import "strconv"

//import . "nc_tools"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

/**
 *
 * @param root TreeNode类
 * @return int整型
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum int

	nums := getAllPath(root)
	for _, v := range nums {
		num, _ := strconv.Atoi(v)
		sum += num
	}
	// write code here
	return sum
}

func getAllPath(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	str := strconv.Itoa(root.Val)
	var res []string

	if root.Left != nil {
		tmp := getAllPath(root.Left)
		for _, v := range tmp {
			res = append(res, str+v)
		}
	}

	if root.Right != nil {
		tmp := getAllPath(root.Right)
		for _, v := range tmp {
			res = append(res, str+v)
		}
	}

	if len(res) == 0 {
		return []string{str}
	}
	return res
}
