package main

import (
	"flag"
	"fmt"
	"github.com/sycki/je"
	"github.com/sycki/je/cmd/je/option"
	"io/ioutil"
	"os"
)

var version string

func main() {
	cmd := flag.CommandLine
	opt := option.NewOption()
	opt.AddArgs(cmd)
	flag.Parse()

	if opt.Version {
		println(version)
		return
	}

	// load input data
	var data, path, value, result []byte
	if opt.File != "" {
		b, err := ioutil.ReadFile(opt.File)
		if err != nil {
			println(err.Error())
			os.Exit(3)
		}
		data = b
	} else {
		b, _ := ioutil.ReadAll(os.Stdin)
		data = b
	}

	// processing
	args := cmd.Args()
	if len(args) == 1 {
		path = []byte(args[0])
		result = je.GetB(data, path)
	} else if len(args) == 2 {
		path = []byte(args[0])
		value = []byte(args[1])
		result = je.SetB(data, path, je.TypeB(value))
	} else {
		cmd.Usage()
		os.Exit(255)
	}

	// print result
	fmt.Println(string(result))
}
