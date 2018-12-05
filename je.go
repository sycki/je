package je

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sycki/je/cmd/je/option"
	"os"
	"strconv"
)

var (
	cons   = [...][]byte{[]byte("."), []byte("..")}
	maxErr = 5
)

const (
	str = iota
	index
	length
)

// same as GetB
func Get(o, path string) string {
	return string(GetB([]byte(o), []byte(path)))
}

// same as SetB
func Set(o, path string, value interface{}) string {
	return string(SetB([]byte(o), []byte(path), value))
}

// get json by path and trim the quotes before and after the result value
func GetB(o, path []byte) []byte {
	keys, err := parsePath(path)
	check(err)

	value := o
	for _, key := range keys {
		if len(value) < 1 {
			return nil
		}
		r, err := getOne(value, key)
		check(err)
		value, _ = json.Marshal(r)
	}

	return trimQuotes(value)
}

// set json by path
func SetB(o, path []byte, v interface{}) []byte {
	keys, err := parsePath(path)
	check(err)

	// json to map
	values := make(map[string]interface{}, len(keys))
	value := o
	for _, key := range keys {
		values[string(key)], err = getOne(value, key)
		check(err)
	}

	// set value
	key := keys[len(keys)-1]
	values[string(key)] = v

	// reverse map
	revalues := make(map[string]interface{}, len(keys))
	for i := len(keys) - 1; i >= 0; i-- {
		key := keys[i]
		revalues[string(key)] = values[string(key)]
	}

	// map to json
	lastKey := keys[len(keys)-1]
	lastValue := revalues[string(lastKey)]
	delete(revalues, string(lastKey))
	for k, v := range revalues {
		vs, _ := json.Marshal(v)
		lastValue, err = setOne(vs, lastKey, lastValue)
		check(err)
		lastKey = []byte(k)
	}

	lastValue, err = setOne(o, lastKey, lastValue)
	check(err)

	o, err = json.Marshal(lastValue)
	check(err)
	return o
}

// same as TypeB
func Type(v string) interface{} {
	return TypeB([]byte(v))
}

// convert string to int, array, struct or string
func TypeB(v []byte) interface{} {
	if len(v) < 1 {
		return ""
	}
	f := v[0]
	var r interface{}
	switch {
	case f == 123:
		r = make(map[string]interface{}, 10)
		json.Unmarshal(v, &r)
	case f == 91:
		r = make([]interface{}, 0, 10)
		json.Unmarshal(v, &r)
	case f >= 48 && f <= 57:
		r, _ = strconv.Atoi(string(v))
	case f == 34:
		r = string(trimQuotes(v))
	default:
		r = string(v)
	}
	return r
}

// get json by single key
func getOne(o, k []byte) (interface{}, error) {
	switch keyType(k) {
	case length:
		return getLen(k)
	case index:
		return getIndex(o, k)
	default:
		return getKey(o, k)
	}
}

// set json by single key
func setOne(o, k []byte, v interface{}) (interface{}, error) {
	switch keyType(k) {
	case index:
		return setIndex(o, k, v)
	default:
		return setKey(o, k, v)
	}
}

// get json by string key
func getKey(o, k []byte) (interface{}, error) {
	m := make(map[string]interface{})

	err := json.Unmarshal(o, &m)
	if err != nil {
		l := len(o)
		if l > maxErr {
			l = maxErr
		}
		return nil, fmt.Errorf("can not parse [%v]", string(o[:l]))
	}

	return m[string(k)], nil
}

// set json by string key
func setKey(o, k []byte, v interface{}) (interface{}, error) {
	m := make(map[string]interface{})

	err := json.Unmarshal(o, &m)
	if err != nil {
		l := len(o)
		if l > maxErr {
			l = maxErr
		}
		return nil, fmt.Errorf("can not parse [%v]", string(o[:l]))
	}

	m[string(k)] = v

	return m, nil
}

// get json by index of the array
func getIndex(o, key []byte) (interface{}, error) {
	i, _ := strconv.Atoi(string(key))
	m := make([]interface{}, 0)

	err := json.Unmarshal(o, &m)
	if err != nil {
		l := len(o)
		if l > maxErr {
			l = maxErr
		}
		return nil, fmt.Errorf("can not parse [%v]", string(o[:l]))
	}

	if i >= len(m) {
		return nil, fmt.Errorf("not found index [%v]", i)
	}

	return m[i], nil
}

// set json by index of the array
func setIndex(o, key []byte, v interface{}) (interface{}, error) {
	i, _ := strconv.Atoi(string(key))
	m := make([]interface{}, 0)

	err := json.Unmarshal(o, &m)
	if err != nil {
		l := len(o)
		if l > maxErr {
			l = maxErr
		}
		return nil, fmt.Errorf("can not parse [%v]", string(o[:l]))
	}

	m[i] = v

	return m, nil
}

// get length of the array
func getLen(o []byte) (int, error) {
	m := make([]interface{}, 0)

	err := json.Unmarshal(o, &m)
	if err != nil {
		l := len(o)
		if l > maxErr {
			l = maxErr
		}
		return 0, fmt.Errorf("can not parse [%v]", string(o[:l]))
	}

	return len(m), nil
}

// parse and check the key path
func parsePath(path []byte) ([][]byte, error) {
	if len(path) < 1 || path[0] != cons[0][0] {
		return nil, fmt.Errorf("first char must is [.] of path")
	}

	if bytes.Contains(path, cons[1]) {
		return nil, fmt.Errorf("key path can not contains [..]")
	}

	keys := bytes.Split(path, cons[0])[1:]

	return keys, nil
}

// parsed the key type
func keyType(key []byte) int {
	if string(key) == "#" {
		return length
	}
	_, err := strconv.Atoi(string(key))
	if err == nil {
		return index
	}
	return str
}

// trim the quotes before and after the string
func trimQuotes(str []byte) []byte {
	str = bytes.TrimPrefix(str, []byte(`"`))
	str = bytes.TrimSuffix(str, []byte(`"`))
	return str
}

// exit process if error occurs
func check(err error) {
	if option.Conf.Cmd {
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err.Error())
			os.Exit(2)
		}
	}
}
