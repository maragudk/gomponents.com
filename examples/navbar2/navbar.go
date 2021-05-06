package main

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Navbar(loggedIn bool) g.Node {
	return Nav(Class("main-nav"),
		Ol(
			NavbarItem("Home", "/"),
			NavbarItem("Contact", "/contact"),
			NavbarItem("About", "/about"),

			g.If(loggedIn,
				NavbarItem("Log out", "/logout"),
			),
		),
	)
}

func NavbarItem(name, path string) g.Node {
	return Li(A(Href(path), g.Text(name)))
}
