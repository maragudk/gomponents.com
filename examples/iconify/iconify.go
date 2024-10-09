package iconify

import (
	"github.com/delaneyj/gomponents-iconify/iconify/maki"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Navbar() Node {
	return Nav(Class("bg-gray-800 flex items-center space-x-4"),
		maki.Beach(),
		A(Href("/"), Text("Home")),
		A(Href("/about"), Text("About")),
	)
}
