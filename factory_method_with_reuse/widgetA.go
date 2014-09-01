package factory_method_with_reuse

// WidgetA's internal data set
type dataA struct {
	m map[interface{}]struct{}
}

// WidgetA's anonymous properties
type WidgetA struct {
	WidgetInfo
	dataA
}

// Create WidgetA data and set
func newWidgetA(wi WidgetInfo) *WidgetA {
	w := &WidgetA{}
	w.WidgetInfo = wi
	w.m = make(map[interface{}]struct{})
	var _ Widget_iface = w  // Enforce interface compliance
	return w
}

// Add any number of items to the set
func (d *dataA) Add(items ...interface{}) {
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		d.m[item] = struct{}{}
	}
}

// Remove items from the set
func (d *dataA) Remove(items ...interface{}) {
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		delete(d.m, item)
	}
}

// Size returns the number of items in a set
func (d *dataA) Size() int {
	return len(d.m)
}

// IsEqual tests whether d and a have the same items
func (d *WidgetA) IsEqual(a Widget_iface) bool {
	if sameSize := len(d.m) == a.Size(); !sameSize {
		return false
	}
	if a.GetInfo().String() != d.GetInfo().String() {
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
func (d *dataA) IsEmpty() bool {
	return d.Size() == 0
}

// Iterate over each item, calling the function (f) for each item
// Continue until all items in the Set have been visited, or if the false is returned
func (d *dataA) Each(f func(item interface{}) bool) {
	for item := range d.m {
		if !f(item) {
			break
		}
	}
}
