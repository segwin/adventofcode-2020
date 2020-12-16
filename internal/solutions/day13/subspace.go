package day13

import (
	"errors"
	"math/big"
)

var (
	ErrNoIntersectionSubspace = errors.New("no intersection exists for the given subspaces")
)

// Subspace in the natural numbers. It is defined by a coefficient (A) and an offset
// (B) such that the space comprises all numbers defined by:
//
//   A*n + B, where n=0,1,2,...
type Subspace struct {
	Coefficient *big.Int
	Offset      *big.Int
}

func NewSubspace(coefficient int64, offset int64) *Subspace {
	return &Subspace{
		Coefficient: big.NewInt(coefficient),
		Offset:      big.NewInt(offset),
	}
}

func (a *Subspace) Intersect(b *Subspace) (c *Subspace, err error) {
	// find the GCD using the extended Euclidean algorithm along with the Bezout
	// identity coefficients (x, y)
	//
	// a*x + b*y = gcd(a,b)
	periodA := big.NewInt(0).Set(a.Coefficient)
	periodB := big.NewInt(0).Set(b.Coefficient)

	x, y := big.NewInt(0), big.NewInt(0)
	gcd := big.NewInt(0).GCD(x, y, periodA, periodB)

	relativeOffset := big.NewInt(0).Sub(a.Offset, b.Offset)

	// ensure an intersection subspace actually exists, i.e. periodA & periodB will
	// eventually sync up with their given offsets
	if big.NewInt(0).Mod(relativeOffset, gcd).Int64() != 0 {
		return nil, ErrNoIntersectionSubspace
	}

	// now let's find the coefficient for the new subspace, which is the least common
	// multiple of the two:
	//   lcm = periodA * periodB / gcd
	periodC := big.NewInt(0).Mul(periodA, periodB)
	periodC.Div(periodC, gcd)

	// finally, let's get the offset for the new subspace
	//
	// goal: find the least common multiple of spaces A and B
	//
	// simple case, without an offset:
	//   x*periodA = y*periodB
	//
	// with offsets:
	//   m * periodA - offsetA = n * periodB - offsetB
	//   m * periodA - n * periodB = offsetA - offsetB
	//   m * periodA - n * periodB = relativeOffset             (1)
	//
	// using Bezout's identity as our 2nd equation:
	//   x * periodA + y * periodB = gcd                        (2)
	//
	// combining (1) and (2), dividing everything by (relativeOffset / gcd) gives us:
	//   m = x * relativeOffset / gcd                           (3)
	//   n = y * relativeOffset / gcd                           (4)
	m := big.NewInt(0).Div(relativeOffset, gcd)
	m.Mul(m, x)

	// since both A and B are at the same point when they intersect, we only need m
	// to compute it (note that we're only interested in the *first* offset, hence
	// the modulo):
	//
	//   offsetC = -(m * periodA - offsetA) % periodC
	offsetC := big.NewInt(0).Mul(m, periodA)
	offsetC.Sub(a.Offset, offsetC)
	offsetC.Mod(offsetC, periodC)

	return &Subspace{
		Coefficient: periodC,
		Offset:      offsetC,
	}, nil
}
