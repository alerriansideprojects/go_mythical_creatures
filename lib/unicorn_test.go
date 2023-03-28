package gomythicalcreatures

import (
	// "fmt"
	"testing"
)

// Unicorn Tests
func TestUnicorn(t *testing.T) {
	assert := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %v\nWant: %v", got, want)
		}
	}

	t.Run("unicorn name", func (t *testing.T)  {
		unicorn := Unicorn{"Robert"}
		got := unicorn.name
		want := "Robert"

		assert(t, got, want)
	})

	t.Run("unicorn default color", func(t *testing.T) {
		unicorn := Unicorn{"Margaret"}
		got := unicorn.defaultColor()
		want := "Silver"

		assert(t, got, want)
	})

	t.Run("unicorn new color", func(t *testing.T) {
		unicorn := Unicorn{"Margaret"}
		got := unicorn.color("Purple")
		want := "Purple"

		assert(t, got, want)
	})

	t.Run("unicorn say", func(t *testing.T) {
		unicorn := Unicorn{"Margaret"}
		got := unicorn.say("Wonderful!")
		want := "**;* Wonderful! **;*"

		assert(t, got, want)
	})
}