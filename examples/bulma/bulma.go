package main

import (
	g "github.com/maragudk/gomponents"
	b "github.com/willoma/bulma-gomponents"
)

func Navbar() g.Node {
	return b.Navbar(
		b.Dark,
		b.NavbarStart(
			b.NavbarAHref("/", "Home"),
			b.NavbarAHref("/about", "About"),
		),
	)
}
