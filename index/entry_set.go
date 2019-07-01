package index

import "github.com/ollie123/jmdict"

// EntrySet represents a set of pointers to jmdict.Entry.
type EntrySet map[*jmdict.Entry]struct{}

// Add adds e to the set.
func (s EntrySet) Add(e *jmdict.Entry) {
	s[e] = struct{}{}
}

// Contains returns true if the set contains e.
func (s EntrySet) Contains(e *jmdict.Entry) bool {
	_, ok := s[e]
	return ok
}

// Intersect calculates the set intersection with r.
func (s EntrySet) Intersect(r EntrySet) EntrySet {
	i := make(EntrySet)
	for e := range s {
		if r.Contains(e) {
			i.Add(e)
		}
	}
	return i
}

// Union calculates the set union with r.
func (s EntrySet) Union(r EntrySet) EntrySet {
	u := make(EntrySet)
	for e := range s {
		u.Add(e)
	}
	for e := range r {
		u.Add(e)
	}
	return u
}

// Difference calculates the set difference with r.
func (s EntrySet) Difference(r EntrySet) EntrySet {
	d := make(EntrySet)
	for e := range s {
		if !r.Contains(e) {
			d.Add(e)
		}
	}
	return d
}
