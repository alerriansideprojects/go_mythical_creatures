package gomythicalcreatures

import "testing"

// Vampire Tests
func TestVampire(t *testing.T) {
	assert := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\nGot: %v\nWant: %v", got, want)
		}
	}

	t.Run("vampire name", func(t *testing.T) {
		vampire := Vampire{"Dracula"}
		got := vampire.name
		want := "Dracula"

		assert(t, got, want)
	})

	t.Run("vampire new name", func(t *testing.T) {
		vampire := Vampire{"Vladimir"}
		got := vampire.name
		want := "Vladimir"

		assert(t, got, want)
	})

	t.Run("vampire default pet", func(t *testing.T) {
		vampire := Vampire{"Ruthven"}
		got := vampire.defaultPet()
		want := "bat"

		assert(t, got, want)
	})
}
