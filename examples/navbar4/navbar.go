package main

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

type NavLink struct {
	Name string
	Path string
}

func Navbar(loggedIn bool, links []NavLink, currentPath string) Node {
	return Nav(Class("navbar"),
		Ol(
			Map(links, func(l NavLink) Node {
				return NavbarItem(l.Name, l.Path, l.Path == currentPath)
			}),

			If(loggedIn,
				NavbarItem("Log out", "/logout", false),
			),
		),
	)
}

func NavbarItem(name, path string, active bool) Node {
	return Li(A(Href(path), Text(name)), Classes{
		"navbar-item": true,
		"active":      active,
		"inactive":    !active,
	})
}
