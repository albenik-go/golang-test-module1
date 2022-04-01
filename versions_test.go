package gotestmod1_test

import (
	"testing"

	gotestmod1 "github.com/albenik-go/golang-test-module1/v2"
)

func TestVersionString(t *testing.T) {
	if gotestmod1.VersionString() != gotestmod1.Version {
		t.FailNow()
	}
}
