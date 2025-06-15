package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {
	Convey("Testing Hello() function", t, func() {
		got := Hello()
		So(got, ShouldEqual, "Hello World")
	})
}

func TestHelloWithArgs(t *testing.T) {
	Convey("Testing HelloWithArgs() function", t, func() {
		got := HelloWithArgs("Chris", "")
		So(got, ShouldEqual, "Hello, Chris!")
	})
}

func TestHelloWithEmptyArgs(t *testing.T) {
	Convey("saying hello to people", t, func() {
		got := HelloWithArgs("Chris", "")
		So(got, ShouldEqual, "Hello, Chris!")
	})
	Convey("say 'Hello, World!' when supplied with an empty string", t, func() {
		got := HelloWithArgs("", "")
		So,(got, ShouldEqual, "Hello, World!")
	})
	Convey("in Spanish", t, func() {
		got := HelloWithArgs("Elodie", "Spanish")
		So(got, ShouldEqual, "Hola, Elodie!")
	})
	Convey("without specifying a language", t, func() {
		got := HelloWithArgs("Elodie", "")
		So(got, ShouldEqual, "Hello, Elodie!")
	})
	Convey("in French", t, func() {
		got := HelloWithArgs("Jules", "French")
		So(got, ShouldEqual, "Bonjour, Jules!")
	})
}
