package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

func NewPermissionedContract(admin, address string, message json.RawMessage) PermissionedContract {
	return PermissionedContract{
		Address:  address,
		Admin:    admin,
		Messages: [][]byte{message},
	}
}

func (pc PermissionedContract) Validate() error {
	if strings.TrimSpace(pc.Address) == "" {
		return fmt.Errorf("invalid permissioned contract address")
	}

	if strings.TrimSpace(pc.Admin) == "" {
		return fmt.Errorf("invalid permissioned contract admin")
	}

	for _, message := range pc.Messages {
		if !json.Valid(message) {
			return fmt.Errorf("invalid contract message json")
		}
	}

	return nil
}

func (pc PermissionedContract) AddMessage(msg json.RawMessage) PermissionedContract {
	pc.Messages = append(pc.Messages, msg)
	return pc
}
