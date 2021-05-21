package types

// NewApplicationData allows to build a new ApplicationData instance
func NewApplicationData(application, username string) *ApplicationData {
	return &ApplicationData{
		Name:     application,
		Username: username,
	}
}

// NewVerificationData allows to build a new VerificationData instance
func NewVerificationData(method, value string) *VerificationData {
	return &VerificationData{
		Method: method,
		Value:  value,
	}
}
