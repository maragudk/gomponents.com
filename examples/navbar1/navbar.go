package main

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Navbar() g.Node {
	return Nav(Class("main-nav"),
		Ol(
			Li(A(Href("/"), g.Text("Home"))),
			Li(A(Href("/contact"), g.Text("Contact"))),
			Li(A(Href("/about"), g.Text("About"))),
		),
	)
}
