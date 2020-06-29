package spiderkit

import (
	"fmt"
	"os"
)

func HandlerErr(err error, when string) {
	if nil != err {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
