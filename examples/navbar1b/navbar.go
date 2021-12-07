package main

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Navbar() g.Node {
	return Nav(Class("navbar"),
		Ol(
			NavbarItem("Home", "/"),
			NavbarItem("Contact", "/contact"),
			NavbarItem("About", "/about"),
		),
	)
}

func NavbarItem(name, path string) g.Node {
	return Li(A(Href(path), g.Text(name)))
}
