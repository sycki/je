package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sycki/je"
)

func main() {
	cmd := flag.CommandLine
	opt := NewDefaultOption()
	opt.AddArgs(cmd)
	flag.Parse()

	var data, path, value, result string
	if opt.file != "" {
		b, err := ioutil.ReadFile(opt.file)
		if err != nil {
			println(err.Error())
			os.Exit(3)
		}
		data = string(b)
	} else {
		b, _ := ioutil.ReadAll(os.Stdin)
		data = string(b)
	}

	args := cmd.Args()
	if len(args) == 1 {
		path = args[0]
		result = je.Get(data, path)
	} else if len(args) == 2 {
		path = args[0]
		value = args[1]
		result = je.Set(data, path, value)
	} else {
		cmd.Usage()
		os.Exit(255)
	}

	fmt.Print(result)
}
