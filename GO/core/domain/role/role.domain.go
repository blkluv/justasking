package roledomain

import (
	"fmt"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/model/role"
	"justasking/GO/core/repo/role"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

var domainName = "RoleDomain"

// GetRolePermissions gets the permissions for a role
func GetRolePermissions(roleId uuid.UUID) (rolemodel.Role, *operationresult.OperationResult) {
	functionName := "GetRolePermissions"
	result := operationresult.New()
	rolePermissions := rolemodel.Role{}

	permissions, err := rolerepo.GetRolePermissions(roleId)
	if err != nil {
		msg := fmt.Sprintf("Unable to retrieve role permissions for role [%v]. Error: [%v]", roleId, err.Error())
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		rolePermissions = mapRolePermissions(permissions)
	}

	return rolePermissions, result
}

// GetRolePermissionsByUserId gets permissions for a user by checking current_account
func GetRolePermissionsByUserId(userId uuid.UUID) (rolemodel.Role, *operationresult.OperationResult) {
	functionName := "GetRolePermissionsByUserId"
	result := operationresult.New()
	rolePermissions := rolemodel.Role{}

	permissions, err := rolerepo.GetRolePermissionsByUserId(userId)
	if err != nil {
		msg := fmt.Sprintf("Unable to retrieve role permissions for user [%v]. Error: [%v]", userId, err.Error())
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		rolePermissions = mapRolePermissions(permissions)
	}

	return rolePermissions, result
}

func mapRolePermissions(rolePermissions []rolemodel.Role) rolemodel.Role {

	var features map[string]string
	features = make(map[string]string)

	for _, row := range rolePermissions {
		features[row.PermissionName] = row.PermissionValue
	}

	role := rolemodel.Role{}

	role.RoleId = rolePermissions[0].RoleId
	role.RoleName = rolePermissions[0].RoleName
	role.AccessAccountDetails, _ = strconv.ParseBool(features["AccessAccountDetails"])
	role.AccessBilling, _ = strconv.ParseBool(features["AccessBilling"])
	role.CloseAccount, _ = strconv.ParseBool(features["CloseAccount"])
	role.ManageUsers, _ = strconv.ParseBool(features["ManageUsers"])
	role.SeeAllBoxes, _ = strconv.ParseBool(features["SeeAllBoxes"])
	role.ManageOwners, _ = strconv.ParseBool(features["ManageOwners"])
	role.EditAccountName, _ = strconv.ParseBool(features["EditAccountName"])

	return role
}
