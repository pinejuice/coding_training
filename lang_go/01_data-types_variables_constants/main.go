// 基本的な文法に関する内容を記載
package main

import "fmt"

// グローバル変数の宣言
var word string = "Hello world."

func main() {
	// グローバル変数の出力
	fmt.Println(word)

	// ローカル変数の宣言
	const x string = "Number"
	fmt.Println(x)

	// ローカル定数の宣言
	const y = 100
	fmt.Println(y)

	// 変数と定数の計算とキャスト
	var z1 int
	var z2 float64
	var z3 float64
	z1 = 10 + y
	fmt.Println(z1)
	z2 = 12.5 + y
	fmt.Println(z2)
	z3 = float64(z1) + z2
	fmt.Println(z3)

	// 複数の定数の定義
	const (
		a int    = 1
		b int    = 2
		c string = "Japan"
	)
	fmt.Println(a, b, c) // 1 2 Japan
	const (
		d int = 1
		e
		f
	)
	fmt.Println(d, e, f)

	// iotaを用いた定義
	const (
		Monday    = iota // 0
		Tuesday   = iota // 1
		Wednesday = iota // 2
		Thursday  = iota // 3
		Friday    = iota // 4
		Saturday  = iota // 5
		Sunday    = iota // 6
	)
	fmt.Println(Wednesday)
	const (
		Mon = iota
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)
	fmt.Println(Thu) //3

	// 文字列の出力
	var str1 string = "Hello "
	fmt.Println(str1)
}
