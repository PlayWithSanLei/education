package global

import (
	"sync"

	"github.com/impact-eintr/WebKits/erbac"
)

type Authority struct {
	sync.RWMutex
	RBAC        *erbac.RBAC
	Permissions erbac.Permissions
}

var (
	Auth = new(Authority)
)
