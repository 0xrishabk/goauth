package generator

import (
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func InitializeNode() {
	sID := os.Getenv("SERVER_ID")
	n, err := strconv.ParseInt(sID, 10, 64)
	if err != nil {
		panic("Error while converting server id")
	}
	snowflake.Epoch = 1747044786000
	node, err = snowflake.NewNode(n)
	if err != nil {
		panic(err)
	}
}

func GenerateID() int64 {
	return node.Generate().Int64()
}
