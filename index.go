package main

import (
	_ "embed"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

//go:embed examples/simple.go
var simpleExample string

func IndexPage() g.Node {
	return Page(
		"gomponents, declarative view components in Go",
		"View components in pure Go, that render to HTML5.",
		Div(
			H1(g.Text("Tired of template languages?")),
			H2(g.Text("Try view components in pure Go.")),
			g.Raw(`<em>gomponents</em> are view components written in pure Go. They render to HTML 5, and make it easy for you to build reusable components. So you can focus on building your app instead of learning yet another templating language.`),

			H2(g.Text("Get started")),
			CodeBlock("$ go get -u github.com/maragudk/gomponents"),

			H3(g.Text("A sample app")),
			CodeBlock(simpleExample),

			P(g.Text("See "), A(Href("https://github.com/maragudk/gomponents"), g.Text("the Github repository")), g.Text(".")),
		),
	)
}
