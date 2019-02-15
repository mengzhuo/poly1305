package poly1305

import (
	"math/big"
)

const TagSize = 16

var (
	p, r, outMod *big.Int
)

func init() {
	p, _ = new(big.Int).SetString("0x3fffffffffffffffffffffffffffffffb", 0) // 2^130 - 5
	r, _ = new(big.Int).SetString("0x0ffffffc0ffffffc0ffffffc0fffffff", 0)
	outMod, _ = new(big.Int).SetString("0xffffffffffffffffffffffffffffffff", 0) // 2^128
}

func Sum(out *[16]byte, m []byte, key *[32]byte) {
	sumGeneric(out, m, key)
}

func sumGeneric(out *[TagSize]byte, msg []byte, key *[32]byte) {

	h := new(big.Int)
	rpart := new(big.Int).SetBytes(key[:TagSize])
	rpart.And(rpart, r)

	for len(msg) >= TagSize {
		mi := new(big.Int).SetBytes(msg[:TagSize])
		h.Add(h, mi)
		h.Mul(h, rpart)
		h.Mod(h, p)
		msg = msg[TagSize:]
	}

	if len(msg) > 0 {
		var buf [TagSize]byte
		off := copy(buf[:], msg)
		buf[off] = 0x01
		mi := new(big.Int).SetBytes(buf[:])
		h.Add(h, mi)
		h.Mul(h, rpart)
		h.Mod(h, p)
	}
	h.Mod(h, p)

	if h.Cmp(p) >= 0 {
		h.Sub(h, p)
	}
	h.Mod(h, outMod)

	spart := new(big.Int).SetBytes(key[TagSize:])
	h.Add(h, spart)
	h.Mod(h, outMod)
	copy(out[:], h.Bytes())
}
