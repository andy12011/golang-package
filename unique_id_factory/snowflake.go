package unique_id_factory

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

const (
	SNOW_START_TIME = "Thu, 02 Sep 2021 07:10:07 GMT" // When production must not modify this value
	SNOW_MACHINE_ID int64 = 1
)

var snow *snowflake.Node

func init() {
	node, err := snowflake.NewNode(SNOW_MACHINE_ID)

	if err != nil {
		fmt.Println(err)
		return
	}

	snow = node
}

func GetSnowID() int64 {
	id := snow.Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id.Int64())
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	return id.Int64()
}
