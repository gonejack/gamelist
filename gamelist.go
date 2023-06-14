package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type GameList struct {
	XMLName xml.Name `xml:"gameList"`
	Game    []Game   `xml:"game"`
}

type Game struct {
	Path      string `xml:"path"`
	Name      string `xml:"name"`
	Developer string `xml:"developer"`
	Desc      string `xml:"desc"`
	Players   string `xml:"players"`
}

type ListStrStr []MapStrStr
type MapStrStr map[string]string

func (m MapStrStr) Set(k, v string) {
	m[k] = v
}
func (m MapStrStr) Get(k string) string {
	return m[k]
}
func (m MapStrStr) Has(k string) (yes bool) {
	_, yes = m[k]
	return
}
func (m MapStrStr) Render(tpl string) string {
	var rps []string
	for k, v := range m {
		rps = append(rps, fmt.Sprintf("{p.%s}", k), v)
	}
	return strings.NewReplacer(rps...).Replace(tpl)
}
