# 自學筆記
[Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU)\


## [變數](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=2148s)
- 宣告
    - 同樣的變數名稱不能重複宣告(declaration)
    - 定義在 package level 的變數，可以用 shadowing 的技巧，在 local scope 重新宣告
    - 所有宣告的變數必須被使用
- Visibility
    - https://golang.org/doc/effective_go.html#names
    - 首字小寫(myVar)，for package scope
    - 首字大寫(MyVar)，會 export 出去，別人可以用
    - 無私有 scope 的概念
    - *the visibility of a name outside a package is determined by whether its first character is upper case.*

## [Primitive types](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=3425s)
- int
    - zero: 0
    - default value: 0 (`var a int`)
    - 支援 bitwise operations
    - 同時有 unsigned 的版本 (int8, uint8)
    - 支援 加減乘除餘
    - 同種型別才能做數學運算
- float
    - zero: 0
    - 3.14
    - 2E10
    - 2.187e12
    - 支援 加減乘除 (沒有餘)
- Complex
    - zero: 0 + 0i
    - real(i): get real part
    - imag(i): get imaginary part
    - 支援 加減乘除 (沒有餘)
- Text
    - String
        - UTF-8 (8bits)
        - immutable, 無法改變特定 index 的值
        - 支援 concatenation (+)
        - 可以透過 []byte 轉換成 array of ascii number         
    - rune
        - UTF-32 (32bits)
        - alias for int32。[rune](https://juejin.im/post/5b44caebf265da0f491b8b83)，與 int32 一樣，但用來區別整數值與字符值

## [Constants](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=5189s)
- 雖然本質上是immutable 物件，但可以被 shadow，當你在 package level 宣告時，inner scope 可以重新宣告、甚至改變型別
- replaced by the compiler at compile time
    - 此例，a + b 會被明確變成 int8 + int16 而造成錯誤
    ```golang
    package main
    import "fmt"
    func main() {
        const a int8 = 42
        var b int16 = 16
        fmt.Printf("%v, %T", a+b, a+b)
    }
    // invalid operation: a + b (mismatched types int and int16)
    ```
    - 此例，a 會被取代成 42，42 + int8是可以的
    ```golang
    package main
    import "fmt"
    func main() {
        const a = 42
        var b int16 = 16
        fmt.Printf("%v, %T", a+b, a+b)
    }
    // 58, int16
    // compiler 會將 a 取代成 42
    ```
- 值必須在 compile time 被決定，故不能用一個 function 的 return 來指定 const 物件
    - `const a float32 = math.Sin(1.57) //const initializer math.Sin(1.57) is not a constant` 
- enumerated type
    - [4 iota enum examples](https://yourbasic.org/golang/iota/)
    - iota 起始值為0，每加一行就+1。不同的 const 區塊會各自計算
    - 注意未於宣告時給定的數值可能有初始值0，並且在布林比對時匹配到 iota 第0個位置的enum
    - type of iota is int
    - Enumerated expressions 允許的計算必須在 compile time 就發生
        - 數學運算
        - bitwise operations
        - bitshifting 
    ```golang
    package main
    import "fmt"
    cosnt (
        a = iota
        b
        c
    )
    func main() {
        fmt.Println(a, b, c) // 0 1 2
    }
    // start with 1
    cosnt (
        a = iota + 1
        b
        c
    )
    // skip value, use 'blank identifier'
    cosnt (
        a = iota + 1
        _
        c
        d
    )
    ```

## [Arrays and Slices](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=6473s)
- [array](https://www.golangprograms.com/go-language/arrays.html) vs. [slice](https://www.golangprograms.com/go-language/slices-in-golang-programming.html)
    - `a =: [3]int` - array
    - `a =: [...]int{1, 2, 3}` - array
    - `a =: []int{}` - slice
- array
    - 固定長度
    - declaration styles
        1. `a := [5]int{1,2,3}` - [1 2 3 0 0] can do Partial assignment
        2. `a := [...]int{1,2,3}` - Initializing an Array with ellipses in Go, the compiler can identify the length of an array, based on the elements specified in the array declaration.
        3. `var a [3]int` - 
    - **len** return size of array; **cap** 回傳預先配置的記憶體長度
    - copy (`a := b`) 會創建一塊一樣大小的記憶體出來
- slice
    - 所有切出來的都是參考
    - creation styles
        - 從既有的 array 中切出
        - slice literal
            - array literal: `[3]int{1,2,3}`
            - slice literal: `[]int{1,2,3}` - 創建同上的 array 之後，再建一個 slice 指向該 array ( [array literal vs. slice literal](https://tour.golang.org/moretypes/9))
        - via make function
            - `a := make([]int, 10)` - create slice with capacity and length == 10
            - `a := make([]int, 10, 100)` - create slice with capacity == 100 and length == 10
    - `len`: return length of slice
    - `cap`: return length of underlying array
    - `append`: add elements to slice
        - slice 無固定長度 
    - copy 的行為會指向 same underlying array
    
## map and struct
- map
    - 一樣可用 `make` 來創建
    - 用 [key] 取值
    - 用 value, ok = a[key] 來確認 key 是否存在
    - map 的 copy 行為為 reference
- struct
    - 可匿名創建 struct 
    - copy 行為為真複製
    - embedded struct
        - 無真正的繼承行為，實際上只是 has-a relationship，也就是 composition 而已
        - golang 有 Interface 的 feature 去實現 OOP 中的繼承概念
    - tag 
        - 用來補充 struct 欄位的額外 meta info
        - 可以用 `reflect` 的一系列方法來取得這些 meta info
        - check out: https://zhengyinyong.com/field-tag-in-go-struct.html

## [control flow: If and Switch](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=10080s)
