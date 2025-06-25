// kmpr is a helper to compare 2 values and returns boolean.
// It's tested to compare slices, struct, function, and interface.
package kmpr

import "fmt"

// Do is going to compare the pointer of arg1, and arg2 that passed as arguments.
func Do(arg1, arg2 any) bool {
	arg1Comparer := fmt.Sprintf("%v", arg1)
	arg2Comparer := fmt.Sprintf("%v", arg2)

	return arg1Comparer == arg2Comparer
}
