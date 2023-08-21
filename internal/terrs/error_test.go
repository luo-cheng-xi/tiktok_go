package terrs

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestError_Eq(t *testing.T) {
	fmt.Println(ErrUnknown.Eq(errors.New("unknown")))
	fmt.Println(ErrUnknown.Eq(ErrUsernameRegistered))
	fmt.Println(ErrUnknown.Eq(ErrUnknown))

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println(now.Nanosecond())
	fmt.Println(time.UnixMilli(now.Unix() * 1000))
}
