package option

import "flag"

var Conf = &option{Cmd: false}

type option struct {
	Cmd     bool
	File	string
	Version bool
}

func (o *option) AddArgs(cmd *flag.FlagSet) {
	cmd.StringVar(&o.File, "f", "", "specify input file")
	cmd.BoolVar(&o.Version, "version", false, "print je version information and exit")
}

func NewOption() *option {
	Conf.Cmd = true
	return Conf
}
