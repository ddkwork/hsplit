package main

import (
	"hsplit"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("hsplit", func(w *unison.Window) {
		hsplit.New().Layout(w.Content())
	})
}
