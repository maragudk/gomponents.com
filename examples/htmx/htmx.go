package main

import (
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func Navbar() Node {
	return Nav(hx.Boost("true"),
		A(Href("/"), Text("Home")),
		A(Href("/about"), Text("About")),
	)
}
