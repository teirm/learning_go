// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
package intset

import (
	"bytes"
	"fmt"
)

const WordSize int = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/WordSize, uint(x%WordSize)

	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/WordSize, (x % WordSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << uint(bit)
}

// UnionWith sets to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements in the set
func (s *IntSet) Len() int {
	if len(s.words) == 0 {
		return 0
	}

	var length int
	for _, word := range s.words {
		for i := 0; i < WordSize; i++ {
			length += int((word >> uint(i)) & 1)

		}
	}

	return length
}

// Remove deletes an entry from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/WordSize, uint(x%WordSize)

	if word < len(s.words) {
		s.words[word] = s.words[word] & (^(1 << bit))
	}
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] = 0
	}
}

// Copy returns a copy of the intset
func (s *IntSet) Copy() *IntSet {

	var t IntSet

	t.words = append(t.words, s.words...)

	return &t
}

// AddAll adds several numbers to the set s.
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// IntersectWith computes the intersection of two
// intsets
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, 0)
		}
	}
}

// DifferenceWith computes the difference of two
// intsets
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword
		} else {
			s.words = append(s.words, 0)
		}
	}
}

// SymmetricDifference computes the symmetric difference
// of two sets
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, 0^tword)
		}
	}
}

// Elems returns a slice of elements in set
func (s *IntSet) Elems() []int {
	var elements []int

	for i, word := range s.words {
		for j := 0; j < WordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elements = append(elements, 64*i+j)
			}
		}
	}
	return elements
}
