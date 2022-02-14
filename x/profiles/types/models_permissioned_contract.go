package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

func NewPermissionedContract(admin, address string, messages [][]byte) PermissionedContract {
	return PermissionedContract{
		Address:  address,
		Admin:    admin,
		Messages: messages,
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

func (pc PermissionedContract) GetMessage() (SudoMsg, error) {
	var msg SudoMsg
	err := json.Unmarshal(pc.Messages[0], &msg)
	if err != nil {
		return SudoMsg{}, err
	}
	return msg, nil
}
