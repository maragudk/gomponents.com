package main

import (
	x "github.com/glsubri/gomponents-alpine"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Navbar() Node {
	return Nav(
		A(Href("/"), Text("Home")),

		x.Data("{ submenu: false }"),
		A(x.On("click", "submenu = ! submenu"), Href("/about"), Text("About")),
		Div(x.Show("submenu"), Div(
		// …
		)),
	)
}
