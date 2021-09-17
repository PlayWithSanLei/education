package service

import (
	"encoding/json"
	"os"

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

	res = res[:0]
	dfs(tmpRBAC, "root")

	global.Auth.Lock()
	defer global.Auth.Unlock()
	global.Auth.RBAC = tmpRBAC
	global.Auth.Permissions = tmpPermissions

	return nil

}

func GetRBAC() (m []map[string][]string, err error) {
	global.Auth.RLock()
	defer global.Auth.RUnlock()

	f1, err := os.Open(global.RBACSetting.CustomerInherFile)
	defer f1.Close()
	if err != nil {
		return nil, inErr.ErrRBACNotFound
	}

	f2, err := os.Open(global.RBACSetting.CustomerRoleFile)
	defer f2.Close()
	if err != nil {
		return nil, inErr.ErrRBACNotFound
	}

	a := make(map[string][]string)
	if err = json.NewDecoder(f1).Decode(&a); err != nil {
		return nil, inErr.ErrRBACDecode
	}

	b := make(map[string][]string)
	if err = json.NewDecoder(f2).Decode(&b); err != nil {
		return nil, inErr.ErrRBACDecode
	}
	m = append(m, a, b)

	return m, nil

}

func QueryRBAC(id string) ([]string, error) {
	res = res[:0]
	dfs(global.Auth.RBAC, id)
	return res, nil
}

var res = []string{}

func dfs(r *erbac.RBAC, id string) {
	c, _ := r.GetParents(id)
	if len(c) != 0 {
		res = append(res, c...)
	}
	for _, v := range c {
		dfs(r, v)
	}
}
