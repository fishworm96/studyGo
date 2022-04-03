package main

import (
	"fmt"
	"sort"
)

func main() {
	// a := [5]int{1, 2, 3, 4, 5}
	// s := a[1:3]
	// fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	// 完整切片表达式
	// a[low : high : max]
	//  a :=[5]int{1, 2, 3, 4, 5}
	//  t := a[1:3:5]
	//  fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))

	// 使用make()函数构造切片
	// make([]T, size, cap)
	// a := make([]int, 2, 10)
	// fmt.Println(a)
	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	// 判断切片是否为空
	// 要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。

	// 切片不能直接比较
	// 切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil

	// 切片的赋值拷贝
	// s1 := make([]int, 3) //[0 0 0]
	// s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	// s2[0] = 100
	// fmt.Println(s1) //[100 0 0]
	// fmt.Println(s2) //[100 0 0]

	// 切片遍历
	// s := []int{1, 3, 5}

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(i, s[i])
	// }

	// for index, value := range s {
	// 	fmt.Println(index, value)
	// }

	// append()方法为切片添加元素
	// var s []int
	// s = append(s, 1)        // [1]
	// s = append(s, 2, 3, 4)  // [1 2 3 4]
	// s2 := []int{5, 6, 7}  
	// s = append(s, s2...)    // [1 2 3 4 5 6 7]

	// 使用copy()函数复制切片
	// copy()复制切片
	// a := []int{1, 2, 3, 4, 5}
	// c := make([]int, 5, 5)
	// copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	// fmt.Println(a) //[1 2 3 4 5]
	// fmt.Println(c) //[1 2 3 4 5]
	// c[0] = 1000
	// fmt.Println(a) //[1 2 3 4 5]
	// fmt.Println(c) //[1000 2 3 4 5]

	// 从切片中删除元素
		// 从切片中删除元素
		// a := []int{30, 31, 32, 33, 34, 35, 36, 37}
		// // 要删除索引为2的元素
		// a = append(a[:2], a[3:]...)
		// fmt.Println(a) //[30 31 33 34 35 36 37]

		// var a = make([]string, 5, 10)
		// for i := 0; i < 10; i++ {
		// 	a = append(a, fmt.Sprintf("%v", i))
		// }
		// fmt.Println(a)

		var a = [...]int{3, 7, 8, 9, 1}
		sort.Ints(a[:])
		fmt.Println(a)
}