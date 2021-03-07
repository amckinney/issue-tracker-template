package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"go.uber.org/zap"
)

const (
	_localhost = "127.0.0.1"
	_defaultPort = 3000
)

var _defaultAddress = fmt.Sprintf("%s:%d", _localhost, _defaultPort)

func main () {
	logger, err := zap.NewDevelopment()
	if err != nil {
		exit(os.Stderr, err)
	}
	_, err = newHandler(logger)
	if err != nil {
		exit(os.Stderr, err)
	}
	exit(os.Stderr, http.ListenAndServe(_defaultAddress, nil))
}

func newHandler(logger *zap.Logger) (*http.Handler, error) {
	return nil, nil
}

// exit writes out the error, and exits with a
// non-zero code if the error is non-nil.
func exit(writer io.Writer, err error) {
	if err == nil {
		os.Exit(0)
	}
	io.WriteString(writer, err.Error())
	os.Exit(1)
}
