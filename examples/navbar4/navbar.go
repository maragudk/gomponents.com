package main

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type NavLink struct {
	Name string
	Path string
}

func Navbar(loggedIn bool, links []NavLink, currentPath string) g.Node {
	return Nav(Class("navbar"),
		Ol(
			g.Group(g.Map(links, func(l NavLink) g.Node {
				return NavbarItem(l.Name, l.Path, l.Path == currentPath)
			})),

			g.If(loggedIn,
				NavbarItem("Log out", "/logout", false),
			),
		),
	)
}

func NavbarItem(name, path string, active bool) g.Node {
	return Li(A(Href(path), g.Text(name)), c.Classes{
		"navbar-item": true,
		"active":      active,
		"inactive":    !active,
	})
}
