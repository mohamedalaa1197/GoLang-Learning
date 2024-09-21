package hello

import "testing"

func TestSayHello(t *testing.T) {

	subTests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, World!",
		},
		{
			items:  []string{"Gopher"},
			result: "Hello, Gopher!",
		},
		{
			items:  []string{"Gopher", "and", "Golang"},
			result: "Hello, Gopher and Golang!",
		},
	}

	for _, st := range subTests {
		if s := Say(st.items); s != st.result {
			t.Errorf("Say(%v) = %s, want %s", st.items, s, st.result)
		}
	}
}
