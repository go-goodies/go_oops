package factory_method

// Provides a common set baseline for both threadsafe and non-ts Sets.
type dataB struct {
	m map[interface{}]struct{} // struct{} doesn't take up space
}

// SetNonTS defines a non-thread safe set dataB structure.
type WidgetB struct {
	WidgetInfo
	dataA
}

// NewNonTS creates and initializes a new non-threadsafe Set.
func newWidgetB(wi WidgetInfo) *WidgetB {
	w := &WidgetB{}
	w.WidgetInfo = wi
	w.m = make(map[interface{}]struct{})

	// Ensure interface compliance
	var _ Interface = w

	return w
}

// Add includes the specified items (one or more) to the set. The underlying
// Set s is modified. If passed nothing it silently returns.
func (d *dataB) Add(items ...interface{}) {
	if len(items) == 0 {
		return
	}

	for _, item := range items {
		d.m[item] = keyExists
	}
}

// Remove deletes the specified items from the set.  The underlying Set s is
// modified. If passed nothing it silently returns.
func (d *dataB) Remove(items ...interface{}) {
	if len(items) == 0 {
		return
	}

	for _, item := range items {
		delete(d.m, item)
	}
}

// Size returns the number of items in a set.
func (d *dataB) Size() int {
	return len(d.m)
}

// IsEqual tests whether d and a are the same in size and have the same items.
func (d *dataB) IsEqual(a Interface) bool {

	// return false if they are no the same size
	if sameSize := len(d.m) == a.Size(); !sameSize {
		return false
	}

	equal := true
	a.Each(func(item interface{}) bool {
		_, equal = d.m[item]
		return equal // if false, Each() will end
	})

	return equal
}

// IsEmpty reports whether the Set is empty.
func (d *dataB) IsEmpty() bool {
	return d.Size() == 0
}

// Each traverses the items in the Set, calling the provided function for each
// set member. Traversal will continue until all items in the Set have been
// visited, or if the closure returns false.
func (d *dataB) Each(f func(item interface{}) bool) {
	for item := range d.m {
		if !f(item) {
			break
		}
	}
}
