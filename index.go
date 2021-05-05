package main

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func IndexPage() g.Node {
	return Page(
		"gomponents, declarative view components in Go",
		"Declarative view components in Go, that can render to HTML5.",
		Div(
			H1(g.Text("gomponents")),
			H2(g.Text("Declarative view components in Go.")),
			P(g.Text("See "), A(Href("https://github.com/maragudk/gomponents"), g.Text("the Github repository")), g.Text(".")),
		),
	)
}
