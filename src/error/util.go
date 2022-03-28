package error

import (
	"fmt"
	"os"
)

func Fatal(err any) {
	fmt.Println(err)
	os.Exit(1)
}
