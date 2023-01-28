package chessimage

import (
	_ "embed"
)

//go:embed "assets/pl.png"
var whitePawn []byte

//go:embed "assets/pd.png"
var blackPawn []byte

//go:embed "assets/bl.png"
var whiteBishop []byte

//go:embed "assets/bd.png"
var blackBishop []byte

//go:embed "assets/nl.png"
var whiteKnight []byte

//go:embed "assets/nd.png"
var blackKnight []byte

//go:embed "assets/rl.png"
var whiteRook []byte

//go:embed "assets/rd.png"
var blackRook []byte

//go:embed "assets/ql.png"
var whiteQueen []byte

//go:embed "assets/qd.png"
var blackQueen []byte

//go:embed "assets/kl.png"
var whiteKing []byte

//go:embed "assets/kd.png"
var blackKing []byte
