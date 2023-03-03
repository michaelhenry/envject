package parser

import (
	"os"
	"testing"
)

func TestVenvInject(t *testing.T){

	os.Setenv("KEY1", "This is value 1")
	os.Setenv("KEY2", "This is value 2")
	os.Setenv("KEY3", "This is value 3")

	configText := `
	struct AppConfig {
		let key1 = "$KEY1"
		let key2 = "${KEY2}"
		let key3 = "$(KEY3)"
	}
	`
	want := `
	struct AppConfig {
		let key1 = "This is value 1"
		let key2 = "This is value 2"
		let key3 = "This is value 3"
	}
	`
	got := ReplaceEnvVariables(configText)

	if got != want {
			t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestVenvInjectWithInvalidFormatting(t *testing.T){

	os.Setenv("KEY1", "This is value 1")
	os.Setenv("KEY2", "This is value 2")
	os.Setenv("KEY3", "This is value 3")
	os.Setenv("KEY4", "This is value 4")
	os.Setenv("KEY5", "This is value 5")

	configText := `
	struct AppConfig {
		let key1 = "$KEY1"
		let key2 = "${KEY2)"
		let key3 = "$((KEY3))"
		let key4 = "$(KEY4}"
		let key5 = "$(KEY5"
	}
	`
	want := `
	struct AppConfig {
		let key1 = "This is value 1"
		let key2 = "${KEY2)"
		let key3 = "$((KEY3))"
		let key4 = "$(KEY4}"
		let key5 = "$(KEY5"
	}
	`
	got := ReplaceEnvVariables(configText)

	if got != want {
			t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestVenvInjectWithNoSetEnvValues(t *testing.T){

	os.Unsetenv("KEY1")
	os.Unsetenv("KEY2")
	os.Unsetenv("KEY3")

	configText := `
	struct AppConfig {
		let key1 = "$KEY1"
		let key2 = "${KEY2}"
		let key3 = "$(KEY3)"
	}
	`
	want := configText
	got := ReplaceEnvVariables(configText)

	if got != want {
			t.Errorf("got %q, wanted %q", got, want)
	}
}