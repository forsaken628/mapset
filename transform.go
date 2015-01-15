package stringset

// Map returns the Set that results from applying f to each element of s.
func (s Set) Map(f func(string) string) Set {
	var out Set
	for k := range s {
		out.Add(f(k))
	}
	return out
}

// Select returns the subset of s for which f returns true.
func (s Set) Select(f func(string) bool) Set {
	var out Set
	for k := range s {
		if f(k) {
			out.Add(k)
		}
	}
	return out
}

// Partition returns two disjoint sets, yes containing the subset of s for
// which f returns true and no containing the subset for which f returns false.
func (s Set) Partition(f func(string) bool) (yes, no Set) {
	for k := range s {
		if f(k) {
			yes.Add(k)
		} else {
			no.Add(k)
		}
	}
	return
}

// Choose returns an element of s for which f returns true, if one exists.  The
// second result reports whether such an element was found.
// If f == nil, chooses an arbitrary element of s.
//
// Example:
//   re := regexp.MustCompile(`[a-z]\d+`)
//   s := stringset.New("a", "b15", "c9", "q").Choose(re.MatchString)
//   fmt.Println(s.Keys()) ⇒ ["b15", "c9"]
//
func (s Set) Choose(f func(string) bool) (string, bool) {
	for k := range s {
		if f == nil || f(k) {
			return k, true
		}
	}
	return "", false
}

// Pop removes and returns an element of s for which f returns true, if one
// exists (essentially Choose + Discard).  The second result reports whether
// such an element was found.  If f == nil, pops an arbitrary element of s.
func (s Set) Pop(f func(string) bool) (string, bool) {
	if v, ok := s.Choose(f); ok {
		delete(s, v)
		return v, true
	}
	return "", false
}
