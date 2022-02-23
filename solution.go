package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

// SidesNumType figure sides number
type SidesNumType int

// Constants of figure sides number
const (
	SidesCircle = iota
	_
	_
	SidesTriangle
	SidesSquare
)

// CalcSquare calculates square of some type of figures defined above
func CalcSquare(sideLen float64, sidesNum SidesNumType) float64 {
	switch (sidesNum) {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)

	case SidesTriangle:
		/*
			 |\
			 | \  a
		   x |  \
			 |   \
			 ------
			  a / 2

			x:
			a^2 = x^2 + (a/2)^2
			a^2 = x^2 + a^2/4
			a^2 - a^2/4 = x^2
			(4a^2-a^2)/4 = x^2
			3a^2/4 = x^2
			x = sqrt(3a^2/4)
			x = sqrt(3a^2)/2
			x = a*sqrt(3)/2


			Square:
			s = x * a/2
			s = a*sqrt(3)/2 * a/2
			s = a^2*sqrt(3)/4
		*/

		return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4.0

	case SidesSquare:
		return float64(sideLen) * float64(sideLen)
	}

	return float64(0)
}
