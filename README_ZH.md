# Json editor
Json editor 是一个易于使用的 json 查询和编辑工具，你可以在命令行和 golang 项目中使用它。

## 文档
* [中文文档](https://github.com/sycki/je/blob/master/README_ZH.md)
* [English doc](https://github.com/sycki/je)

## 安装
```
curl -Lo je https://sycki.com/f/bin/je          // for linux
curl -Lo je https://sycki.com/f/bin/je-darwin   // for mac
chmod +x je
mv je /usr/local/bin/je                         // move to your PATH
```

## 在命令行中使用
### 查询操作
- Get string
    ```bash
    echo '{"k1":"v1","k2":2,"k3":[{"k31":"v31"},{"k32":"v32"},{"k33":"v33"}]}' | je .k1
    v1
    ```
- Get string
    ```bash
    echo '{"k1":"v1","k2":2,"k3":[{"k31":"v31"},{"k32":"v32"},{"k33":"v33"}]}' | je .k2
    2
    ```
- Get length string
    ```bash
    echo '{"k1":"v1","k2":2,"k3":[{"k31":"v31"},{"k32":"v32"},{"k33":"v33"}]}' | je .k3.#
    3
    ```
- Get string
    ```bash
    echo '{"k1":"v1","k2":2,"k3":[{"k31":"v31"},{"k32":"v32"},{"k33":"v33"}]}' | je .k3.0.k31
    v31
    ```

### 编辑操作
- Set struct
    ```bash
    echo '{"k1":{"k11":"v11","k12":"v12"},"k2":"v2"}' | je .k2 '{"k21":"v21"}'
    {"k1":{"k11":"v11","k12":"v12"},"k2":{"k21":"v21"}}
    ```
- Set array
    ```bash
    echo '{"k1":{"k11":"v11","k12":"v12"},"k2":"v2"}' | je .k2 '[{"k21":"v21"}]'
    {"k1":{"k11":"v11","k12":"v12"},"k2":[{"k21":"v21"}]}
    ```
- Set int
    ```bash
    echo '{"k1":{"k11":"v11","k12":"v12"},"k2":"v2"}' | je .k2 '5'
    {"k1":{"k11":"v11","k12":"v12"},"k2":5}
    ```
- Set string
    ```bash
    echo '{"k1":{"k11":"v11","k12":"v12"},"k2":"v2"}' | je .k2 '"5"'
    {"k1":{"k11":"v11","k12":"v12"},"k2":"5"}
    ```
- Set string
    ```bash
    echo '{"k1":{"k11":"v11","k12":"v12"},"k2":"v2"}' | je .k2 \"5\"
    {"k1":{"k11":"v11","k12":"v12"},"k2":"5"}
    ```
- Set string
    ```bash
    echo '{"k1":{"k11":"v11","k12":"v12"},"k2":"v2"}' | je .k2 v0
    {"k1":{"k11":"v11","k12":"v12"},"k2":"v0"}
    ```

## 在golang项目中使用
### Get
```go
import github.com/sycki/je

func main() {
    str := `{"k1":"v1","k2":2,k3":[{"k31":"v31"},{"k32":"v32"},{"k33":"v33"}]}`
    r := je.Get(str, ".k3.0.k31")
    println(r) // v31
}
```

### Set
```go
import github.com/sycki/je

func main() {
    str := `{"k1":"v1","k2":2,k3":[{"k31":"v31"},{"k32":"v32"},{"k33":"v33"}]}`
    r := je.Set(str, ".k3", 3)
    println(r) // {"k1":"v1","k2":2,k3":3}
}
```

更多的使用例子请看源码文件 `je_test.go`
