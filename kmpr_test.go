package kmpr

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	t.Run("Comparing function", func(t *testing.T) {
		func1 := func() {
			fmt.Println("First function")
		}
		func2 := func() {
			fmt.Println("Second function")
		}
		func3 := func1

		tableTest := []struct {
			arg1 any
			arg2 any
			want bool
		}{
			{func1, func1, true},
			{func1, func2, false},
			{func1, func3, true},
		}

		for i, tt := range tableTest {
			got := Do(tt.arg1, tt.arg2)
			if got != tt.want {
				t.Errorf("got %v want %v, on index: %d", got, tt.want, i)
			}
		}
	})

	t.Run("Comparing slice", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := []string{"1st", "2nd", "3rd", "4th"}
		slice3 := slice1

		tableTest := []struct {
			arg1 any
			arg2 any
			want bool
		}{
			{slice1, slice1, true},
			{slice1, slice2, false},
			{slice1, slice3, true},
		}

		for i, tt := range tableTest {
			got := Do(tt.arg1, tt.arg2)
			if got != tt.want {
				t.Errorf("got %v want %v, on index: %d", got, tt.want, i)
			}
		}
	})

	t.Run("Comparing struct", func(t *testing.T) {
		struct1 := struct {
			item string
		}{
			item: "dummy",
		}

		struct2 := struct {
			qty int
		}{
			qty: 1,
		}

		struct3 := struct1

		tableTest := []struct {
			arg1 any
			arg2 any
			want bool
		}{
			{struct1, struct1, true},
			{struct1, struct2, false},
			{struct1, struct3, true},
		}

		for i, tt := range tableTest {
			got := Do(tt.arg1, tt.arg2)
			if got != tt.want {
				t.Errorf("got %v want %v, on index: %d", got, tt.want, i)
			}
		}
	})
}

func BenchmarkDo(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []string{"1st", "2nd", "3rd", "4th"}

	for i := 0; i < b.N; i++ {
		Do(slice1, slice2)
	}
}

func ExampleDo() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []string{"1st", "2nd", "3rd", "4th"}

	fmt.Println(Do(slice1, slice2))

	// Output: false
}

