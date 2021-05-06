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
	return Nav(Class("main-nav"),
		Ol(
			g.Group(g.Map(len(links), func(i int) g.Node {
				return NavbarItem(links[i].Name, links[i].Path)
			})),

			g.If(loggedIn,
				NavbarItem("Log out", "/logout"),
			),
		),
	)
}

func NavbarItem(name, path string) g.Node {
	return Li(A(Href(path), g.Text(name)))
}
