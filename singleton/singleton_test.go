package singleton

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

func (s *mySuite) TestSingleton() {

	var AppContext *Singleton

	AppContext = NewSingleton()
	AppContext.Data["username"] = "joesample"

	s.Equal(AppContext.Data["username"], "joesample")

}
