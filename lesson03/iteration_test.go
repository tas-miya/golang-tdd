package iteration

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRepeat(t *testing.T) {
	Convey("testing Repeat()", t, func() {
		repeated := Repeat("a")
		So(repeated, ShouldEqual, "aaaaa")
	})
}

func TestCompareString(t *testing.T) {
	Convey("testing CompareStrings()", t, func() {
		compare := CompareStrings("aaaa", "bb")
		So(compare, ShouldEqual, -1)
	})
}

func TestCutPrefixString(t * testing.T) {
	Convey("testing CutPrefix()", t, func() {
		cut := CutPrefixString("hello", "hel")
		So(cut, ShouldEqual, true)
	})
	Convey("testing CutPrefix()", t, func() {
		cut := CutPrefixString("hello", "lo")
		So(cut, ShouldEqual, false)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}