package model

import (
	"errors"
	"fmt"
	"testing"
	"tiktok/internal/terrs"
)

func TestNewErrorRsp(t *testing.T) {
	fmt.Printf("%#v\n", NewErrorRsp(terrs.ErrUsernameRegistered))
	fmt.Printf("%#v\n", NewErrorRsp(errors.New("an error")))
}
