package main

import (
	. "maragu.dev/gomponents"
	ds "maragu.dev/gomponents-datastar"
	. "maragu.dev/gomponents/html"
)

func Navbar() Node {
	return Nav(
		A(Href("/"), Text("Home"), ds.On("click", "@get('/')")),
		A(Href("/about"), Text("About"), ds.On("click", "@get('/about')")),
	)
}
