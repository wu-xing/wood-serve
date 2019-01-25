package util

import (
	"github.com/dimiro1/banner"

	"bytes"
	"fmt"
	"github.com/mattn/go-colorable"
)

func LogBanner() {

	banner.Init(colorable.NewColorableStdout(), true, true, bytes.NewBufferString("\n\n██╗    ██╗ ██████╗  ██████╗ ██████╗ \n██║    ██║██╔═══██╗██╔═══██╗██╔══██╗\n██║ █╗ ██║██║   ██║██║   ██║██║  ██║\n██║███╗██║██║   ██║██║   ██║██║  ██║\n╚███╔███╔╝╚██████╔╝╚██████╔╝██████╔╝\n ╚══╝╚══╝  ╚═════╝  ╚═════╝ ╚═════╝ \n\n\n"))
	fmt.Println("")

}
