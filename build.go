package main

import (
	"log"
	"os"
	"path"
	"path/filepath"

	g "maragu.dev/gomponents"
)

func main() {
	os.Exit(build())
}

var pages = map[string]g.Node{
	"index.html":      IndexPage(),
	"plus/index.html": PlusPage(),
}

func build() int {
	for pagePath, n := range pages {
		fullPagePath := path.Join("docs", pagePath)

		if err := os.MkdirAll(filepath.Dir(fullPagePath), os.ModePerm); err != nil {
			log.Println("Error creating directory:", err)
			return 1
		}

		f, err := os.Create(fullPagePath)
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
