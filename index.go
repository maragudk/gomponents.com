package main

import (
	_ "embed"
	"strings"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

//go:embed examples/simple.go
var simpleExample string

//go:embed examples/navbar1/navbar.go
var navbar1 string

//go:embed examples/navbar1b/navbar.go
var navbar1b string

//go:embed examples/navbar2/navbar.go
var navbar2 string

//go:embed examples/navbar3/navbar.go
var navbar3 string

//go:embed examples/navbar4/navbar.go
var navbar4 string

//go:embed examples/navbar5/navbar.go
var navbar5 string

func IndexPage() g.Node {
	return Page(
		"gomponents, view components in pure Go",
		"View components in pure Go, that render to HTML5.",
		"https://www.gomponents.com",
		"/",

		Div(
			H1(g.Text("Tired of complex template languages?")),
			H2(g.Text("Try view components in pure Go.")),
			P(Class("lead"), g.Raw(`<em>gomponents</em> are view components in pure Go. They render to HTML 5, and make it easy for you to build reusable components. So you can focus on building your app instead of learning yet another templating language.`)),
			BashBlock("$ go get -u github.com/maragudk/gomponents"),
			P(
				A(Href("https://github.com/maragudk/gomponents"), g.Text("See gomponents on Github")),
			),

			H3(g.Text("Components are just functions")),
			P(g.Text("Have a look at this component. If you know HTML, you know what it does. Easy, right?")),
			CodeBlock(stripLines(navbar1, 7)),
			P(g.Text("Let's deduplicate a bit.")),
			CodeBlock(stripLines(navbar1b, 7)),

			H3(g.Text("Inline if")),
			P(g.Text("Sometimes you only want to show a component based on some condition. Enter "), Code(g.Text("If")), g.Text(":")),
			CodeBlock(stripLines(navbar2, 7)),

			H3(g.Text("Map data to components")),
			P(g.Text("What if you have data and want to map it to components? No problem.")),
			CodeBlock(stripLines(navbar3, 7)),

			H3(g.Text("Styling")),
			P(g.Text("Want to apply CSS classes based on state? Use the "), Code(g.Text("Classes")), g.Text(" helper.")),
			CodeBlock(stripLines(navbar4, 8)),

			H3(g.Text("Sometimes you just want HTML")),
			P(g.Text("Miss those "), Code(g.Text("<tags>")), g.Text(" or need to inject HTML from somewhere else? Use "), Code(g.Text("Raw")), g.Text(".")),
			CodeBlock(stripLines(navbar5, 8)),

			H3(g.Text("It's all just Go")),
			P(g.Text("Your editor helps you with auto-completion. It's type-safe. Nice formatting using gofmt. And all common HTML elements and attributes are included!")),

			H2(g.Text("Get started")),
			BashBlock("$ go get -u github.com/maragudk/gomponents"),

			H3(g.Text("A sample app")),
			P(
				g.Text("Sometimes there's nothing like seeing it in action. Try running this complete program and going to "),
				A(Href("http://localhost:8080"), g.Text("localhost:8080")),
				g.Text(" ."),
			),
			CodeBlock(simpleExample),

			P(
				g.Text("See also "),
				A(Href("https://github.com/maragudk/gomponents"), g.Text("the Github repository")),
				g.Text(" or "),
				A(Href("https://www.maragu.dk/blog/gomponents-declarative-view-components-in-go/"), g.Text("the blog post that started it all")),
				g.Text("."),
			),
		),
	)
}

func stripLines(s string, n int) string {
	lines := strings.Split(s, "\n")
	return strings.Join(lines[n:], "\n")
}
