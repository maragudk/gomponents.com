package main

import (
	"github.com/eduardolat/gomponents-lucide"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func myPage() g.Node {
	return Div(
		lucide.Star(),
		lucide.Languages(),
		lucide.Usb(),
		//...
		lucide.Cherry(
			// You can add any attributes you want
			// to customize the SVG
			Class("w-6 h-6 text-blue-500"),
		),
	)
}
