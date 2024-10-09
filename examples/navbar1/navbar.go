package main

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Navbar() Node {
	return Nav(Class("navbar"),
		Ol(
			Li(A(Href("/"), Text("Home"))),
			Li(A(Href("/contact"), Text("Contact"))),
			Li(A(Href("/about"), Text("About"))),
		),
	)
}
