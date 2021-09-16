package service

import (
	"github.com/impact-eintr/WebKits/erbac"
	"github.com/impact-eintr/education/global"
	"github.com/impact-eintr/education/internal/inErr"
)

func UpdateRBAC(newInher, newRole map[string][]string) error {
	tmpRBAC := erbac.NewRBAC()
	tmpPermissions := make(erbac.Permissions)

	for rid, pids := range newRole {
		role := erbac.NewStdRole(rid)
		for _, pid := range pids {
			_, ok := tmpPermissions[pid]
			if !ok {
				tmpPermissions[pid] = erbac.NewStdPermission(pid)
			}
			role.Assign(tmpPermissions[pid])
		}
		tmpRBAC.Add(role)
	}

	for rid, parents := range newInher {
		if err := tmpRBAC.SetParents(rid, parents); err != nil {
			return inErr.ErrInvalidInher
		}
	}
	// 保存权限文件
	err := tmpRBAC.SaveUserRBACWithSort(
		global.RBACSetting.CustomerRoleFile, global.RBACSetting.CustomerInherFile)
	if err != nil {
		return err
	}

	global.Auth.Lock()
	defer global.Auth.Unlock()
	global.Auth.RBAC = tmpRBAC
	global.Auth.Permissions = tmpPermissions

	return nil

}
