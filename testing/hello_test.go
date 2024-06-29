package testing

import (
	"testing"

	"github.com/moabdelazem/gowithtests"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		input := gowithtests.Hello("Mohamed", "English")
		expected := "Hello, Mohamed"

		assertCorrectAnswer(t, input, expected)
	})

	t.Run("saying 'Hello, World' when empty string is supplied", func(t *testing.T) {
		input := gowithtests.Hello("", "")
		expected := "Hello, World"

		assertCorrectAnswer(t, input, expected)
	})

	t.Run("saying hello in french", func(t *testing.T) {
		input := gowithtests.Hello("Mohamed", "French")
		expected := "Bonjour, Mohamed"

		assertCorrectAnswer(t, input, expected)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		input := gowithtests.Hello("Mohamed", "Spanish")
		expected := "Hola, Mohamed"

		assertCorrectAnswer(t, input, expected)
	})
}

func assertCorrectAnswer(t testing.TB, input, expected string) {
	t.Helper()
	if input != expected {
		t.Errorf("got %q want %q", expected, input)
	}
}
