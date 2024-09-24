package main

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

type NavLink struct {
	Name string
	Path string
}

func Navbar(loggedIn bool, links []NavLink) g.Node {
	return Nav(Class("navbar"),
		Ol(
			g.Map(links, func(l NavLink) g.Node {
				return NavbarItem(l.Name, l.Path)
			}),

			g.If(loggedIn,
				NavbarItem("Log out", "/logout"),
			),
		),
	)
}

func NavbarItem(name, path string) g.Node {
	return Li(A(Href(path), g.Text(name)))
}
