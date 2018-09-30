package main

import "flag"

type option struct {
	file	string
}

func (o *option) AddArgs(cmd *flag.FlagSet) {
	cmd.StringVar(&o.file, "f", "", "")
}

func NewDefaultOption() *option {
	return &option{}
}
