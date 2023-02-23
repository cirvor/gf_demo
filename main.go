package main

import (
	_ "gf_demo/internal/logic"
	_ "gf_demo/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf_demo/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
