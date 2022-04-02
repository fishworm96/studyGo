package main

import (
	"fmt"
)

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func Demo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func Demo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func testSwitch3() {
	switch n := 7; n {
	case 1, 3 ,5 ,7 ,9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func ff() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %d\n", j, i, j * i)
		}
		fmt.Printf("")
	}
}

func main() {
	// ifDemo1()
	// ifDemo2()
	// forDemo()
	// Demo2()
	ff()
}