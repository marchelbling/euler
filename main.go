package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/marchelbling/euler/solutions"
)

func AnswerFor(fn func() int64) {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	start := time.Now()
	res := fn()
	fmt.Printf("Answer for %s: %d (took %s)\n", name, res, time.Since(start))
}

func main() {
	AnswerFor(solutions.Problem14)
}
