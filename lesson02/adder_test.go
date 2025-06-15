package integers

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAdder(t * testing.T) {
	Convey("Testing Add Function", t, func() {
		got := Add(2, 2)
		So(got, ShouldEqual, 4)
	})
}

func ExampleAdd() {
	Convey("Example of Add Function", func() {
		got := Add(1, 2)
		So(got, ShouldEqual, 3)
	})
}
