package service_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/magicsea/ganet/service"

	//"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type hello struct {
	Who string
}

func Example(t *testing.T) {
	fmt.Println("service_test Example pass")
	props := actor.FromInstance(&BaseServer{})
	pid := actor.Spawn(props)
	pid.Tell(&hello{Who: "Roger"})
	time.Sleep(1)
	fmt.Println("service_test Example pass")
	pid.GracefulStop()
}
