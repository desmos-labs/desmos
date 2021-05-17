package types

type VerificationData struct {
	Method string `obi:"validation_type"`
	Value  string `obi:"value"`
}

func NewVerificationData(method, value string) VerificationData {
	return VerificationData{
		Method: method,
		Value:  value,
	}
}
