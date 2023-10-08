package assertions

import "testing"

func AssertEquals(t *testing.T, toAssert interface{}, toCompare interface{}) {
	if toAssert != toCompare {
		t.Fatalf("AssertionError: %v did not equal %v", toAssert, toCompare)
	}
}
