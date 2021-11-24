package prose

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}

}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}

}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	got := JoinWithCommas(list)
	want := "apple, orange, and pear"
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

//Replaces code above since its not duplicated
func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf(errorString(test.list, got, test.want))
		}
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
