package profiling

import "testing"

func BenchmarkFactorial(b *testing.B) {
	// Run the function b.N times
	for i := 0; i < b.N; i++ {
		factorial(10) // Change the number to the desired input value
	}
}

// // Run benchmarks using: go test -bench=.
// func TestMain(m *testing.M) {
// 	// Perform any setup here if needed
// 	// ...

// 	// Run benchmarks and tests
// 	m.Run()
// }
