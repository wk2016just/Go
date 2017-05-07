# Go
大晚上闲来无事写写go，主要几点：
1，规范：使用“comma, ok”的习惯用法

来安全地测试值是否为一个字符串：
str, ok := value.(string)if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}

2，new和make

3，接口的各种组合，变量转换以及receiver的通用性。go的接口还是很好用的，完全免去java弱势~

4，goroutine chan waitgroup
