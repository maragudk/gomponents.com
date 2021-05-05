package main

import (
	"log"
	"os"
	"path"

	g "github.com/maragudk/gomponents"
)

func main() {
	os.Exit(build())
}

var pages = map[string]g.Node{
	"index.html": IndexPage(),
}

func build() int {
	for pagePath, n := range pages {
		f, err := os.Create(path.Join("docs", pagePath))
		if err != nil {
			log.Println("Error creating file:", err)
			return 1
		}
		if err := n.Render(f); err != nil {
			log.Println("Error writing page:", err)
			return 1
		}
		_ = f.Close()
	}
	return 0
}
