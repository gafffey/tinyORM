package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Error("test error")
	Errorf("test %s", "errorf")
	Info("test info")
	Infof("test %s", "infof")
}
