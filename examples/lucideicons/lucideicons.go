package main

import (
	"github.com/eduardolat/gomponents-lucide"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func myPage() Node {
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
