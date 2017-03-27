//go:generate go-bindata -o tpl.go tpl
package main

import (
	"flag"
	"html/template"
	"os"
)

func main() {
	var useGB bool
	flag.BoolVar(&useGB, "go-bindata", false, "")
	flag.Parse()

	var t *template.Template
	if useGB {
		b, _ := Asset("tpl/page.html")
		t, _ = template.New("").Parse(string(b))
	} else {
		t = template.Must(template.ParseFiles("tpl/page.html"))
	}

	if t != nil {
		t.ExecuteTemplate(os.Stdout, "base", nil)
	}
}
