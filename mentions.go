package main

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func MentionsPage() g.Node {
	return Page(
		"Mentions",
		"Mentions of gomponents around the internet.",
		"https://www.gomponents.com",
		"/mentions/",
		Div(
			H1(g.Text("Mentions")),

			H2(g.Text(`Twitter`)),
			P(Class("lead"), g.Text("Shoutouts to gomponents around the internet.")),
			g.Raw(`<blockquote class="twitter-tweet" data-dnt="true" data-theme="light"><p lang="en" dir="ltr">This is so cool. It&#39;s the most interesting GO library I&#39;ve seen lately. <a href="https://t.co/8bq841yvtq">https://t.co/8bq841yvtq</a></p>&mdash; mpj_marshall (@mpj_marshall) <a href="https://twitter.com/mpj_marshall/status/1392298708860014595?ref_src=twsrc%5Etfw">May 12, 2021</a></blockquote> <script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>`),

			H2(g.Text(`Blogs`)),
			Ul(
				Li(A(Href("https://www.yellowduck.be/posts/building-view-components-with-go-and-tailwind-css/"), g.Text(`"Building view components with Go and Tailwind CSS" on yellowduck.be`))),
			),
		),
	)
}
