package main

import (
	"hsplit"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("hsplit", func(w *unison.Window) {
		w.Content().AddChild(hsplit.New().Layout())
	})
}
