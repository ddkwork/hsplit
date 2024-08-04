package main

import (
	"hsplit"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("hsplit", func(w *unison.Window) {
		w.Content().AddChild(hsplit.New().Layout())
	})
}
