package je

import (
	"encoding/json"
	"fmt"
	"strings"
)

// get json by path
func Get(o, path string) string {
	keys, err := parsePath(path)
	if err != nil {
		println(err.Error())
	}

	value := o
	for _, key := range keys {
		if value == "" {
			return value
		}
		value = GetByKey(value, key)
	}

	return value
}

// set json by path
func Set(o, path, v string) string {
	keys, err := parsePath(path)
	if err != nil {
		println(err.Error())
	}

	// json to map
	values := make(map[string]string, len(keys))
	value := o
	for _, key := range keys {
		value = GetByKey(value, key)
		values[key] = value
	}

	// set value
	values[keys[len(keys)-1]] = v

	// reverse map
	revalues := make(map[string]string, len(keys))
	for i := len(keys) - 1; i >= 0; i-- {
		key := keys[i]
		revalues[key] = values[key]
	}

	// map to json
	lastKey := keys[len(keys)-1]
	lastValue := revalues[lastKey]
	delete(revalues, lastKey)
	for k, v := range revalues {
		lastValue = SetByKey(v, lastKey, lastValue)
		lastKey = k
	}

	o = SetByKey(o, lastKey, lastValue)
	return o
}

// get json by short key
func GetByKey(o, k string) string {
	if k == "" {
		return o
	}
	m := make(map[string]interface{})

	err := json.Unmarshal([]byte(o), &m)
	if err != nil {
		println(err.Error())
		return ""
	}

	rb, err := json.Marshal(m[k])
	if err != nil {
		println(err.Error())
		return ""
	}

	return string(rb)
}

// set json by short key
func SetByKey(o, k, v string) string {
	m := make(map[string]interface{})

	err := json.Unmarshal([]byte(o), &m)
	if err != nil {
		println(err.Error())
		return ""
	}

	var vo interface{} = ""
	err = json.Unmarshal([]byte(v), &vo)
	if err != nil {
		vo = v
	}

	m[k] = vo

	rb, err := json.Marshal(m)
	if err != nil {
		println(err.Error())
		return ""
	}

	return string(rb)
}

// parse and check the key path
func parsePath(path string) ([]string, error) {
	if len(path) < 1 || string(path[0]) != "." {
		return nil, fmt.Errorf("first char must is dot of path.")
	}

	keys := strings.Split(path, ".")[1:]

	return keys, nil
}
