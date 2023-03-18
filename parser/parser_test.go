package parser

import (
	"michaelhenry/envject/obfuscators"
	"michaelhenry/envject/value_encoders"
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
	got := ReplaceEnvVariables(configText, "", &value_encoders.RawValueEncoder{})

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
	got := ReplaceEnvVariables(configText, "", &value_encoders.RawValueEncoder{})

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
	got := ReplaceEnvVariables(configText, "", &value_encoders.RawValueEncoder{})

	if got != want {
			t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestVenvInjectSwiftFormat(t *testing.T){

	os.Setenv("KEY1", "hello world")
	os.Setenv("KEY2", "another secret key")
	os.Setenv("KEY3", "321jdsj3244+=&&3298329894ahdha***@3232>>>ad")

	configText := `
enum Secrets {
	let key1 = "$KEY1"
	let key2 = "$KEY2"
	let key3 = "$KEY3"
    let key4 = "${IGNORE_ME}"
    let key5 = "${IGNORE_THIS}"
}`

	want := `
enum Secrets {
	let key1 = "\({ deobfuscate([0x1a, 0x04, 0x02, 0x08, 0x00, 0x4d, 0x28, 0x1c, 0x17, 0x0f, 0x16], salt: salt) }())"
	let key2 = "\({ deobfuscate([0x13, 0x0f, 0x01, 0x10, 0x07, 0x08, 0x2d, 0x53, 0x16, 0x06, 0x11, 0x17, 0x11, 0x06, 0x41, 0x05, 0x01, 0x16], salt: salt) }())"
	let key3 = "\({ deobfuscate([0x41, 0x53, 0x5f, 0x0e, 0x0b, 0x1e, 0x35, 0x40, 0x57, 0x57, 0x46, 0x4e, 0x49, 0x54, 0x47, 0x5d, 0x56, 0x56, 0x55, 0x6c, 0x41, 0x5c, 0x5b, 0x4b, 0x51, 0x15, 0x1a, 0x05, 0x06, 0x05, 0x45, 0x47, 0x75, 0x33, 0x56, 0x51, 0x41, 0x57, 0x4a, 0x4c, 0x5f, 0x0f, 0x00], salt: salt) }())"
    let key4 = "${IGNORE_ME}"
    let key5 = "${IGNORE_THIS}"
}
// MARK: - Auto generated by envject (https://github.com/michaelhenry/envject)
private let salt: [UInt8] = [0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74]
private func deobfuscate(_ bytes: [UInt8], salt: [UInt8]) -> String { bytes.enumerated().reduce(into: "") { $0.append(Character(UnicodeScalar($1.element ^ salt[$1.offset % salt.count])))}}`

	got := ReplaceEnvVariables(configText, "^IGNORE_", &value_encoders.SwiftValueEncoder{Salt: []byte("random_secret")})

	if got != want {
			t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeobfuscate(t *testing.T){
	salt := []byte("random_secret")
	got, _ := obfuscators.Deobfuscate([]byte{0x1a, 0x04, 0x02, 0x08, 0x00, 0x4d, 0x28, 0x1c, 0x17, 0x0f, 0x16}, salt)
	want := "hello world"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}