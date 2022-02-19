[![Go Reference](https://pkg.go.dev/badge/gonih.org/rnd.svg)](https://pkg.go.dev/gonih.org/rnd)

Package rnd provides a pre-seeded PRNG.

It is meant for cases where a library wants to guarantee non-deterministic
behavior. In that case, the API of `math/rand` provides several challenges:

 	1. `*rand.Rand` is not concurrency safe, so you need to wrap it into a
 	   mutex or use a `sync.Pool`. Locking is less efficient, as it holds the
 	   mutex for longer than required. `sync.Pool` means random number
 	   generation might allocate.
 	2. The default source (used by the top-level function) is concurrency
 	   safe, but not seeded. A library can not assume its user remembered
 	   to properly seed it. It can also not be sure a different library did
 	   not seed it with a bad seed (e.g. many people use the current time).

This package works around that by using a concurrency safe and properly
seeded shared source and not allowing to seed it manually.
