package main

import (
	"fmt"
)

// 方法一
// func main() {
// 	var testArray [3]int
// 	numArray := [3]int{1, 2}
// 	cityArray := [3]string{"北京", "上海", "深圳"}
// 	fmt.Println(testArray)
// 	fmt.Println(numArray)
// 	fmt.Println(cityArray)
// }

// 方法二
// func main() {
// 	var testArray [3]int
// 	var numArray = [...]int{1, 2}
// 	var cityArray = [...]string{"北京", "上海", "深圳"}
// 	fmt.Println(testArray)
// 	fmt.Println(numArray)
// 	fmt.Println(cityArray)
// }

// 方法三
// func main() {
// 	a := [...]int{1: 1, 3: 5}
// 	fmt.Println(a)                  // [0 1 0 5]
// 	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
// }

// 求数组[1, 3, 5, 7, 8]所有元素的和
// func main() {
// 	arr := [...]int{1, 3 ,5 ,7 ,8}
// 	sum := 0
// 	for _, i := range arr {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func main() {
	arr := [...]int{1, 3, 5, 7, 8}
	for _, i := range arr {
		for _, j := range arr {
			if 8 - i == j {
				fmt.Println(i, j)
			}
		}
	}
}
