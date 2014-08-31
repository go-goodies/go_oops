package factory_method_with_registry

import (
	"testing"
	"github.com/remogatto/prettytest"
)

type mySuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	prettytest.Run(t, new(mySuite))
}

func (s *mySuite) TestFactoryMethod() {

	expected_widgets := []widget{
		widget{"A", 100},
		widget{"B", 250},
		widget{"C", 300}}

	factory := &widgetFactory{}
	widgets := []widget_iface{
		factory.create(factory, widget{"A", 100}),
		factory.create(factory, widget{"B", 250}),
		factory.create(factory, widget{"C", 300}),
	}

	for i, widget := range widgets {
		s.Equal(widget.getInfo(), expected_widgets[i])
	}
}
