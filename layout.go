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

func Page(title, description string, body g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       title,
		Description: description,
		Language:    "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href("/styles/app.css"), Type("text/css")),
			Link(Rel("stylesheet"), Href("/styles/highlightjs.min.css"), Type("text/css")),
			Script(Src("/scripts/highlight.min.js")),
			Script(g.Raw("hljs.highlightAll();")),
		},
		Body: []g.Node{
			Container(
				Prose(
					body,
				),
				Footer(Class("mt-32 prose prose-sm prose-indigo"),
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
func Container(children ...g.Node) g.Node {
	return Div(Class("max-w-7xl mx-auto p-4 sm:p-6 lg:p-8"), g.Group(children))
}

func Prose(children ...g.Node) g.Node {
	return Div(Class("prose lg:prose-lg xl:prose-xl prose-indigo"), g.Group(children))
}

func CodeBlock(text string) g.Node {
	return Pre(Code(Class("language-go"), g.Text(text)))
}

func BashBlock(text string) g.Node {
	return Pre(Code(Class("language-bash"), g.Text(text)))
}