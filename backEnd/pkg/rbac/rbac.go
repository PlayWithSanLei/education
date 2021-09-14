package rbac

import (
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/impact-eintr/WebKits/erbac"
	"github.com/impact-eintr/education/global"
)

var RootDir string
var once = new(sync.Once)

var R = struct{}{}

func init() {

	var err error
	once.Do(func() {
		inferRootDir()
		global.RBACSetting.CustomerInherFile = path.Join(RootDir, global.RBACSetting.CustomerInherFile)
		global.RBACSetting.CustomerRoleFile = path.Join(RootDir, global.RBACSetting.CustomerRoleFile)
		global.RBACSetting.DefaultRoleFile = path.Join(RootDir, global.RBACSetting.DefaultRoleFile)
		global.RBACSetting.DefaultInherFile = path.Join(RootDir, global.RBACSetting.DefaultInherFile)
		log.Println(global.RBACSetting.DefaultInherFile)
	})

	if exists(global.RBACSetting.CustomerInherFile) &&
		exists(global.RBACSetting.CustomerRoleFile) {

		global.Auth.RBAC, global.Auth.Permissions, err =
			erbac.BuildRBAC(global.RBACSetting.CustomerRoleFile,
				global.RBACSetting.CustomerInherFile)

	} else if exists(global.RBACSetting.DefaultInherFile) &&
		exists(global.RBACSetting.DefaultRoleFile) {

		global.Auth.RBAC, global.Auth.Permissions, err =
			erbac.BuildRBAC(global.RBACSetting.DefaultRoleFile,
				global.RBACSetting.DefaultInherFile)

		err = global.Auth.RBAC.SaveUserRBACWithSort(
			global.RBACSetting.CustomerRoleFile, global.RBACSetting.CustomerInherFile)
	} else {
		err = errors.New("权限配置文件不存在")
	}
	if err != nil {
		panic(err)
	}

}

func inferRootDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var infer func(d string) string
	infer = func(d string) string {
		// 确保当前目录下有配置文件
		if pathExists(d + "/confs") {
			return d
		}
		return infer(filepath.Dir(d))
	}

	RootDir = infer(cwd)
}

func pathExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
