# chessimage

`chessimage` is a [golang](https://golang.org) library for rendering a chess board PNG in specific state.


![](./docs/board_with_moves.png)

> FEN Notation: `rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1`
>
> `go run examples/board_with_moves.go | open -f -a /Applications/Preview.app/`

## Usage

Include in your go path.

```bash
go get github.com/cjsaylor/chessimage
```

Initialize the renderer with a FEN notation.

```go
board, _ := chessimage.NewRendererFromFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1")
```

Render the chess board to a png `image.Image` interface.

```go
f, _ := os.Create("board.png")
defer f.Close()
image, _ := board.Render(chessimage.Options{AssetPath: "./assets/")})
png.Encode(f, image)
```

## Options

You can define rendering options at render time:

```go
options := chessimage.Options{
	AssetPath: "./assets/"
}
renderer.Render(options)
```

#### AssetPath (**Required**)

Specify the path of the image assets for the individual pieces. Feel free to use the assets packaged in this repo, but be aware they are under CC license.

#### Resizer (`draw.CatmullRom`)

Change the algorhythm for asset resizing. Depending on your performance requirements, you may need to use a faster (but more lossy) resizing method (like `draw.NearestNeighbor`).

#### BoardSize (`512`)

Square board size in pixels

#### PieceRatio (`0.8`)

Size of the pieces relative as a percentage to the game board tile size. If the game board size is `800`, each board tile would be `100` pixels wide, and the pieces would render at `80` pixels with the default ratio.

#### LastMove (`nil`)

Set of tiles to highlight on the board.

```go
// Example highlight tiles E2 and E4
options := chessimage.Options{
    LastMove: &chessimage.LastMove{ chessImage.E2, chessImage.E4 }
}
```

## Todo

* Add support for `PGN` notation for rendering a board (similar to the `FEN` notation setup now)