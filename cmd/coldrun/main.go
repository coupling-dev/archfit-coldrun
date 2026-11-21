package main

import (
	"fmt"
	"os"

	"github.com/coupling-dev/archfit-coldrun/internal/api"
)

func main() {
	srv := api.New()
	if err := srv.Put("hello", "world"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	v, _ := srv.Get("hello")
	fmt.Println(v)
}
