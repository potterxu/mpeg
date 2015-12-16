package util

import "io"

type BitReader interface {
	io.Reader
	ReadBit() (bool, error)
	PeekBit() (bool, error)
	Trash(uint) error
	IsByteAligned() bool
	ByteAlign() error
}

type BitReader32 interface {
	BitReader
	Read32(uint) (uint32, error)
	Peek32(uint) (uint32, error)
}

var NewBitReader func(io.Reader) BitReader32 = NewSimpleBitReader