package metrics

import "testing"

func TestUnits(t *testing.T) {
	actual := Milliseconds
	expected := "milliseconds"

	if actual != expected {
		t.Errorf("wrong unit, got %s, want %s", actual, expected)
	}
}
