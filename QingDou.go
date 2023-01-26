package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
*现有 n 个宽度为 1 的柱子，给出 n 个非负整数依次表示柱子的高度，
此时均匀从上空向下撒青豆，计算按此排列的柱子能接住多少青豆。（不考虑边角堆积）
输入：height = [5,0,2,1,4,0,1,0,3]
输出：17
解析：上面是由数组 [5,0,2,1,4,0,1,0,3] 表示的柱子高度，在这种情况下，可以接 17 个单位的青豆。
*/
var channel = make(chan int, 1)

func main() {
	fmt.Println("欢迎运行青豆小游戏！")
	fmt.Println("现有 n 个宽度为 1 的柱子，给出 n 个非负整数依次表示柱子的高" +
		"度，\n此时均匀从上空向下撒青豆，计算按此排列的柱子能接住多少青豆。（不考虑边角堆积）\n输入：height = [5,0" +
		",2,1,4,0,1,0,3]\n输出：17\n解析：上面是由数组 [5,0,2,1,4,0,1,0,3] 表示的柱子高度，在这种情况下，可以接 17 个单位的青豆")
	for true {
		arr := make([]int, 0)
		fmt.Println("请选择难度(1：简单；2：困难；3：噩梦；else：exit)")
		level := 0
		fmt.Scanln(&level)
		rand.Seed(time.Now().Unix())
		switch level {
		case 1:
			for i := 0; i < 5; i++ {
				arr = append(arr, rand.Intn(10))
			}
		case 2:
			for i := 0; i < 10; i++ {
				arr = append(arr, rand.Intn(100))
			}
		case 3:
			for i := 0; i < 20; i++ {
				arr = append(arr, rand.Intn(1000))
			}
		default:
			os.Exit(404)
		}
		go JieDou(arr)
		fmt.Printf("柱子数组为:%v\n", arr)
		fmt.Println("请输入你的答案：")
		var key int
		userkey := -1
		key = <-channel
		fmt.Scanln(&userkey)
		if userkey == key {
			fmt.Println("恭喜回答正确")
		} else {
			fmt.Println("回答错误")
			fmt.Printf("正确答案为：%d\n", key)
		}
	}

}
func JieDou(arr []int) (ans int) {
	len := len(arr)
	left, right := 0, len-1
	maxLeft, maxRight := arr[left], arr[right]
	for {
		if maxLeft < maxRight {
			ans += maxLeft - arr[left]
			left++
			if left >= right {
				break
			}
			maxLeft = max(maxLeft, arr[left])

		} else {
			ans += maxRight - arr[right]
			right--
			if left >= right {
				break
			}
			maxRight = max(maxRight, arr[right])
		}
	}
	channel <- ans
	return ans
}
func max(x int, y int) int {
	if x <= y {
		return y
	}
	return x
}
