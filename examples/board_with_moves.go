// +build ignore

package main

import (
	"image/png"
	"log"
	"os"

	"github.com/cjsaylor/chessimage"
)

func main() {
	board, err := chessimage.NewRendererFromFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1")
	if err != nil {
		log.Fatalln(err)
	}
	image, err := board.Render(chessimage.Options{
		AssetPath: "./assets/",
		LastMove: &chessimage.LastMove{
			From: chessimage.E2,
			To:   chessimage.E4,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	png.Encode(os.Stdout, image)
}
