package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

func NewPermissionedContract(admin, address string, message json.RawMessage) PermissionedContract {
	return PermissionedContract{
		Address: address,
		Admin:   admin,
		Message: message,
	}
}

func (pc PermissionedContract) Validate() error {
	if strings.TrimSpace(pc.Address) == "" {
		return fmt.Errorf("invalid permissioned contract address")
	}

	if strings.TrimSpace(pc.Admin) == "" {
		return fmt.Errorf("invalid permissioned contract admin")
	}

	if !json.Valid(pc.Message) {
		return fmt.Errorf("invalid contract message json")
	}

	return nil
}
