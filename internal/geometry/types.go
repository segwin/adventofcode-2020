package geometry

type Number interface {
	Int() int64
	UInt() uint64
	Float() float64
}

type Int int64

func (i Int) Int() int64     { return int64(i) }
func (i Int) UInt() uint64   { return uint64(i) }
func (i Int) Float() float64 { return float64(i) }

type UInt uint64

func (u UInt) Int() int64     { return int64(u) }
func (u UInt) UInt() uint64   { return uint64(u) }
func (u UInt) Float() float64 { return float64(u) }

type Float float64

func (f Float) Int() int64     { return int64(f) }
func (f Float) UInt() uint64   { return uint64(f) }
func (f Float) Float() float64 { return float64(f) }
