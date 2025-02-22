// the Æ chess engine
package main
// This file includes all the basic types that are used in this package


type PieceType int8; const (
	NoType PieceType = 0
	Pawn   PieceType = 1
	Rook   PieceType = 2
	Knight PieceType = 3
	Bishop PieceType = 4
	Queen  PieceType = 5
	King   PieceType = 6
)

type Color int8; const (
	NoColor Color = 0 << 3
	White   Color = 1 << 3
	Black   Color = 2 << 3
)

type Piece int8; const (
	NoPiece = Piece(NoColor) | Piece(NoType)

	WPawn   = Piece(White) | Piece(Pawn)
	WRook   = Piece(White) | Piece(Rook)
	WKnight = Piece(White) | Piece(Knight)
	WBishop = Piece(White) | Piece(Bishop)
	WQueen  = Piece(White) | Piece(Queen)
	WKing   = Piece(White) | Piece(King)

	BPawn   = Piece(Black) | Piece(Pawn)
	BRook   = Piece(Black) | Piece(Rook)
	BKnight = Piece(Black) | Piece(Knight)
	BBishop = Piece(Black) | Piece(Bishop)
	BQueen  = Piece(Black) | Piece(Queen)
	BKing   = Piece(Black) | Piece(King)
)

// to be inlined
func opposite_color(col Color) Color {
	switch col {
		case White:
			return Black
		case Black:
			return White
		default:
			return NoColor
	}
}

// to be inlined
func color_of(piece Piece) Color {
	return (3 << 3) & Color(piece)
}

// to be inlined
func type_of(piece Piece) PieceType {
	return (8 - 1) & PieceType(piece)
}

// to be inlined
func make_piece(pieceType PieceType, color Color) Piece {
	return Piece(color) | Piece(pieceType)
}

// to be inlined
func seperate(piece Piece) (PieceType, Color) {
	pieceType := type_of(piece)
	pieceColor := color_of(piece)
	return pieceType, pieceColor
}

type State struct {
	board [64]Piece
	turnToMove Color

	castlings Castlings
}

type Castlings int8; const (
	NoCastling       Castlings = 0
	WCastleKingside  Castlings = 1 << 0
	WCastleQueenside Castlings = 1 << 1
	BCastleKingside  Castlings = 1 << 2
	BCastleQueenside Castlings = 1 << 3
)


type Square struct {
	file int
	rank int
}

// to be inlined
func square_legal(sq Square) bool {
	return (sq.file < 8) && (sq.file >= 0) && (sq.rank < 8) && (sq.rank >= 0)
}

// to be inlined
func square_index(sq Square) int {
	return (sq.rank << 3) + sq.file
}

type moveFlags int8; const (
	StandardMove  moveFlags = 0
	CastlingMove  moveFlags = 1 << 0
	EnPassantMove moveFlags = 1 << 1
	PromotionMove moveFlags = 1 << 2
)

type Move struct {
	from Square
	to Square
	additionalFlags moveFlags
}

// to be inlined
func isMoveCastling(mv Move) bool {
	return mv.additionalFlags & CastlingMove != 0
}

// to be inlined
func isMoveEnPassant(mv Move) bool {
	return mv.additionalFlags & EnPassantMove != 0
}

// to be inlined
// This function could probably be omitted, by just saying "if the
// promoted pieceType is not NoPiece, it's a promotion"
func isMovePromotion(mv Move) bool {
	return mv.additionalFlags & PromotionMove != 0
}

// to be inlined
func getPromotedPieceType(mv Move) PieceType {
	return PieceType(mv.additionalFlags >> 3)
}

func makePromotion(pt PieceType) moveFlags {
	return moveFlags(pt << 3) | PromotionMove
}

// I might port the entire engine to C once I'm done with it in go.
// The reason I'm writing it in Go in the first place is since C has
// no native support for +Inf and -Inf and I hate the workarounds I
// have to do to achieve the same thing. I also expect the easy
// concurrency model of Go to come in handy later. What I miss from C
// the most is the lack of explicit inlining of functions. I hate
// the fact, that I am at the mercy of the Go compiler when it comes
// to which functions are inlined and which are not. I hope recognizes
// small functions well, but I have no idea.

// Well I've found out +Inf and -Inf DO in fact exist in C (i mean duh,
// it's C). Only thing remaining is the ease and the memory management
// (movelists would be a pain).