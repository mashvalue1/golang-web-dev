package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

type region struct {
	Region string
	Hotels hotels
}

type hotels []hotel

type hotel struct {
	Name, Address, City, Zip string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := region{
		Region: "Southern",
		Hotels: hotels{
			hotel{
				"Hotel1", "Address1", "City1", "Zip1",
			},
			hotel{
				"Hotel2", "Address2", "City2", "Zip2",
			},
		},
	}

	tpl.Execute(os.Stdout, r)
}
