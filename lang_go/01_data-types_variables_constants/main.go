// 基本的な文法に関する内容を記載
package main

import (
	"fmt"
	"strconv"
)

// グローバル変数の宣言
var word string = "Hello world."

func main() {
	// 変数と定数について
	variablesAndConstants()

	// 文字列について
	strings()

	// 配列とスライスについて
	arraysAndSlices()

	// マップについて
	maps()
}

func variablesAndConstants() {
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
}

func strings() {
	// 文字列の出力
	var str1 string = "Hello "
	fmt.Println(str1)
	// 文字列の結合
	var str2 string = str1 + "world!"
	fmt.Println(str2)
	fmt.Println(str1 + "world!")
	fmt.Printf("%c %c\n", str2[0], str2[8])
	// 文字列の書き換え
	str3 := []byte(str2)
	str3[0] = 'h'
	str2 = string(str3)
	fmt.Println(str2)
	// 複数行の文字列の宣言
	var str4 = `今日の
夕飯は
お寿司でした。
`
	fmt.Println(str4)
}

func arraysAndSlices() {
	// 配列
	var list1 [5]int
	list1[0] = 10
	list1[1] = 20
	fmt.Println(list1)
	fmt.Println(list1[3])

	// スライス
	var slice1 []int
	fmt.Println(slice1)
	var slice2 = make([]int, 3)
	fmt.Println("slice2: ", slice2)
	fmt.Println("slice2[2]: ", slice2[2])
	slice2 = append(slice2, 10)
	fmt.Println("slice2: ", slice2)
	var slice3 = []int{1, 2, 3}
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice3[1]: ", slice3[1])
	// スライスの要素の削除
	// (1)の場合
	var slice4 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice4: ", slice4)
	var slice5 = make([]int, 0)
	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			slice5 = append(slice5, i)
		}
	}
	fmt.Println("slice5: ", slice5)
	slice4 = slice5
	fmt.Println("slice4: ", slice4)
	// (2)の場合
	var slice6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice6: ", slice6) // [1 2 3 4 5 6 7 8 9 10]
	target_index1 := 4
	slice6 = append(slice6[:target_index1], slice6[target_index1+1:]...)
	fmt.Println("slice6: ", slice6)
	// (3)の場合
	var slice7 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice7: ", slice7) // [1 2 3 4 5 6 7 8 9 10]
	n := 4
	slice7 = slice7[:n+copy(slice7[n:], slice7[n+1:])]
	fmt.Println("slice7: ", slice7) // [1 2 3 4 6 7 8 9 10]
}

func maps() {
	var map1 = make(map[string]int)
	map1["Hokkaido"] = 1
	fmt.Println(map1)
	map1["Hokkaido"] = 3
	fmt.Println("map1", map1)

	map2 := make(map[int]string)
	map2[1] = "Hokkaido"
	map2[47] = "Okinawa"
	fmt.Println("map2", map2)

	map3 := make(map[int]string)
	for i := 0; i < 10; i++ {
		map3[i] = "val_" + strconv.Itoa(i)
	}
	fmt.Println("map3", map3)
	fmt.Println("map3[1]", map3[1])
	fmt.Println("順不同で表示")
	for key, value := range map3 {
		fmt.Println(key, value)
	}

	delete(map3, 3)
	fmt.Println("map3", map3)
}
