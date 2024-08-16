package ws

import (
	"context"
	"testing"
	"time"
)

func Test_runListen(t *testing.T) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()
	final(timeout)
}
