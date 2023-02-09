// A package containing computations related to prime numbers

package prime

//---------------------------------------------------------------
// Functions

// IsPrime function returns true if n is prime
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
