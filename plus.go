package main

import (
	_ "embed"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents-heroicons/solid"
	. "github.com/maragudk/gomponents/html"
)

//go:embed examples/tailwindcss/tailwindcss.go
var tailwindcss string

//go:embed examples/heroicons/heroicons.go
var heroicons string

func PlusPage() g.Node {
	return Page(
		"gomponents +",
		"Let's build reusable components together. Add yours here!",
		"/plus/",

		Div(
			H1(g.Text("gomponents +")),
			P(Class("lead"), g.Text(`Let's build reusable components together. 🌟`)),
			P(g.Raw(`Have you built something with gomponents that other developers can use? <a href="https://github.com/maragudk/gomponents.com/issues/new">Create an issue on Github</a> and we can work together to add it here. 😎`)),

			Headline2(`TailwindCSS`, `tailwindcss`),
			P(g.Raw(`gomponents works great together with <a href="https://tailwindcss.com">TailwindCSS</a>. In fact, this page is built using gomponents and TailwindCSS! Check out <a href="https://github.com/maragudk/gomponents.com">the source of this page</a> or <a href="https://github.com/maragudk/gomponents-tailwind-example">a gomponents + TailwindCSS example project</a>.`)),
			H3(g.Text(`Example`)),
			CodeBlock(tailwindcss),

			Headline2(`Heroicons`, `heroicons`),
			P(g.Raw(`<a href="https://heroicons.com">Heroicons</a> are a collection of handcrafted SVG icons, by the makers of TailwindCSS. They work great together with gomponents! Check out <a href="https://github.com/maragudk/gomponents-heroicons">the gomponents-heroicons library on Github</a>.`)),
			H3(g.Text(`Example`)),
			CodeBlock(heroicons),
		),
	)
}

func Headline2(text, id string) g.Node {
	return H2(Class("inline-flex items-center"),
		g.Text(text),
		ID(id),
		A(Href("#"+id), solid.Link(Class("ml-2 h-5 w-5 lg:h-6 lg:w-6 xl:h-7 xl:w-7"))),
	)
}
