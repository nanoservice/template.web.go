package main

import "github.com/nanoservice/libs-go/web"

func main() {
	web.Route("/", func(w *web.Response, r *web.Request) {
		w.Println("Hello, World! (from `--web --go` template)")
	})

	web.Start()
}
