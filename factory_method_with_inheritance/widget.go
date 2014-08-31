package factory_method

// WidgetType denotes which type of set is created. WidgetA or WidgetB
type WidgetType int

const (
	Widget_A = iota
	Widget_B
)

type WidgetInfo struct {
	id	 int
	name string
}

// widget is the domain type that this example manipulates
type Widget struct {
	WidgetInfo
}

func (wi WidgetInfo) SetInfo(id int, name string) {
	wi.id = id
	wi.name = name
}

func (wi WidgetInfo) GetInfo() WidgetInfo {
	return wi
}

// Interface is describing a Widget. Widgets are an unordered, unique list of values.
type Interface interface {
	SetInfo(id int, name string)	// common
	GetInfo() WidgetInfo			// common
	Add(items ...interface{})
	Remove(items ...interface{})
	IsEqual(a Interface) bool
	Size() int
	Each(func(interface{}) bool)
}


func (s WidgetType) String() string {
	switch s {
	case Widget_A:
		return "WidgetA"
	case Widget_B:
		return "WidgetB"
	}
	return ""
}

// helpful to not write everywhere struct{}{}
var keyExists = struct{}{}

// New creates and initalizes a new Widget interface. Its single parameter denotes the type of set to create.
func New(wt WidgetType, wi WidgetInfo) Interface {
	switch wt {
	case Widget_A:
		return newWidgetA(wi)
	case Widget_B:
		return newWidgetB(wi)
	}
	return nil
}
