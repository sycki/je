# json editor

## use in command line
### get value by path
```bash
echo '{"k1": {"kk1": "vv1", "kk2": "vv2"}, "k2": "v2"}' | je .k1.kk2
"vv2"
```

### set value
```bash
echo '{"k1": {"kk1": "vv1", "kk2": "vv2"}, "k2": "v2"}' | je .k1.kk2 vv0
{"k1": {"kk1": "vv1", "kk2": "vv0"}, "k2": "v2"}
```

## use in go code
### get value by path
```go
import github.com/sycki/je

func main() {
    str := `{"k1": {"kk1": "vv1", "kk2": "vv2"}, "k2": "v2"}`
    str2 := je.Get(str, ".k1.kk2")
    println(str2) // "vv2"
}
```

### set value by path
```go
import github.com/sycki/je

func main() {
    str := `{"k1": {"kk1": "vv1", "kk2": "vv2"}, "k2": "v2"}`
    str2 := je.Set(str, ".k1.kk2", "vv0")
    println(str2) // {"k1": {"kk1": "vv1", "kk2": "vv0"}, "k2": "v2"}
}
```
