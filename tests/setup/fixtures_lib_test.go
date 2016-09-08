package setup

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFixturesSearcherFilter(t *testing.T) {
	FixturesPath = "tests/setup"
	defer func() {
		FixturesPath = "tests/fixtures"
	}()
	Convey("Should fetch fixtures", t, func() {
		Convey("Shoud fetch fixtures with A == asdasd", func() {
			c := GetFixture("test_lib").Filter("A", Equal, "asdasd").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldEqual, "asdasd")
			}
		})

		Convey("Shoud fetch fixtures with A != asdasd", func() {
			c := GetFixture("test_lib").Filter("A", NotEqual, "asdasd").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldNotEqual, "asdasd")
			}
		})

		Convey("Shoud fetch fixtures which starts with 'cc'", func() {
			c := GetFixture("test_lib").Filter("A", StartsWith, "cc").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldStartWith, "cc")
			}
		})
		Convey("Shoud fetch fixtures which not starts with 'cc'", func() {
			c := GetFixture("test_lib").Filter("A", NotStartsWith, "cc").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldNotStartWith, "cc")
			}
		})
		Convey("Shoud fetch fixtures which starts with 'aa'", func() {
			c := GetFixture("test_lib").Filter("A", EndsWith, "aa").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldEndWith, "aa")
			}
		})
		Convey("Shoud fetch fixtures which not ends with 'aa'", func() {
			c := GetFixture("test_lib").Filter("A", NotEndsWith, "aa").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldNotEndWith, "cc")
			}
		})
		Convey("Shoud fetch fixtures which contains 'wer'", func() {
			c := GetFixture("test_lib").Filter("A", Contains, "wer").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldContainSubstring, "wer")
			}
		})
		Convey("Shoud fetch fixtures which not contains 'wer'", func() {
			c := GetFixture("test_lib").Filter("A", NotContains, "wer").All()
			So(len(c), ShouldNotEqual, 0)
			for _, f := range c {
				So(f["A"], ShouldNotContainSubstring, "wer")
			}
		})
	})

}