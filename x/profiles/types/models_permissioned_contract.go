package types

import (
	"fmt"
	"strings"
)

func NewPermissionedContract(admin, address string) PermissionedContract {
	return PermissionedContract{
		Address: address,
		Admin:   admin,
	}
}

func (pc PermissionedContract) Validate() error {
	if strings.TrimSpace(pc.Address) == "" {
		return fmt.Errorf("invalid permissioned contract address")
	}

	if strings.TrimSpace(pc.Admin) == "" {
		return fmt.Errorf("invalid permissioned contract admin")
	}

	return nil
}
