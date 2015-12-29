package video

import "github.com/32bitkid/huffman"
import "github.com/32bitkid/bitreader"

func (fp *VideoSequence) motion_vectors(s int, mb *Macroblock) error {

	f_code := fp.PictureCodingExtension.f_code

	mv_count, mv_format, dmv := mv_info(fp, mb)

	motion_vector_part := func(r, s, t int) error {
		code, err := decodeMotionCode(fp)
		if err != nil {
			val, _ := fp.Peek32(11)
			log.Printf("%011b", val)
			panic(err)
			return err
		}
		if f_code[s][t] != 1 && code != 0 {
			r_size := uint(f_code[s][t] - 1)

			_, err := fp.Read32(r_size)
			if err != nil {
				return err
			}
		}
		if dmv == 1 {
			panic("unsupported: dmv[]")
		}
		return nil
	}

	motion_vector := func(r, s int) error {
		err := motion_vector_part(r, s, 0)
		if err != nil {
			return err
		}
		return motion_vector_part(r, s, 1)
	}

	if mv_count == 1 {
		if mv_format == MotionVectorFormat_Field && dmv != 1 {
			panic("not implemented: field format")
		}
		return motion_vector(0, s)
	} else {
		panic("not implement: field format")
	}

	return nil
}

func decodeMotionCode(br bitreader.BitReader) (int, error) {
	val, err := motionCodeDecoder.Decode(br)
	if err != nil {
		return 0, err
	} else if code, ok := val.(int); ok {
		return code, nil
	} else {
		return 0, huffman.ErrMissingHuffmanValue
	}
}

var motionCodeDecoder = huffman.NewHuffmanDecoder(huffman.HuffmanTable{
	"0000 0011 001 ": -16,
	"0000 0011 011 ": -15,
	"0000 0011 101 ": -14,
	"0000 0011 111 ": -13,
	"0000 0100 001 ": -12,
	"0000 0100 011 ": -11,
	"0000 0100 11 ":  -10,
	"0000 0101 01 ":  -9,
	"0000 0101 11 ":  -8,
	"0000 0111 ":     -7,
	"0000 1001 ":     -6,
	"0000 1011 ":     -5,
	"0000 111 ":      -4,
	"0001 1 ":        -3,
	"0011 ":          -2,
	"011 ":           -1,
	"1":              0,
	"010":            1,
	"0010":           2,
	"0001 0":         3,
	"0000 110":       4,
	"0000 1010":      5,
	"0000 1000":      6,
	"0000 0110":      7,
	"0000 0101 10":   8,
	"0000 0101 00":   9,
	"0000 0100 10":   10,
	"0000 0100 010":  11,
	"0000 0100 000":  12,
	"0000 0011 110":  13,
	"0000 0011 100":  14,
	"0000 0011 010":  15,
	"0000 0011 000":  16,
})
