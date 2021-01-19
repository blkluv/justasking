package rolerepo

import (
	"justasking/GO/core/model/role"
	"justasking/GO/core/startup/flight"

	uuid "github.com/satori/go.uuid"
)

// GetRolePermissions gets the permissions for a role
func GetRolePermissions(roleId uuid.UUID) ([]rolemodel.Role, error) {
	db := flight.Context(nil, nil).DB

	rolePermissions := []rolemodel.Role{}
	err := db.Raw(`SELECT r.id as role_id, r.name as role_name, p.name as permission_name, rp.permission_value
		FROM roles r JOIN role_permissions rp ON r.id = rp.role_id JOIN permissions p ON p.id = rp.permission_id
		WHERE r.id = ?`, roleId).Scan(&rolePermissions).Error

	return rolePermissions, err
}

// GetRolePermissionsByUserId gets permissions for a user by checking current_account
func GetRolePermissionsByUserId(userId uuid.UUID) ([]rolemodel.Role, error) {
	db := flight.Context(nil, nil).DB

	rolePermissions := []rolemodel.Role{}
	err := db.Raw(`SELECT r.id as role_id, r.name as role_name, p.name as permission_name, rp.permission_value
		FROM users u JOIN account_users au ON u.id = au.user_id JOIN roles r ON au.role_id = r.id JOIN role_permissions rp ON r.id = rp.role_id JOIN permissions p ON p.id = rp.permission_id JOIN account_users
		WHERE u.id = ? AND au.current_account = 1`, userId).Scan(&rolePermissions).Error

	return rolePermissions, err
}
