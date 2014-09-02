package factory_method_with_reuse

// WidgetA's internal data map
type dataA struct {
	m map[interface{}]struct{}
}

// WidgetA's anonymous properties
type WidgetA struct {
	WidgetInfo
	dataA
}

// Create WidgetA data map
func newWidgetA(wi WidgetInfo) *WidgetA {
	w := &WidgetA{}
	w.WidgetInfo = wi
	w.m = make(map[interface{}]struct{})
	var _ Widget_iface = w  // Enforce interface compliance
	return w
}

// Add any number of items to the map
func (d *dataA) Add(items ...interface{}) {
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		d.m[item] = struct{}{}
	}
}

// Remove items from the data map
func (d *dataA) Remove(items ...interface{}) {
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		delete(d.m, item)
	}
}

// Size returns the number of items
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
	return true
}

// IsEmpty indicates whether the map is empty
func (d *dataA) IsEmpty() bool {
	return d.Size() == 0
}
