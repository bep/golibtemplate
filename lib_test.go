package golibtemplate

import (
	"testing"

	"github.com/go-quicktest/qt"
)

func TestFoo(t *testing.T) {
	qt.Assert(t, qt.Equals(Foo(), "foo"))
}
