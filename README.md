# Json editor
The Json editor is an easy-to-use json query and editing tool that you can use in command line and golang projects.

## Documents
* [中文文档](https://github.com/sycki/je/blob/master/README_ZH.md)
* [English doc](https://github.com/sycki/je)

## Install
```
curl -Lo je https://sycki.com/f/bin/je          // for linux
curl -Lo je https://sycki.com/f/bin/je-darwin   // for mac
chmod +x je
mv je /usr/local/bin/je                         // move to your PATH
```

## Use in command line
### Get
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

### Set
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

## Use in golang code
More demo see `je_test.go`

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
