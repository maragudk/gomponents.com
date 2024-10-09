package main

import (
	b "github.com/willoma/bulma-gomponents"
	. "maragu.dev/gomponents"
)

func Navbar() Node {
	return b.Navbar(
		b.Dark,
		b.NavbarStart(
			b.NavbarAHref("/", "Home"),
			b.NavbarAHref("/about", "About"),
		),
	)
}
