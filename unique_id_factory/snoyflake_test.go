package unique_id_factory

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetSony(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 30000; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("stop")
					return
				default:
					id, err := GetID()
					if nil != err {
						fmt.Println(err.Error())
						cancel()
						return
					}
					fmt.Println(id)
				}
			}
		}(ctx)
	}

	time.Sleep(10 * time.Minute)
}
