package main

import (
	_ "embed"

	"github.com/maragudk/gomponents-heroicons/v2/solid"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed examples/tailwindcss/tailwindcss.go
var tailwindcss string

//go:embed examples/bulma/bulma.go
var bulma string

//go:embed examples/heroicons/heroicons.go
var heroicons string

//go:embed examples/lucideicons/lucideicons.go
var lucideicons string

//go:embed examples/htmx/htmx.go
var htmx string

//go:embed examples/alpinejs/alpinejs.go
var alpinejs string

//go:embed examples/iconify/iconify.go
var iconify string

func PlusPage() Node {
	type section struct {
		title string
		id    string
		body  Node
	}
	sections := []section{
		{
			title: `TailwindCSS`, id: `tailwindcss`, body: Div(
				P(Raw(`gomponents works great together with <a href="https://tailwindcss.com">TailwindCSS</a>. In fact, this page is built using gomponents and TailwindCSS! Check out <a href="https://github.com/maragudk/gomponents.com">the source of this page</a>.`)),

				P(Raw(`There's also the <a href="https://github.com/maragudk/gomponents-starter-kit">gomponents starter kit</a>, which includes TailwindCSS.`)),

				H3(Text(`Example`)),
				CodeBlock(tailwindcss),

				H3(Text(`IntelliSense`)),
				P(Raw(`You can get TailwindCSS auto-completion directly in your Go files! Check the <a href="https://github.com/maragudk/gomponents-starter-kit">readme in the gomponents starter kit</a> for instructions.`)),
			),
		},
		{
			title: `Bulma`, id: `bulma`, body: Div(
				P(Raw(`gomponents and <a href="https://bulma.io">Bulma</a> also go well together. Check out <a href="https://willoma.github.io/bulma-gomponents/">bulma-gomponents</a> by <a href="https://github.com/willoma">willoma</a> for easy integration.`)),
				H3(Text(`Example`)),
				CodeBlock(bulma),
			),
		},
		{
			title: "HTMX", id: "htmx", body: Div(
				P(Raw(`<a href="https://htmx.org">HTMX</a> is a tiny Javascript library to give access to AJAX, websockets, and more, using HTML attributes. This fits perfectly with gomponents when you need that extra bit of client-side interactivity. Check out <a href="https://github.com/maragudk/gomponents-htmx">the gomponents-htmx library on Github</a>.`)),

				P(Raw(`There's also the <a href="https://github.com/maragudk/gomponents-starter-kit">gomponents starter kit</a>, which includes HTMX.`)),

				H3(Text(`Example`)),
				CodeBlock(htmx),
			),
		},
		{
			title: "AlpineJS", id: "alpinejs", body: Div(
				P(Raw(`<a href="https://alpinejs.dev">AlpineJS</a> is a minimal framework for composing JavaScript behavior in your markup. It's a great fit with gomponents for adding client-side interactivity to your components. Check out <a href="https://github.com/glsubri/gomponents-alpine">gomponents-alpine</a> on Github for integration with gomponents.`)),
				H3(Text(`Example`)),
				CodeBlock(alpinejs),
			),
		},
		{
			title: `Heroicons`, id: `heroicons`, body: Div(
				P(Raw(`<a href="https://heroicons.com">Heroicons</a> are a collection of handcrafted SVG icons, by the makers of TailwindCSS. They work great together with gomponents! Check out <a href="https://github.com/maragudk/gomponents-heroicons">the gomponents-heroicons library on Github</a>.`)),
				H3(Text(`Example`)),
				CodeBlock(heroicons),
			),
		},
		{
			title: "Iconify", id: "iconify", body: Div(
				P(Raw(`<a href="https://github.com/delaneyj/gomponents-iconify/">Iconify</a> is a large set of gomponents icons created by <a href="https://github.com/delaneyj">delaneyj</a>, from the <a href="https://iconify.design">Iconify project</a>.`)),
				H3(Text(`Example`)),
				CodeBlock(iconify),
			),
		},
		{
			title: `Lucide Icons`, id: `lucideicons`, body: Div(
				P(Raw(`<a href="https://lucide.dev/">Lucide Icons</a> are a collection of many beautiful & consistent icons made by the community. You can easily use them with gomponents, check out <a href="https://github.com/eduardolat/gomponents-lucide">the gomponents-lucide library on Github</a>.`)),
				H3(Text(`Example`)),
				CodeBlock(lucideicons),
			),
		},
	}

	return Page(
		"gomponents +",
		"Let's build reusable components together. Add yours here!",
		"https://www.gomponents.com",
		"/plus/",

		Div(
			H1(Text("gomponents +")),
			P(Class("lead"), Text(`Let's build reusable components together. 🌟`)),
			P(Raw(`Have you built gomponents that other developers can use? <a href="https://github.com/maragudk/gomponents.com/issues/new">Create an issue on Github</a> and let's add them here. 😎`)),

			P(Text("Jump to:")),
			Ul(
				Map(sections, func(s section) Node {
					return Li(A(Href("#"+s.id), Text(s.title)))
				}),
			),

			Map(sections, func(s section) Node {
				return Div(
					Headline2(s.title, s.id),
					s.body,
				)
			}),
		),
	)
}

func Headline2(text, id string) Node {
	return H2(Class("inline-flex items-center"),
		Text(text),
		ID(id),
		A(Href("#"+id), solid.Link(Class("ml-2 h-5 w-5 lg:h-6 lg:w-6 xl:h-7 xl:w-7"))),
	)
}
