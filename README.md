# 進度

https://youtu.be/YS4e4q9oBaU?t=3158


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

