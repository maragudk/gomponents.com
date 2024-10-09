package icons

import (
	"github.com/maragudk/gomponents-heroicons/v2/solid"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Navbar() Node {
	return Nav(Class("bg-gray-800 flex items-center space-x-4"),
		solid.ArrowRight(Class("h-6 w-6")),
		A(Href("/"), Text("Home")),
		A(Href("/about"), Text("About")),
	)
}
