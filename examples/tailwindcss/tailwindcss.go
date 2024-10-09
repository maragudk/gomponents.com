package main

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Navbar() Node {
	return Nav(Class("bg-gray-800 flex items-center space-x-4"),
		A(Href("/"), Text("Home")),
		A(Href("/about"), Text("About")),
	)
}
