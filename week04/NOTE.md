学习笔记
代码
package main

import "fmt"

func bsearch(nums []int) int{
    right := len(nums)                   
    left := 0
    for left < right{
    	mid := (right + left) / 2             //取中间
    	if  nums[left] < nums[mid]{            // 中间值跟作业比较，找到不规则的区域在左面还是右面 
    		left = mid                         //在不规则区域， 逐渐缩小范围
		}else {
			right = mid
		}
	}
	return right + 1                          // 最后left right变为相当，取right值，那不连续区域即 right和right+1的之间 
}


func main(){
	var nums = []int{5, 6, 7, 8,0, 1, 2, 4}
	fmt.Println(bsearch(nums))
}
