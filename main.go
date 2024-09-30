package main

import (
	"github.com/alecthomas/kong"
)

type Context struct {
	Debug bool
}

type RmCmd struct {
	Pattern string `help:"File Pattern, e.g. Landrary_$date.log" short:"p" default:"localhost.$date.log"`
	Day     int    `help:"Number of days to keep." short:"d" default:"30"`

	Paths []string `arg:"" name:"path" help:"Paths to remove." type:"path"`
}

func (r *RmCmd) Run(ctx *Context) error {
	RemoveLogFiles(r.Paths, r.Pattern, r.Day, ctx.Debug)
	return nil
}

type LsCmd struct {
	Pattern string   `help:"File Pattern, e.g. Landrary_#date#.log" optional:"false" short:"p"  default:"localhost.#date#.log"`
	Day     int      `help:"Number of days to keep." short:"d" default:"30"`
	Paths   []string `arg:"" name:"path" help:"Paths to list." type:"path"`
}

func (l *LsCmd) Run(ctx *Context) error {
	PrintLogFiles(l.Paths, l.Pattern, l.Day, ctx.Debug)
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Rm RmCmd `cmd:"" help:"Remove files."`
	Ls LsCmd `cmd:"" help:"List paths."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
