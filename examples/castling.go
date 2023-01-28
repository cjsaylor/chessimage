package main

import (
	"image/png"
	"log"
	"os"

	"github.com/cjsaylor/chessimage"
)

func main() {
	board, err := chessimage.NewRendererFromFEN("r1bqkbnr/ppp2ppp/2np4/4p3/4P3/3B1N2/PPPP1PPP/RNBQ1RK1 b kq - 3 4")
	if err != nil {
		log.Fatalln(err)
	}
	board.SetLastMove(chessimage.LastMove{MoveType: chessimage.MoveTypeCastlingWK})
	image, err := board.Render(chessimage.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	png.Encode(os.Stdout, image)
}
