package factory_method_with_reuse

// WidgetB's internal data map
type dataB struct {
	m map[interface{}]struct{}
}

// WidgetB's anonymous properties
type WidgetB struct {
	WidgetInfo
	dataA
}

// Create WidgetB data (deferred creation of a widget)
func newWidgetB(wi WidgetInfo) *WidgetB {
	w := &WidgetB{}
	w.WidgetInfo = wi
	w.m = make(map[interface{}]struct{})
	var _ Widget_iface = w  // Enforce interface compliance
	return w
}

// Add any number of items to data map
func (d *dataB) Add(items ...interface{}) {
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		d.m[item] = struct{}{}
	}
}

// Remove items from the map
func (d *dataB) Remove(items ...interface{}) {
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		delete(d.m, item)
	}
}

// Size returns the number of items
func (d *dataB) Size() int {
	return len(d.m)
}

// IsEqual tests whether d and a have the same items
func (d *WidgetB) IsEqual(a Widget_iface) bool {
	if sameSize := len(d.m) == a.Size(); !sameSize {
		return false
	}
	if a.GetInfo().String() != d.GetInfo().String() {
		return false
	}
	return true
}

// IsEmpty indicates whether the map is empty
func (d *dataB) IsEmpty() bool {
	return d.Size() == 0
}
