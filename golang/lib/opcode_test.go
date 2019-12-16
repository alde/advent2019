package lib
import (
	"fmt"
	"testing"
)

func Test_OpCode(t *testing.T) {
	testData := []struct {
		in []int
		out []int
	} {
		{ []int{1,0,0,0,99}, []int{2,0,0,0,99} },
		{ []int{2,3,0,3,99}, []int{2,3,0,6,99} },
		{ []int{2,4,4,5,99,0}, []int{2,4,4,5,99,9801} },
		{ []int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99} },
	}
	for _, td := range testData {
		t.Run(fmt.Sprintf("%+v", td.in), func(t *testing.T) {
			actual, _ := OpCode(td.in)
			validate(t, actual, td.out)
		})
	}
}

func validate(t *testing.T, actual []int, expected []int) {
	for idx, e := range actual {
		if expected[idx] != e {
			fmt.Printf("expected: %+v\nactual: %+v\n", expected, actual)
			t.Fail()
		}
	}
}