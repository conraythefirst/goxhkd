package main

import (
	"flag"
	"net/rpc"

	"github.com/cupnoodles14/scratchpad/go/goxhkd/pkg/shared"
)

func main() {
	btn := flag.String("button", "", "specify a button")
	cmd := flag.String("command", "", "set command for the button")
	// clear := flag.Bool("clear", false, "clear the button")
	clearAll := flag.Bool("clearall", false, "clear all bindings")
	flag.Parse()

	conn := shared.Connection{
		Network: "unix",
		Address: shared.DefaultSocketAddr,
	}

	c, err := rpc.Dial(conn.Network, conn.Address)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	if btn != nil && cmd != nil {
		err = c.Call("GoRPC.BindCommand", shared.Binding{
			Cmd: *cmd,
			Btn: *btn,
		}, nil)
	} else if clearAll != nil {
		err = c.Call("GoRPC.UnbindAll", nil, nil)
	}

	if err != nil {
		panic(err)
	}
}
