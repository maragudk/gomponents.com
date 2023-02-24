package main

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type PageProps struct {
	Title string
	Body  g.Node
}

func Page(title, description, baseURL, path string, body g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       title,
		Description: description,
		Language:    "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href("/styles/app.css"), Type("text/css")),
			Link(Rel("stylesheet"), Href("/styles/prism.css"), Type("text/css")),
			Script(Src("https://engaging-classical.gomponents.com/script.js"), DataAttr("site", "MOGDWGQV"), Defer()),
			Script(Src("/scripts/prism.js")),
			FavIcons(),
			OpenGraph(title, description, baseURL+"/images/logo.png", baseURL+path),
			TwitterCard(),
		},
		Body: []g.Node{
			Class("dark:bg-gray-900"),
			Navbar(path),
			Container(true,
				Img(Src("/images/logo.png"), Alt("Gopher logo"), Class("hidden h-64 w-auto float-right lg:block")),
				Prose(
					Img(Src("/images/logo.png"), Alt("Gopher logo"), Class("h-24 sm:h-32 w-auto float-right lg:hidden")),
					body,
				),
				Footer(Class("mt-32 prose dark:prose-light prose-sm prose-indigo"),
					P(
						g.Text("made in 🇩🇰 by "),
						A(Href("https://www.maragu.dk"), g.Text("maragu")),
					),
					P(
						g.Text("maker of "),
						A(Href("https://www.golang.dk"), g.Text("online Go courses")),
					),
				),
			),
		},
	})
}

// Container restricts the width of the children, centers, and adds some padding.
func Container(padY bool, children ...g.Node) g.Node {
	return Div(c.Classes{"max-w-7xl mx-auto px-4 sm:px-6 lg:px-8": true, "py-4 sm:py-6 lg:py-8": padY}, g.Group(children))
}

func Navbar(path string) g.Node {
	return Nav(Class("bg-gray-700 dark:bg-gray-800 mb-6"),
		Container(false,
			Div(Class("flex items-center space-x-4 sm:space-x-6 lg:space-x-8 h-16"),
				NavbarLink("/", "Home", path),
				NavbarLink("/plus/", "gomponents +", path),
				NavbarLink("/mentions/", "Mentions", path),
			),
		),
	)
}

func NavbarLink(path, text, currentPath string) g.Node {
	return A(Href(path), g.Text(text),
		c.Classes{
			"text-sm font-medium focus:outline-none focus:text-white hover:text-white": true,
			"text-white":    path == currentPath,
			"text-gray-300": path != currentPath,
		},
	)
}

func Prose(children ...g.Node) g.Node {
	return Div(Class("prose dark:prose-light prose-sm sm:prose sm:dark:prose-light sm:prose-indigo lg:prose-lg xl:prose-xl prose-indigo"), g.Group(children))
}

func CodeBlock(text string) g.Node {
	return Pre(Code(Class("language-go"), g.Text(text)))
}

func BashBlock(text string) g.Node {
	return Pre(Code(Class("language-bash"), g.Text(text)))
}

func FavIcons() g.Node {
	return g.Group([]g.Node{
		Link(Rel("apple-touch-icon"), g.Attr("sizes", "180x180"), Href("/apple-touch-icon.png")),
		Link(Rel("icon"), Type("image/png"), g.Attr("sizes", "32x32"), Href("/favicon-32x32.png")),
		Link(Rel("icon"), Type("image/png"), g.Attr("sizes", "16x16"), Href("/favicon-16x16.png")),
		Link(Rel("manifest"), Href("/manifest.json")),
		Link(Rel("mask-icon"), Href("/safari-pinned-tab.svg"), g.Attr("color", "#000000")),
		Meta(Name("msapplication-TileColor"), Content("#ffffff")),
		Meta(Name("theme-color"), Content("#ffffff")),
	})
}

func OpenGraph(title, description, image, url string) g.Node {
	return g.Group([]g.Node{
		Meta(g.Attr("property", "og:type"), Content("website")),
		Meta(g.Attr("property", "og:title"), Content(title)),
		g.If(description != "", Meta(g.Attr("property", "og:description"), Content(description))),
		g.If(image != "", Meta(g.Attr("property", "og:image"), Content(image))),
		g.If(url != "", Meta(g.Attr("property", "og:url"), Content(url))),
	})
}

func TwitterCard() g.Node {
	return g.Group([]g.Node{
		Meta(Name("twitter:card"), Content("summary")),
		Meta(Name("twitter:site"), Content("@markusrgw")),
		Meta(Name("twitter:creator"), Content("@markusrgw")),
	})
}
