package generator

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func InitializeNode(n int64) {
	var err error
	snowflake.Epoch = 1747044786000
	node, err = snowflake.NewNode(n)
	if err != nil {
		panic(err)
	}
}

func GenerateID() int64 {
	return node.Generate().Int64()
}
