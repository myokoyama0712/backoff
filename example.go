package backoff

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"
)

func ExampleRetry() {
	// An operation that may fail.
	operation := func() error {
		//return nil // or an error
		fmt.Printf("Time: %v\n", time.Now())
		return fmt.Errorf("test error")
	}

	err := Retry(operation, NewExponentialBackOff())
	if err != nil {
		// Handle error.
		fmt.Println("Failure...")
		return
	}

	// Operation is successful.
	fmt.Println("Success!")
}

func ExampleRetryContext() {
	// A context
	ctx := context.Background()

	// An operation that may fail.
	operation := func() error {
		return nil // or an error
	}

	b := WithContext(NewExponentialBackOff(), ctx)

	err := Retry(operation, b)
	if err != nil {
		// Handle error.
		return
	}

	// Operation is successful.
}

func ExampleTicker() {
	// An operation that may fail.
	operation := func() error {
		return nil // or an error
	}

	ticker := NewTicker(NewExponentialBackOff())

	var err error

	// Ticks will continue to arrive when the previous operation is still running,
	// so operations that take a while to fail could run in quick succession.
	for _ = range ticker.C {
		if err = operation(); err != nil {
			log.Println(err, "will retry...")
			continue
		}

		ticker.Stop()
		break
	}

	if err != nil {
		// Operation has failed.
		return
	}

	// Operation is successful.
	return
}
