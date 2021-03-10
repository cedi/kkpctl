package utils

// IsOneOf probes if the needle is found in a haystack
// the runtime complexity is probably terrible since it's O(n), but who cares
func IsOneOf(needle interface{}, haystack ...interface{}) bool {
	for _, probe := range haystack {
		if needle == probe {
			return true
		}
	}

	return false
}
