package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/alecthomas/kong"
	"log"
	"os"
	"strings"
)

func main() {
	new(cmd).run()
}

//go:embed game.xml
var tpl string

type options struct {
	Template string   `name:"template" help:"Output template for game items"`
	About    bool     `help:"Show about."`
	Files    []string `arg:"" optional:""`
}

type cmd struct {
	options
}

func (c *cmd) run() {
	kong.Parse(&c.options,
		kong.Name("gamelist"),
		kong.Description("Convert Peagaus's metadata.txt to EmulationStation's gamelist.xml."),
		kong.UsageOnError(),
	)
	for _, f := range c.Files {
		c.convert(f)
	}
}

func (c *cmd) convert(name string) {
	src, err := os.Open(name)
	if err != nil {
		log.Printf("cannot open %s: %s", name, err)
		return
	}
	defer src.Close()
	var a ListStrStr
	var m MapStrStr
	for sc := bufio.NewScanner(src); sc.Scan(); {
		k, v, ok := strings.Cut(sc.Text(), ":")
		if !ok {
			continue
		}
		k, v = strings.TrimSpace(k), strings.TrimSpace(v)
		if k == "game" {
			if m.Has("game") {
				a = append(a, m)
			}
			m = nil
		}
		if m == nil {
			m = make(MapStrStr)
		}
		m.Set(k, v)
	}
	for _, m := range a {
		fmt.Println(m.Render(tpl))
	}
}
