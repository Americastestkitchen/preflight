package main

import (
	"fmt"
	"github.com/americastestkitchen.com/preflight/ruby"
	"os"
)

func main() {
	errors := []string{}
	for _, file := range FindChanges() {
		if file.Language == Ruby {
			if msg, err := ruby.CheckSyntax(file.Path); err != nil {
				errors = append(errors, msg)
			}
		}
	}
	if len(errors) > 0 {
		for _, msg := range errors {
			fmt.Println(msg)
		}
		os.Exit(1)
	}
}
