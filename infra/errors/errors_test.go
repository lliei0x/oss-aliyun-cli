package errors

import (
	"testing"
)

func TestError(t *testing.T) {
	err := New("user not exists")
	if err.Error() != "user not exists" {
		t.Error()
	}
	err = NewWithCode("user not exists", 1000)
	if err.Error() != "user not exists" {
		t.Error()
	}
	err = Wrap(err, 404)
	if err.Error() != "user not exists" {
		t.Error()
	}
	if e, ok := err.(DError); ok {
		if e.Code() != 404 {
			t.Error()
		}
		if e.String() != "404: user not exists" {
			t.Error()
		}
	}
}
