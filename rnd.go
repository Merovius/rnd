// Package rnd provides a pre-seeded PRNG.
//
// It is meant for cases where a library wants to guarantee non-deterministic
// behavior. In that case, the API of math/rand provides several challenges:
//
//		1. *rand.Rand is not concurrency safe, so you need to wrap it into a
//		   mutex or use a sync.Pool. Locking is less efficient, as it holds the
//		   mutex for longer than required. sync.Pool means random number
//		   generation might allocate.
//		2. The default source (used by the top-level function) is concurrency
//		   safe, but not seeded. A library can not assume its user remembered
//		   to properly seed it. It can also not be sure a different library did
//		   not seed it with a bad seed (e.g. many people use the current time).
//
// This package works around that by using a concurrency safe and properly
// seeded shared source and not allowing to seed it manually.
package rnd

import (
	"hash/maphash"
	"math"
	"sync/atomic"

	"golang.org/x/exp/rand"
)

var (
	global = rand.New(new(rand.LockedSource))
	// calls counts the approximate number of calls to Source.Uint64, for
	// re-seeding occasionally.
	calls uint64
)

func init() {
	global.Seed(new(maphash.Hash).Sum64())
}

// reseed increments calls by n and perhaps re-seeds the global source.
func reseed(n int) {
	if atomic.AddUint64(&calls, uint64(n)) > math.MaxUint32 {
		// Concurrent calls might run into this branch. That's fine, re-seeding
		// happens very infrequently and isn't that expensive anyways.
		global.Seed(new(maphash.Hash).Sum64())
	}
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func Int63() int64 {
	defer reseed(1)
	return global.Int63()
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func Uint32() uint32 {
	defer reseed(1)
	return global.Uint32()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func Uint64() uint64 {
	defer reseed(1)
	return global.Uint64()
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func Int31() int32 {
	defer reseed(1)
	return global.Int31()
}

// Int returns a non-negative pseudo-random int.
func Int() int {
	defer reseed(1)
	return global.Int()
}

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Int63n(n int64) int64 {
	defer reseed(1)
	return global.Int63n(n)
}

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Int31n(n int32) int32 {
	defer reseed(1)
	return global.Int31n(n)
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Intn(n int) int {
	defer reseed(1)
	return global.Intn(n)
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func Float64() float64 {
	defer reseed(1)
	return global.Float64()
}

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func Float32() float32 {
	defer reseed(1)
	return global.Float32()
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func Perm(n int) []int {
	defer reseed(n)
	return global.Perm(n)
}

// Shuffle pseudo-randomizes the order of elements of s.
func Shuffle[T any](s []T) {
	global.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	reseed(len(s))
}

// Read generates len(p) random bytes and writes them into p. It always returns
// len(p) and a nil error.
func Read(p []byte) (n int, err error) {
	defer reseed(len(p) / 8)
	return global.Read(p)
}

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with
// standard normal distribution (mean = 0, stddev = 1).
// To produce a different normal distribution, callers can
// adjust the output using:
//
//  sample = NormFloat64() * desiredStdDev + desiredMean
//
func NormFloat64() float64 {
	defer reseed(1)
	return global.NormFloat64()
}

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1).
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
//
//  sample = ExpFloat64() / desiredRateParameter
//
func ExpFloat64() float64 {
	defer reseed(1)
	return global.ExpFloat64()
}
