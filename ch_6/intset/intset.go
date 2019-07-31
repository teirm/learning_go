// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
package intset

import "bytes"

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)

	return word < len(s.words) && s.words[word]&(1<<bit) != 0
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
		for j := 0; j < 64; j++ {
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
	for word := range s.words {
		for i := 0; i < 64; i++ {
			length += (word >> i) & 1

		}
	}

	return length
}

// Remove deletes an entry from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)

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
