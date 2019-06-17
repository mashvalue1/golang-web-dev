package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

type hotels []hotel

type hotel struct {
	Name, Address, City, Zip, Region string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	ch := hotels{
		// hotel{} で宣言したほうが良い
		{
			"Hotel1", "Address1", "City1", "Zip1", "Region1",
		},
		{
			"Hotel2", "Address2", "City2", "Zip2", "Region2",
		},
	}

	tpl.Execute(os.Stdout, ch)
}
