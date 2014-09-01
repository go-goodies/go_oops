package factory_method_with_reuse

import (
	"strconv"
)

// WidgetType indicates which type (WidgetA or WidgetB) is created
type WidgetType int

const (
	Widget_A = iota  // start at 0 and increment by 1
	Widget_B
)

// Widget is the domain type that this example manipulates
type Widget struct {
	WidgetInfo		// Anonymous field
}

// WidgetInfo is an Embedded type that contains widget data
type WidgetInfo struct {
	id	 int
	name string
}

// SetInfo is the WidgetInfo method used to set the widget's id and name
func (wi *WidgetInfo) SetInfo(id int, name string) {
	wi.id = id
	wi.name = name
}

// GetInfo is the WidgetInfo method used to return the WidgetInfo object
func (wi WidgetInfo) GetInfo() WidgetInfo {
	return wi
}

// GetInfo is the WidgetInfo method used to return the WidgetInfo object
func (wi WidgetInfo) String() string {
	return strconv.Itoa(wi.id) + "_" + wi.name
}

// Widget_iface is describing a Widget. Widgets are an unordered, unique list of values.
type Widget_iface interface {
	SetInfo(id int, name string)	// reuse
	GetInfo() WidgetInfo			// reuse
	Add(items ...interface{})
	Remove(items ...interface{})
	IsEqual(a Widget_iface) bool
	Size() int
	Each(func(interface{}) bool)
}

// String displays the string name of the specified widget
func (s WidgetType) String() string {
	switch s {
	case Widget_A:
		return "WidgetA"
	case Widget_B:
		return "WidgetB"
	}
	return ""
}

// Create a new Widget interface based on WidgetType and set WidgetInfo
func New(wt WidgetType, wi WidgetInfo) Widget_iface {
	switch wt {
	case Widget_A:
		return newWidgetA(wi)
	case Widget_B:
		return newWidgetB(wi)
	}
	return nil
}
