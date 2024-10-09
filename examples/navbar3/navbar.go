package main

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type NavLink struct {
	Name string
	Path string
}

func Navbar(loggedIn bool, links []NavLink) Node {
	return Nav(Class("navbar"),
		Ol(
			Map(links, func(l NavLink) Node {
				return NavbarItem(l.Name, l.Path)
			}),

			If(loggedIn,
				NavbarItem("Log out", "/logout"),
			),
		),
	)
}

func NavbarItem(name, path string) Node {
	return Li(A(Href(path), Text(name)))
}
