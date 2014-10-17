package a

import (
	"strings"
	"testing"
	"runtime"
	"reflect"
)

func previousStack() string {
	buf := make([]byte, 100000)
	runtime.Stack(buf, false)
	return string(buf)
}

func ExpectEquals(t *testing.T, item interface{}, equals interface{}, msg string) {
	if !reflect.DeepEqual(item, equals) {
		t.Errorf("%s != %s: %s @ %s", item, equals, msg, previousStack())
	}
}

func ExpectNil(t *testing.T, item interface{}) {
	if item != nil {
		t.Errorf("Not nil: %s", item)
	}
}

func ExpectNotNil(t *testing.T, item interface{}) {
	if item == nil {
		t.Errorf("Nil: %s", item)
	}
}

func ExpectNotEquals(t *testing.T, item interface{}, equals interface{}, msg string) {
	if reflect.DeepEqual(item, equals) {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)
		t.Errorf("%s != %s: %s @ %s", item, equals, msg, previousStack())
	}
}

func ExpectContains(t *testing.T, item string, substr string, msg string) {
	if !strings.Contains(item, substr) {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)
		t.Errorf("%s DOES NOT CONTAIN %s: %s @ %s", item, substr, msg, previousStack())
	}
}
