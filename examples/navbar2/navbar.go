package main

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Navbar(loggedIn bool) Node {
	return Nav(Class("navbar"),
		Ol(
			NavbarItem("Home", "/"),
			NavbarItem("Contact", "/contact"),
			NavbarItem("About", "/about"),

			If(loggedIn,
				NavbarItem("Log out", "/logout"),
			),
		),
	)
}

func NavbarItem(name, path string) Node {
	return Li(A(Href(path), Text(name)))
}
