# Go言語の勉強
## 基礎知識
- Goで開発をするには、、、
  1. 下記コマンドでプロジェクトフォルダを作成する。  
     ```
     $ mkdir [プロジェクトフォルダ]
     $ cd [プロジェクトフォルダ]
     $ go mod init [アプリケーション名]
     ```
  2. 拡張子が`.go`のファイルを作成する。
  3. 下記コマンドでgoファイルをビルドする。
     ```
     go build [goファイル名]
     ```
  4. 下記コマンドでビルドされた実行ファイルが動作することを確認する。
     ```
     [生成されたexeファイル]
     ```
  5. なお、ビルドと実行を一括で実施する場合は、下記コマンドを使用する。
     ```
     go run [goファイル名]
     ```

## 環境情報
- Goのバージョン: 1.22.0
- 使用するOS: Windows 11
- 使用するエディター: Visual Studio Code
  - プラグイン:
- 使用するCLI: コマンドプロンプト 

## 目次
1. [データ型と変数、定数](#データ型と変数定数)
    1. [データ型](#データ型)
    2. [変数の宣言](#変数の宣言)
    3. [定数の宣言](#定数の宣言)
    4. [特別な定数-`iota`](#特別な定数-iota)
    5. [文字列](#文字列)
    6. [配列とスライス](#配列とスライス)
    7. [マップ](#マップ)
    8. [struct（構造体）](#struct構造体)
    9. interface{}またはany
2. 基本文法
    1. 四則演算
    2. 条件分岐
    3. 繰り返し処理
    4. メソッド
    5. ポインタ
    6. コンストラクタ
    7. インターフェース
    8. 後処理
    9. ゴルーチン
    10. チャネル
    11. 制御構文（select）


## データ型と変数、定数
### データ型
#### 整数型  
|データ型|概要|備考|
|:--:|----|----|
|int8|符号付き8ビット整数型||
|int16|符号付き16ビット整数型||
|int32|符号付き32ビット整数型|int (32ビット環境)、ルーン (rune)ともいう|
int64|符号付き64ビット整数型|int (64ビット環境)ともいう|
|uint8|符号なし8ビット整数型|byteともいう|
|uint16|符号なし16ビット整数型||
|uint32|符号なし32ビット整数型|uint (32ビット環境)ともいう|
|uint64|符号なし64ビット整数型|uint (32ビット環境)ともいう|

#### 浮動小数点型
|データ型|概要|備考|
|:--:|----|----|
|float32|32ビット浮動小数点型||
|float64|64ビット浮動小数点型||

#### その他のデータ型
|データ型|概要|備考|
|:--:|----|----|
|string|文字列型。`"`で囲って宣言する。||
|bool|真偽型。`true` (真) や `false` (偽) を表す。||

### 変数の宣言
変数は下記のように、`var`を用いて宣言する。  
また、変数の宣言と初期値の入力を同時にすることも可能。  
但し、宣言した変数に異なるデータ型の値を入れたり、変数をそもそも使用しなかった場合、コンパイルエラーになるので、注意すること。
```go
// 宣言方法１
var [変数名] [データ型]
[変数名] = [値]

// 宣言方法２
var [変数名] [データ型] = [値]
```
変数の宣言時に、`var`を使用しない方法もある。  
この場合、右辺の値のデータ型が推測できる場合は、明示的にデータ型を宣言する必要はない。
```go
[変数名] := [値]
// 例（x, zはint型、yはfloat64型）
x := 1
y := 2.5
z := 1 + x * 2
```

### 定数の宣言
変数は下記のように、`const`を用いて宣言する。  
宣言で使用できるデータ型は、「真偽型・数値型・複素数型・文字列型」のみである。  
なお、変数と違い、宣言した定数を使用しなくても、コンパイルエラーにはならない。
```go
// 宣言方法
const [定数名] = [値]
```
定数に関しては、データ型を明示する必要はなく、使われ方で自動的にデータ型が決定する。  
その為、下記のように記載してもコンパイルエラーにはならないが、異なるデータ型の変数に対して処理をする場合は、キャストする必要がある。
```go
const x = 100
var num1 int = 10 + x
var num2 float64 = 12.5 + x
var num3 float64 = float64(num1) + num2
```
また、定数は下記のようにすることで、複数個をまとめて宣言することができる。  
一つずつ宣言すると、それぞれの定数に値が入力されるが、  
最初の一つだけ宣言すると、全ての定数に同じ値が入力される。
```go
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
fmt.Println(d, e, f) // 1 1 1
```

### 特別な定数 `iota`
iotaは値の並びを表現するために使われる特別な定数。  
各定数が一意な値を持つ列挙型を表現するためによく使用される。  
iotaは行が変わるたびに自動的にインクリメントされる。
```go
// 曜日を表す定数のセットを定義する例
const (
  Monday    = iota // 0
  Tuesday   = iota // 1
  Wednesday = iota // 2
  Thursday  = iota // 3
  Friday    = iota // 4
  Saturday  = iota // 5
  Sunday    = iota // 6
)
fmt.Println(Wednesday) // 2

// 曜日を表す定数のセットを定義する例
// 連続して記述する場合、2行目以降は省略可能
const (
  Mon = iota
  Tue
  Wed
  Thu
  Fri
  Sat
  Sun
)
```

### 文字列
Goで文字列を扱う場合、`"`で囲って使用する。  
コンソール出力する場合は、fmtパッケージの`Println`メソッドを使用する。  
文字列同士の結合も可能である。
```go
var str1 string = "Hello "
fmt.Println(str1) // Hello 
var str2 string = str1 + "world!"
fmt.Println(str2) // Hello world!
```
また、文字列はバイト列で構成されているため、添え字でのアクセスが可能である。  
但し、string型で宣言すると、イミュータブルな値になるため、値の書き換えができなくなる。  
書き換える場合は、バイト列に変換してから行う必要がある。
```go
fmt.Printf("%c, %c", str2[0], str2[8]) // H, r

str3 := []byte(str2)
str3[0] = 'h'
str2 = string(str3)
fmt.Println(str2) // hello world!
```
Goで複数行の文字列を扱う場合は、`` ` ``で囲うことで使用できる。
```go
var str4 = `今日の
夕飯は
お寿司でした。
`
fmt.Println(str4)
// 今日の
// 夕飯は
// お寿司でした。
// 
```

### 配列とスライス
Goには、配列とスライスという2つの型がある。  
配列は、固定長で、要素の型は1種類であり、indexは0から数える。    
スライスは、可変長で要素の型は1種類であり、こちらもindexは0から数える。  
  
配列は、下記のように宣言し、値を入力していない場合、int型であれば0が入力され、string型であれば`""`が入力される。  
なお、配列は固定長なので、配列の長さを超えた場合、panicが発生する。
```go
var list [5]int
list[0] = 10
list[1] = 20
fmt.Println(list) // [10 20 0 0 0]
```
一方のスライスでは、配列と同じように宣言してしまうと、`nil`という特別な値が代入されてしまい、要素にアクセスしようとすると、panicが発生してしまう。  
その為、`make`を用いて変数宣言も兼ねつつスライスの宣言を行う。  
`make`は、配列やスライス、マップを宣言する際に使用し、`make([型], [長さ], [最大長(容量)])`といった引数が必要となる。
```go
// panicが発生するスライスの宣言方法
var slice1 []int
// makeを用いて、スライスの宣言と初期化
slice1 := make([]int, 3)
fmt.Println(slice1) // [0 0 0]
// スライスの宣言と初期化の別の方法
var slice2 = []int{1, 2, 3}
fmt.Println(slice2) // [1 2 3]
```
スライスは可変長なので、スライスに要素を追加したり、削除することが可能である。  
要素を追加するには、`append`メソッドを用いて、スライスを上書きするイメージで更新する。  
一方で削除するには3種類の方法がある。  
(1)新しく同じ型のスライスを用意し、要素を追加しなおす。  
(2)`append`を使う。  
(3)部分参照と`copy`を使用する。  
なお、(2)ではアロケーションが1回発生してしまうため、(1)か(3)の方法が好ましい。  
※アロケーション：メモリ領域を2重で確保してしまうこと。
```go
// スライスに要素を追加する
var slice3 = []int{1, 2, 3}
fmt.Println(slice3) // [1 2 3]
slice3 = append(slice3, 4, 5, 6)
fmt.Println(slice3) // [1 2 3 4 5 6]

// スライスから要素を削除する
// (1)新しく同じ型のスライスを用意し、要素を追加しなおす。
// 例）1 - 10が記載されたスライスから奇数を削除する。
var slice4 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
fmt.Println(slice4) // [1 2 3 4 5 6 7 8 9 10]
var slice5 = make([]int, 0)
for i := 1; i < 11; i++ {
  if i%2 == 0 {
    slice5 = append(slice5, i)
  }
}
fmt.Println(slice5) // [2 4 6 8 10]
slice4 = slice5
fmt.Println(slice4) // [2 4 6 8 10]

// (2)appendを使う。
// 例）1 - 10が記載されたスライスから5を削除する。
var slice6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
fmt.Println(slice6) // [1 2 3 4 5 6 7 8 9 10]
target_index := 4
slice6 = append(slice6[:target_index], slice6[target_index+1:]...)
fmt.Println(slice6) // [1 2 3 4 6 7 8 9 10]

// (3)部分参照とcopyを使用する。
var slice7 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
fmt.Println("slice7: ", slice7) // [1 2 3 4 5 6 7 8 9 10]
n := 4
slice7 = slice7[:n+copy(slice7[n:], slice7[n+1:])]
fmt.Println("slice7: ", slice7) // [1 2 3 4 6 7 8 9 10]
```

### マップ
マップは、順序を持たない、キーと値のペアである。  
キーには、文字列以外も指定可能であり、宣言時にはmakeを用いて宣言する。  
なお、値は後から上書きすることは可能である。  
また、繰り返し処理を用いることでキーと値のペアを追加することも可能である。
```go
map1 := make(map[string]int)
map1["Hokkaido"] = 0
fmt.Println(map1) // map[Hokkaido:0]
map1["Hokkaido"] = 1
fmt.Println(map1) // map[Hokkaido:1]

map2 := make(map[int]string)
map2[1]  = "Hokkaido"
map2[47] = "Okinawa"
fmt.Println(map2) // map[1:Hokkaido 47:Okinawa]

map3 := make(map[int]string)
for i := 0; i < 10; i++ {
  map3[i] = "val_" + strconv.Itoa(i)
}
fmt.Println("map3", map3) // map[0:val_0 1:val_1 2:val_2 3:val_3 4:val_4 5:val_5 6:val_6 7:val_7 8:val_8 9:val_9]
```
for文を使ってキーと値のペアを取得することも可能であるが、マップは順不同な為、毎回ランダムな順番で取得される。  
順番を固定させるのであれば、キーだけを配列として取得し、ソート後に繰り返し処理をすることになる。
```go
// 繰り返し処理で表示
for key, value := range map3 {
  fmt.Println(key, value)
}
// ソートして表示
keys := []int{}
for k := range map3 {
  keys = append(keys, k)
}
sort.Ints(keys)
for k := range keys {
  fmt.Println(k, ":", map3[k])
}
```
マップから要素を削除するには、`delete(map, [削除対象のキー])`とする。
```go
delete(map3, 3)
fmt.Println("map3", map3) // map[0:val_0 1:val_1 2:val_2 4:val_4 5:val_5 6:val_6 7:val_7 8:val_8 9:val_9]
```
マップに存在しないキーを指定して値を取ろうとすると、値が文字列型の場合はブランク、数値型の場合は 0 となる。  
また、マップにキーが存在しているかを確認するには、値を取得する際に2つの変数で取得することで確認ができる。
```go
// 存在しないキーを指定した場合（バリュー：文字列）
map4 := make(map[int]string)
map4[1] = "One"
map4[2] = "two"
fmt.Println("map4", map4) // map4 map[1:One 2:two]
fmt.Println("map4[3]", map4[3]) //map4[3]
// 存在しないキーを指定した場合（バリュー；数値）
map5 := make(map[string]int)
map5["One"] = 1
map5["two"] = 2
fmt.Println("map5", map5) //map5 map[One:1 two:2]
fmt.Println("map5[Three]", map5["Three"]) // map5[Three] 0
// キーの存在確認
v_one, exists_one := map5["One"]
fmt.Println("Key One: Value =", v_one, "/ Exists = ", exists_one) // Key One: Value = 1 / Exists =  true
v_three, exists_three := map5["Three"]
fmt.Println("Key Three: Value =", v_three, "/ Exists = ", exists_three) // Key Three: Value = 0 / Exists =  false
```

### struct（構造体）
Goでは、C言語と同じような構造体を宣言可能である。  
これは、`type`を用いて型に名前を付けることの応用みたいなものである。  
この構造体を利用することで、JSON形式のデータも変数に格納することが可能である。
```go
type User struct {
  Name string
  age int
}
var user1 User 
user1.Name = "Demo1"
user1.age = 30
fmt.Println("user1: ", user1) // user1:  {Demo1 30}
fmt.Println("user1.Name: ", user1.Name) // user1.Name:  Demo1

user2 := User{
  Name: "Demo2",
  age:  31,
}
fmt.Println("user2: ", user2) // user2:  {Demo2 31}
fmt.Println("user2.Name: ", user2.Name) // user2.Name:  Demo2
```