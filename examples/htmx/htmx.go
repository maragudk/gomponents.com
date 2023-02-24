package main

import (
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

func Navbar() g.Node {
	return Nav(hx.Boost("true"),
		A(Href("/"), g.Text("Home")),
		A(Href("/about"), g.Text("About")),
	)
}
