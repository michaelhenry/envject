package value_encoders

type RawValueEncoder struct {}

func (v *RawValueEncoder) Encode(value string) string {
	return value
}

func (v *RawValueEncoder) AdditionalCode() string {
	return ""
}