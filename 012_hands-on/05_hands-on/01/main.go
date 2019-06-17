package main

import (
	"os"
	"text/template"
)

type item struct {
	Name, Desc, Meal string
	Price            int
}

type items []item

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	i := items{
		item{
			Name:  "foo",
			Desc:  "desc",
			Meal:  "Breakfast",
			Price: 1000,
		},
		{
			Name:  "foo",
			Desc:  "desc",
			Meal:  "Lunch",
			Price: 10000,
		},
	}

	tpl.Execute(os.Stdout, i)
}
