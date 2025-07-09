package bin

import (
	_ "embed"
	"os"
)

var BinaryPath string

func init() {
	var err error
	BinaryPath, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}
