package icons

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents-heroicons/solid"
	. "github.com/maragudk/gomponents/html"
)

func Navbar() g.Node {
	return Nav(Class("bg-gray-800 flex items-center space-x-4"),
		solid.ArrowRight(Class("h-6 w-6")),
		A(Href("/"), g.Text("Home")),
		A(Href("/about"), g.Text("About")),
	)
}
