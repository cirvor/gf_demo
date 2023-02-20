package main

import (
	_ "gf_demo/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gf_demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
