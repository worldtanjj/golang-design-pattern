package idgenerator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerate(t *testing.T) {

	Convey("LogTraceIdGenerator test", t, func() {
		var i IdGenerator = LogTraceIdGenerator{}
		var _, err = i.Generate()
		So(err == nil, ShouldBeTrue)
	})
}

func TestGenerateRandomAlphameric(t *testing.T) {
	Convey("TestGenerateRandomAlphameric", t, func() {
		var actualSubstr = generateRandomAlphameric(8)
		So(len(actualSubstr) > 0, ShouldBeTrue)
	})
}
