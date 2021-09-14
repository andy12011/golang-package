package unique_id_factory

// use sonyflake
// go get github.com/sony/sonyflake
// 1 bit sign  & 39 bits Millisecond & 8 bits serial number & 16 bit machine num
import (
	"fmt"
	"golang-package/utils"
	"time"

	"github.com/sony/sonyflake"
)

const (
	SONY_START_TIME = "Thu, 02 Sep 2021 07:10:07 GMT" // When production must not modify this value
	SONY_MACHINE_ID uint16 = 1
)

var sony *sonyflake.Sonyflake

func init() {
	t ,err := time.Parse(time.RFC1123, SONY_START_TIME)
	fmt.Println(t.Format(utils.TIME_FORMAT_ONE))
	 if nil != err {
		 panic("GetSony time parse error")
	 }

	setting := sonyflake.Settings{
		StartTime: t,
		MachineID: func() (uint16, error) {
			return SONY_MACHINE_ID, nil
		},
		CheckMachineID: func(u uint16) bool {
			return true
		},
	}

	sony = sonyflake.NewSonyflake(setting)
}

func GetSonyID() (uint64, error) {
	return sony.NextID()
}