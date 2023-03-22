package value_encoders

type ValueEncoder interface {
	Encode(value string) string
	AdditionalCode() string
}
