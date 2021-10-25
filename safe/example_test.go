package safe_test

import (
	"fmt"
	"github.com/v1990/go-utils/safe"
)

func ExamplePanicHandler() {
	var executor safe.Executor = safe.PanicHandler(func(recovered interface{}) {
		fmt.Println("recovered:", recovered)
	})

	executor.Execute(func() {
		fmt.Println("execute test1")
		panic("test1")
	})

	executor.Execute(func() {
		fmt.Println("execute test2")
		panic("test2")
	})

	// Output:
	// execute test1
	// recovered: test1
	// execute test2
	// recovered: test2

}

func ExamplePanicRetryHandler() {
	var executor = safe.PanicRetryHandler(func(recovered interface{}, attempt uint) (retry bool) {
		fmt.Println("recovered:", recovered, "attempt:", attempt)
		return attempt < 3
	})

	executor.Execute(func() {
		fmt.Println("execute test")
		panic("test")
	})

	// Output:
	// execute test
	// recovered: test attempt: 0
	// execute test
	// recovered: test attempt: 1
	// execute test
	// recovered: test attempt: 2
	// execute test
	// recovered: test attempt: 3
}
