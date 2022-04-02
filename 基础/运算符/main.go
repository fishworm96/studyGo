package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "math"
)

// func main() {
// 	var s, sep string
// 	for i:= 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func main() {
// 	var a int = 10
// 	fmt.Printf("%d \n", a)
// 	fmt.Printf("%b \n", a)

// 	var b int = 077
// 	fmt.Printf("%o \n", b)

// 	var c int = 0xff
// 	fmt.Printf("%x \n", c)
// 	fmt.Printf("%x \n", c)
// }

// func main() {
// 	fmt.Printf("%f\n", math.Pi)
// 	fmt.Printf("%.2f\n", math.Pi)
// }

// func traversalString() {
// 	s := "hello沙河"
// 	for i := 0; i < len(s); i++ { //byte
// 		fmt.Printf("%v(%c) ", s[i], s[i])
// 	}
// 	fmt.Println()
// 	for _, r := range s { //rune
// 		fmt.Printf("%v(%c) ", r, r)
// 	}
// 	fmt.Println()
// }

// func changeString() {
// 	s1 := "big"
// 	byteS1 := []byte(s1)
// 	byteS1[0] = 'p'
// 	fmt.Println(string(byteS1))

// 	s2 := "hello白萝卜"
// 	runeS2 := []rune(s2)
// 	runeS2[0] = '红'
// 	fmt.Println(string(runeS2))
// }

// func sqrtDemo() {
// 	var a, b = 3, 4
// 	var c int
// 	c = int(math.Sqrt(float64(a*a + b*b)))
// 	fmt.Println(c)
// }

func main() {
	// sqrtDemo()
	// changeString()
	// traversalString()
	// s := [5]int{1, 2, 3, 1, 2}

	fmt.Print(1^1^2^3^4^3^4)
}
