package icons

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Navbar() g.Node {
	return Nav(Class("bg-gray-800 flex items-center space-x-4"),
		A(Href("/"), g.Text("Home")),
		A(Href("/about"), g.Text("About")),
	)
}
