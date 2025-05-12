package generator

import "github.com/bwmarrin/snowflake"

var n *snowflake.Node

func InitializeSnowFlakeNode(node int64) {
	var err error
	snowflake.Epoch = 1747044786000
	n, err = snowflake.NewNode(node)
	if err != nil {
		panic(err)
	}
}

func GenerateUserID() int64 {
	return n.Generate().Int64()
}
