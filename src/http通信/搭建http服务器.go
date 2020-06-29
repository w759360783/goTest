package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func SHandlerError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := ioutil.ReadFile("F:/golandWoke/src/http通信/虎扑.html")
		SHandlerError(err, "ReadFile")
		_, err = writer.Write(bytes)
		SHandlerError(err, "Write")
	})
	err := http.ListenAndServe("0.0.0.0:8800", nil)
	SHandlerError(err, "ListenAndServe")
}
