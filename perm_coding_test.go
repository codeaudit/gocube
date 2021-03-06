package gocube

import "testing"

func BenchmarkAllPermutations8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		allPermutations(8)
	}
}

func BenchmarkEncodePermutationInPlace8(b *testing.B) {
	perms := allPermutations(8)
	for i := 0; i < b.N/len(perms); i++ {
		for _, p := range perms {
			var array [8]int
			copy(array[:], p)
			encodePermutationInPlace(array[:])
		}
	}
}

func TestEncodeChoose(t *testing.T) {
	answer := 0
	for i := 0; i < 12; i++ {
		for j := i + 1; j < 12; j++ {
			for k := j + 1; k < 12; k++ {
				for l := k + 1; l < 12; l++ {
					perm := make([]bool, 12)
					perm[i] = true
					perm[j] = true
					perm[k] = true
					perm[l] = true
					if answer != encodeChoice(perm) {
						t.Error("Expected", answer, "got", encodeChoice(perm))
					}
					answer++
				}
			}
		}
	}
}

func TestEncodePermutation(t *testing.T) {
	for length := 0; length < 8; length++ {
		testSet := allPermutations(length)
		for j := 0; j < len(testSet); j++ {
			if j != encodePermutation(testSet[j]) {
				t.Error("Encoded", testSet[j], "expected",
					j, "but got", encodePermutation(testSet[j]))
			}
		}
	}
}
