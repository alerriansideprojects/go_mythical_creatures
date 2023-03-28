package gomythicalcreatures

import (
	// "fmt"
	"testing"
)

// Unicorn Tests
func TestUnicorn(t *testing.T) {
	assertHelper := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %v\nWant: %v", got, want)
		}
	}

	t.Run("unicorn name", func (t *testing.T)  {
		unicorn := Unicorn{"Robert"}
		got := unicorn.name
		want := "Robert"

		assertHelper(t, got, want)
	})

	t.Run("unicorn color", func(t *testing.T) {
		unicorn := Unicorn{"Margaret"}
		got := unicorn.color("Silver")
		want := "Silver"

		assertHelper(t, got, want)
	})

	t.Run("unicorn say", func(t *testing.T) {
		unicorn := Unicorn{"Margaret"}
		got := unicorn.say("Wonderful!")
		want := "**;* Wonderful! **;*"

		assertHelper(t, got, want)
	})
}