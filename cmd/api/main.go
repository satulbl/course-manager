package main

import (
	"fmt"
	"os"

	"course-manager/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "app stopped:", err)
		os.Exit(1)
	}
}
