package poly1305

import (
	"math/big"
)

const TagSize = 16

var (
	p, rAnd *big.Int
)

func init() {
	p, _ = new(big.Int).SetString("0x3fffffffffffffffffffffffffffffffb", 0) // 2^130 - 5
	rAnd, _ = new(big.Int).SetString("0x0ffffffc0ffffffc0ffffffc0fffffff", 0)
}

func Sum(out *[16]byte, m []byte, key *[32]byte) {
	sumGeneric(out, m, key)
}

func changeEndian(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
}
func numToLeBytes(b *big.Int, out *[TagSize]byte) {
	bl := len(b.Bytes())
	copy(out[:], make([]byte, 16))
	if bl == 0 {
		return
	}
	offset := 0
	if bl > TagSize {
		offset = bl - TagSize
		bl = TagSize
	}
	copy(out[TagSize-bl:], b.Bytes()[offset:])
	changeEndian(out[:])
}

func sumGeneric(out *[TagSize]byte, msg []byte, key *[32]byte) {
	h := new(big.Int)
	changeEndian(key[:TagSize])
	rpart := new(big.Int).SetBytes(key[:TagSize])
	rpart.And(rpart, rAnd)

	changeEndian(key[TagSize:])
	spart := new(big.Int).SetBytes(key[TagSize:])
	var buf [TagSize + 1]byte
	mi := new(big.Int)

	for len(msg) >= TagSize {
		buf = [TagSize + 1]byte{}
		copy(buf[:], msg[:TagSize])
		buf[16] = 0x1
		changeEndian(buf[:])
		mi.SetBytes(buf[:])
		h.Add(h, mi)
		h.Mul(h, rpart)
		h.Mod(h, p)
		msg = msg[TagSize:]
	}

	if len(msg) > 0 {
		buf = [TagSize + 1]byte{}
		off := copy(buf[:], msg)
		buf[off] = 0x01
		changeEndian(buf[:])
		mi.SetBytes(buf[:])
		h.Add(h, mi)
		h.Mul(h, rpart)
		h.Mod(h, p)
	}
	h.Add(h, spart)
	numToLeBytes(h, out)
}
