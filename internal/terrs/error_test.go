package terrs

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_Eq(t *testing.T) {
	fmt.Println(ErrUnknown.Eq(errors.New("unknown")))
	fmt.Println(ErrUnknown.Eq(ErrUsernameRegistered))
	fmt.Println(ErrUnknown.Eq(ErrUnknown))
}
