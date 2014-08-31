package factory_method_with_reuse

// WidgetB's internal data set
type dataB struct {
	m map[interface{}]struct{}
}

// WidgetB's anonymous properties
type WidgetB struct {
	WidgetInfo
	dataA
}

// Create WidgetB data and set
func newWidgetB(wi WidgetInfo) *WidgetB {
	w := &WidgetB{}
	w.WidgetInfo = wi
	w.m = make(map[interface{}]struct{})
	var _ Widget_iface = w  // Enforce interface compliance
	return w
}

// Add any number of items to the set
func (d *dataB) Add(items ...interface{}) {
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		d.m[item] = NewEmptyStruct
	}
}

// Remove items from the set
func (d *dataB) Remove(items ...interface{}) {
	if len(items) == 0 {
		return
	}

	for _, item := range items {
		delete(d.m, item)
	}
}

// Size returns the number of items in a set
func (d *dataB) Size() int {
	return len(d.m)
}

// IsEqual tests whether d and a have the same items
func (d *WidgetB) IsEqual(a Widget_iface) bool {
	// return false if they are no the same size
	if sameSize := len(d.m) == a.Size(); !sameSize {
		return false
	}
	equal := true
	a.Each(func(item interface{}) bool {
		_, equal = d.m[item]
		return equal  // if false Each() will return
	})
	return equal
}

// IsEmpty indicates whether the Set is empty
func (d *dataB) IsEmpty() bool {
	return d.Size() == 0
}

// Iterate over each item, calling the function (f) for each item
// Continue until all items in the Set have been visited, or if the false is returned
func (d *dataB) Each(f func(item interface{}) bool) {
	for item := range d.m {
		if !f(item) {
			break
		}
	}
}

