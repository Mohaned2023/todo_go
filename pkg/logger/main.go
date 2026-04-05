package logger

import (
	"fmt"
	"os"
	"time"
)


func Info(msg any) {
	fmt.Printf(
		"%s\tINFO %#v\n",
		time.Now().UTC().Format(time.RFC3339),
		msg,
	)
}

func Err(msg any) {
	fmt.Fprintf(
		os.Stderr,
		"%s\tError %#v\n",
		time.Now().UTC().Format(time.RFC3339),
		msg,
	)
}
