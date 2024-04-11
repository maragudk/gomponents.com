package main

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func myPage() gomponents.Node {
	return html.Div(
		lucide.Star(),
		lucide.Languages(),
		lucide.Usb(),
		//...
		lucide.Cherry(
			// You can add any attributes you want
			// to customize the SVG
			html.Class("w-6 h-6 text-blue-500"),
		),
	)
}
