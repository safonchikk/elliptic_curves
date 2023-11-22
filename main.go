package main

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"math/rand"
)

var curve = elliptic.P224()

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

func RandBigInt(bits int) (dest big.Int) {
	ar := make([]byte, bits/8)
	for i := 0; i < len(ar); i++ {
		ar[i] = byte(rand.Intn(256))
	}
	dest.SetBytes(ar)
	return dest
}

func BasePointGGet() (point ECPoint) {
	return ECPointGen(curve.Params().Gx, curve.Params().Gy)
}

func ECPointGen(x, y *big.Int) (point ECPoint) {
	point.X = x
	point.Y = y
	return point
}

func IsOnCurveCheck(a ECPoint) (c bool) {
	return curve.IsOnCurve(a.X, a.Y)
}

func AddECPoints(a, b ECPoint) (c ECPoint) {
	c.X, c.Y = curve.Add(a.X, a.Y, b.X, b.Y)
	return c
}

func DoubleECPoints(a ECPoint) (c ECPoint) {
	c.X, c.Y = curve.Double(a.X, a.Y)
	return c
}

func ScalarMult(k big.Int, a ECPoint) (c ECPoint) {
	c.X, c.Y = curve.ScalarMult(a.X, a.Y, k.Bytes())
	return c
}

func ECPointToString(point ECPoint) (s string) {
	return fmt.Sprintf("x: %v\ny: %v\n", point.X, point.Y)
}

func StringToECPoint(s string) (point ECPoint) {

	return point
}
func PrintECPoint(point ECPoint) {
	fmt.Print(ECPointToString(point))
}

func IsEqual(p1, p2 ECPoint) (res bool) {
	return p1.X.Cmp(p2.X) == 0 && p1.Y.Cmp(p2.Y) == 0
}

func main() {
	G := BasePointGGet()

	k := RandBigInt(256)
	d := RandBigInt(256)

	H1 := ScalarMult(d, G)
	H3 := ScalarMult(k, G)

	result := IsEqual(H1, H3)
	fmt.Println(result)

	H2 := ScalarMult(k, H1)
	H4 := ScalarMult(d, H3)

	result = IsEqual(H2, H4)
	fmt.Println(result)

	fmt.Println(IsOnCurveCheck(H2))
}
