package load

import (
	"testing"
	"time"
)

func TestRunXn(t *testing.T) {
	RunXn(time.Duration(time.Second*2), 20)
}
