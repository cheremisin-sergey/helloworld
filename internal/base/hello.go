package base

import (
	"fmt"

	"github.com/helloworld/internal/base/internal"
)

func Hello() {
	fmt.Println("hello")
	internal.Bye()
}
