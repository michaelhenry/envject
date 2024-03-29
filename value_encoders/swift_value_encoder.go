package value_encoders

import (
	"fmt"
	"michaelhenry/envject/obfuscators"
	"michaelhenry/envject/utils"
)

type SwiftValueEncoder struct {
	Salt []byte
}

func (v *SwiftValueEncoder) Encode(value string) string {
	bytes := obfuscators.Obfuscate(value, v.Salt)
	return fmt.Sprintf("\\(deobfuscate([%v]))", utils.ConvertBytesToHexString(bytes))
}

func (v *SwiftValueEncoder) AdditionalCode() string {
	return fmt.Sprintf(`
// MARK: - Auto generated by envject (https://github.com/michaelhenry/envject)
private let salt: [UInt8] = [%s]
private func deobfuscate(_ bytes: [UInt8]) -> String { bytes.enumerated().reduce(into: "") { $0.append(Character(.init($1.element ^ salt[$1.offset %% salt.count])))}}`, utils.ConvertBytesToHexString(v.Salt))
}

func NewSwiftValueEncoder() *SwiftValueEncoder {
	bytes, _ := utils.GenerateRandomBytes(16)
	return &SwiftValueEncoder{Salt: bytes}
}