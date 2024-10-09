package main

import (
	. "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

type PageProps struct {
	Title string
	Body  Node
}

func Page(title, description, baseURL, path string, body Node) Node {
	return c.HTML5(c.HTML5Props{
		Title:       title,
		Description: description,
		Language:    "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("/styles/app.css"), Type("text/css")),
			Link(Rel("stylesheet"), Href("/styles/prism.css"), Type("text/css")),
			Script(Src("https://cdn.usefathom.com/script.js"), DataAttr("site", "MOGDWGQV"), Defer()),
			Script(Src("/scripts/prism.js")),
			FavIcons(),
			OpenGraph(title, description, baseURL+"/images/logo.png", baseURL+path),
			TwitterCard(),
		},
		Body: []Node{
			Class("dark:bg-gray-900"),
			Navbar(path),
			Container(true,
				Img(Src("/images/logo.png"), Alt("Gopher logo"), Class("hidden h-64 w-auto float-right lg:block")),
				Prose(
					Img(Src("/images/logo.png"), Alt("Gopher logo"), Class("h-24 sm:h-32 w-auto float-right lg:hidden")),
					body,
				),
				Footer(Class("mt-32 prose dark:prose-invert prose-sm prose-indigo"),
					P(
						Text("made with ✨sparkles✨ by "),
						A(Href("https://www.maragu.dev"), Text("maragu")),
					),
				),
			),
		},
	})
}

// Container restricts the width of the children, centers, and adds some padding.
func Container(padY bool, children ...Node) Node {
	return Div(c.Classes{"max-w-7xl mx-auto px-4 sm:px-6 lg:px-8": true, "py-4 sm:py-6 lg:py-8": padY}, Group(children))
}

func Navbar(path string) Node {
	return Nav(Class("bg-gray-700 dark:bg-gray-800 mb-6"),
		Container(false,
			Div(Class("flex items-center space-x-4 sm:space-x-6 lg:space-x-8 h-16"),
				NavbarLink("/", "Home", path),
				NavbarLink("/plus/", "gomponents +", path),
			),
		),
	)
}

func NavbarLink(path, text, currentPath string) Node {
	return A(Href(path), Text(text),
		c.Classes{
			"text-sm font-medium focus:outline-none focus:text-white hover:text-white": true,
			"text-white":    path == currentPath,
			"text-gray-300": path != currentPath,
		},
	)
}

func Prose(children ...Node) Node {
	return Div(Class("prose dark:prose-invert lg:prose-lg xl:prose-xl prose-indigo"), Group(children))
}

func CodeBlock(text string) Node {
	return Pre(Code(Class("language-go"), Text(text)))
}

func BashBlock(text string) Node {
	return Pre(Code(Class("language-bash"), Text(text)))
}

func FavIcons() Node {
	return Group([]Node{
		Link(Rel("apple-touch-icon"), Attr("sizes", "180x180"), Href("/apple-touch-icon.png")),
		Link(Rel("icon"), Type("image/png"), Attr("sizes", "32x32"), Href("/favicon-32x32.png")),
		Link(Rel("icon"), Type("image/png"), Attr("sizes", "16x16"), Href("/favicon-16x16.png")),
		Link(Rel("manifest"), Href("/manifest.json")),
		Link(Rel("mask-icon"), Href("/safari-pinned-tab.svg"), Attr("color", "#000000")),
		Meta(Name("msapplication-TileColor"), Content("#ffffff")),
		Meta(Name("theme-color"), Content("#ffffff")),
	})
}

func OpenGraph(title, description, image, url string) Node {
	return Group([]Node{
		Meta(Attr("property", "og:type"), Content("website")),
		Meta(Attr("property", "og:title"), Content(title)),
		If(description != "", Meta(Attr("property", "og:description"), Content(description))),
		If(image != "", Meta(Attr("property", "og:image"), Content(image))),
		If(url != "", Meta(Attr("property", "og:url"), Content(url))),
	})
}

func TwitterCard() Node {
	return Group([]Node{
		Meta(Name("twitter:card"), Content("summary")),
		Meta(Name("twitter:site"), Content("@markusrgw")),
		Meta(Name("twitter:creator"), Content("@markusrgw")),
	})
}
