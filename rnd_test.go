package rnd

import "testing"

func Test(t *testing.T) {
	// We can't test a lot, as the behavior of the package is intentionally
	// non-deterministic. Also, the package only thinly wraps x/exp/rand
	// anyways. But we can at least test that we can call every function,
	// without panics.

	Int63()
	Uint32()
	Uint64()
	Int31()
	Int()
	Int63n(420)
	Int31n(420)
	Intn(420)
	Float64()
	Float32()
	Perm(420)
	Shuffle[int](nil)
	Shuffle(make([]int, 420))
	type myIntSlice []int
	Shuffle(make(myIntSlice, 420))
	if n, err := Read(nil); n != 0 || err != nil {
		t.Errorf("Read(<nil>) = %d, %v, want 0, <nil>", n, err)
	}
	if n, err := Read(make([]byte, 420)); n != 420 || err != nil {
		t.Errorf("Read(<nil>) = %d, %v, want 420, <nil>", n, err)
	}
	NormFloat64()
	ExpFloat64()
}
