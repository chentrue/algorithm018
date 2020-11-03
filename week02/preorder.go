func preorder(root *Node) []int {
	var data []int
	if root != nil{
		data = append(data, root.Val)
		if len(root.Children) > 0{
			for _,node := range root.Children{
				data = append(data, preorder(node)...)
			}
		}
	}
	return data    
}
