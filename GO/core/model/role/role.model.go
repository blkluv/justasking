package rolemodel

import uuid "github.com/satori/go.uuid"

// Role is a role and its permissions
type Role struct {
	RoleId               uuid.UUID
	RoleName             string
	PermissionName       string `json:"-"`
	PermissionValue      string `json:"-"`
	AccessBilling        bool
	SeeAllBoxes          bool
	ManageUsers          bool
	AccessAccountDetails bool
	CloseAccount         bool
	ManageOwners         bool
	EditAccountName      bool
}
