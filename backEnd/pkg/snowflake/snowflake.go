package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/impact-eintr/education/global"
)

var node *snowflake.Node

var S = struct{}{}

func init() {
	var st time.Time
	st, err := time.Parse("2006-01-02", global.ServerSetting.StatTime)
	if err != nil {
		panic(err)
	}

	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(global.ServerSetting.MachineID)
	if err != nil {
		panic(err)
	}
}

func GenID() int64 {
	return node.Generate().Int64()
}
