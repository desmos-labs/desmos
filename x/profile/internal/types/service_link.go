package types

import (
	"fmt"
	"strings"
)

// ServiceLink represents the data that have been used to connect a verifiable trusted service
type ServiceLink struct {
	Name       string `json:"service_name"`       // Name of the trusted service (e.g. keybase)
	Credential string `json:"service_credential"` // Credential to be used to verify the user (eg. the Keybase identity)
	Proof      string `json:"service_proof"`      // Proof to verify the user (e.g. the
}

// Validate implements validator
func (sl ServiceLink) Validate() error {
	if len(strings.TrimSpace(sl.Name)) == 0 {
		return fmt.Errorf("name of the trusted service cannot be empty or blank")
	}

	if len(strings.TrimSpace(sl.Credential)) == 0 {
		return fmt.Errorf("credential of %s service cannot be empty or blank", sl.Name)
	}

	if len(strings.TrimSpace(sl.Proof)) == 0 {
		return fmt.Errorf("%s service proof cannot be empty or blank", sl.Name)
	}

	return nil
}

// Equals allows to check whether the contents of sl are the same of other
func (sl ServiceLink) Equals(other ServiceLink) bool {
	return sl.Name == other.Name &&
		sl.Credential == other.Credential &&
		sl.Proof == other.Proof
}
