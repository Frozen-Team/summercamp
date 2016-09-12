package forms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsStrongPass(t *testing.T) {
	Convey("short pass", t, func() {
		p := "!@#"
		res := isStrongPass(p)
		So(res, ShouldBeFalse)
	})

	Convey("weak pass: only numbers", t, func() {
		p := "12345"
		res := isStrongPass(p)
		So(res, ShouldBeFalse)
	})

	Convey("weak pass: only low letters", t, func() {
		p := "abcdef"
		res := isStrongPass(p)
		So(res, ShouldBeFalse)
	})

	Convey("weak pass: only up letters", t, func() {
		p := "ABCDF"
		res := isStrongPass(p)
		So(res, ShouldBeFalse)
	})

	Convey("weak pass: only up letters", t, func() {
		p := "ABCDF"
		res := isStrongPass(p)
		So(res, ShouldBeFalse)
	})

	Convey("strong pass: letters + lower", t, func() {
		p := "123ab"
		res := isStrongPass(p)
		So(res, ShouldBeTrue)
	})

	Convey("strong pass: letters + upper", t, func() {
		p := "123AB"
		res := isStrongPass(p)
		So(res, ShouldBeTrue)
	})

	Convey("strong pass: lower + upper", t, func() {
		p := "abcDE"
		res := isStrongPass(p)
		So(res, ShouldBeTrue)
	})

	Convey("strong pass: special", t, func() {
		p := "1234$"
		res := isStrongPass(p)
		So(res, ShouldBeTrue)
	})
}
