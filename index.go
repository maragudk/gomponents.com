package main

import (
	_ "embed"
	"strings"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

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

func IndexPage() Node {
	return Page(
		"gomponents, HTML components in pure Go",
		"HTML components in pure Go, that render to HTML 5.",
		"https://www.gomponents.com",
		"/",

		Div(
			H1(Text("Tired of complex template languages?")),
			H2(Text("Try HTML components in pure Go.")),
			P(Class("lead"), Raw(`<em>gomponents</em> are HTML components in pure Go. They render to HTML 5, and make it easy for you to build reusable components. So you can focus on building your app instead of learning yet another templating language.`)),
			BashBlock("$ go get maragu.dev/gomponents"),
			P(
				A(Href("https://github.com/maragudk/gomponents"), Text("See gomponents on Github")),
			),

			P(Raw(`Does your company depend on this project? <a href="mailto:markus@maragu.dk?Subject=Supporting%20your%20project">Contact me at markus@maragu.dk</a> to discuss options for a one-time or recurring invoice to ensure its continued thriving.`)),

			H3(Text("Video introduction")),

			P(Raw(`Into video? See this lightning talk from GopherCon 2021.`)),

			Raw(`<iframe class="w-full h-80 sm:h-96" src="https://www.youtube-nocookie.com/embed/5XiewS8ZbH8?start=570" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`),

			H3(Text("Components are just functions")),
			P(Text("Have a look at this component. If you know HTML, you know what it does. Easy, right?")),
			CodeBlock(stripLines(navbar1, 2)),
			P(Text("Let's deduplicate a bit.")),
			CodeBlock(stripLines(navbar1b, 2)),

			H3(Text("Inline if")),
			P(Text("Sometimes you only want to show a component based on some condition. Enter "), Code(Text("If")), Text(":")),
			CodeBlock(stripLines(navbar2, 2)),

			P(Raw(`PS: There's also <code>Iff</code>, which takes a callback function instead, to avoid those pesky nil pointer errors.`)),

			H3(Text("Map data to components")),
			P(Text("What if you have data and want to map it to components? No problem.")),
			CodeBlock(stripLines(navbar3, 2)),

			H3(Text("Styling")),
			P(Text("Want to apply CSS classes based on state? Use the "), Code(Text("Classes")), Text(" helper.")),
			CodeBlock(stripLines(navbar4, 2)),

			H3(Text("Sometimes you just want HTML")),
			P(Text("Miss those "), Code(Text("<tags>")), Text(" or need to inject HTML from somewhere else? Use "), Code(Text("Raw")), Text(".")),
			CodeBlock(stripLines(navbar5, 2)),

			H3(Text("It's all just Go")),
			P(Text("Your editor helps you with auto-completion. It's type-safe. Nice formatting using gofmt. You can even use the debugger. And all common HTML elements and attributes are included!")),

			H2(Text("Get started")),
			BashBlock("$ go get maragu.dev/gomponents"),

			H3(Text("Examples & templates")),

			P(Raw(`<a href="https://github.com/maragudk/gomponents/tree/main/internal/examples/app">There’s an example app inside the gomponents repository</a>. It’s a simple web server that serves two HTML pages using gomponents and TailwindCSS.`)),

			P(Raw(`There’s also the <a href="https://github.com/maragudk/gomponents-starter-kit">gomponents starter kit</a>, a full template repository for building a web app with gomponents, TailwindCSS, and HTMX.`)),

			H3(Text(`Online HTML-to-gomponents converters`)),

			P(Raw(`Need to convert HTML or an SVG into Gomponent calls? Our community has your back:`)),

			Ul(
				Li(Raw(`<a href="https://github.com/PiotrKowalski">Piotr Kowalski's</a> original <a href="https://htg.piotrkowalski.me">online HTML-to-gomponents</a> converter tool!`)),
				Li(Raw(`<a href="https://github.com/traherom">Ryan Morehart's</a> <a href="https://gomponents.morehart.dev/">alternative</a>`)),
			),

			H3(Text(`The end`)),

			P(
				Text("See also "),
				A(Href("https://github.com/maragudk/gomponents"), Text("the Github repository")),
				Text(" or "),
				A(Href("https://www.maragu.dk/blog/gomponents-declarative-view-components-in-go/"), Text("the blog post that started it all")),
				Text("."),
			),
			P(
				Text("PS: "), A(Href("https://maragu.gumroad.com/l/gomponents"), Text("Buy the Official gomponents <marquee> Element here! 😁")),
			),
		),
	)
}

func stripLines(s string, n int) string {
	lines := strings.Split(s, "\n")
	return strings.Join(lines[n:], "\n")
}
