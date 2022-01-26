package main
import "testing"

func TestGreeting(t *testing.T) {
	want := "Hello"
	got := Greeting()
	if got != want {
		t.Errorf("Greeting() = %q, want %q", got, want)
	}
}
