package factory_method_with_registry

// maker_iface defines makeWidget and registerWidget methods
type maker_iface interface {
	makeWidget(w widget) widget_iface
	registerWidget(widget_iface)
}

// widget_iface defines the getInfo method
type widget_iface interface {
	getInfo() widget
}

// factory has the create method, used to create and store widgets
type factory struct {
}

// factory method lets a class defer instantiation (dependency injection)
// The concrete function used to create widgets is encapsulated
func (self *factory) create(factory maker_iface, w widget) widget_iface {
	widget := factory.makeWidget(w)
	factory.registerWidget(widget)
	return widget
}

// widget is the domain type that this example manipulates
type widget struct {
	name string
	count int
}
// getInfo is the the widget's only method
func (w *widget) getInfo() widget {
	return *w
}

// widgetFactory used to create and store widgets
type widgetFactory struct {
	*factory
	widgets []*widget
}

// makeWidget takes a pointer to the widgetFactory and returns the new widget
func (self *widgetFactory) makeWidget(w widget) widget_iface {
	return &widget{w.name, w.count}
}

// registerWidget takes a pointer to the widgetFactory and stores the new widget in the factory
// Note: this state may not be suitable for most functional programming paradigms
func (wf *widgetFactory) registerWidget(wi widget_iface) {
	widget := wi.getInfo()
	wf.widgets = append(wf.widgets, &widget)
}
