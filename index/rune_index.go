// Package index contains utilities for indexing the JMdict to support quick searching.
package index

import "jmdict"

// RuneIndex represents a map of runes to entry sets.
type RuneIndex map[rune]EntrySet

// Add adds an entry to the rune index for the specified rune.
func (i RuneIndex) Add(key rune, entry *jmdict.Entry) {
	mm, ok := i[key]
	if !ok {
		mm = make(EntrySet)
		i[key] = mm
	}
	mm.Add(entry)
}

// Get returns the entry set for a rune, or nil, false if the rune is not present.
func (i RuneIndex) Get(key rune) (EntrySet, bool) {
	mm, ok := i[key]
	if !ok {
		return nil, false
	}
	return mm, true
}

// GetIntersection returns the intersection of entry sets for the given runes.
func (i RuneIndex) GetIntersection(keys ...rune) EntrySet {
	is := make(EntrySet)
	initial := false
	for _, r := range keys {
		mm, ok := i.Get(r)
		if !ok {
			// If any runes are missing, the intersection will be empty.
			return nil
		}
		if !initial {
			is = mm
			initial = true
		} else {
			is = is.Intersect(mm)
			if len(is) == 0 {
				// All further intersections will be empty.
				return nil
			}
		}
	}
	return is
}
