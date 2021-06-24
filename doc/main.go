package doc

import (
	"os"

	"github.com/mattn/go-colorable"

	"github.com/tshihad/terraform-plugin-docs/pkg/cmd"
)

func Run() {
	name := "tfplugindocs"
	version := name + " Version " + version
	if commit != "" {
		version += " from commit " + commit
	}

	os.Exit(cmd.Run(
		name,
		version,
		[]string{"generate", "."},
		os.Stdin,
		colorable.NewColorableStdout(),
		colorable.NewColorableStderr(),
	))
}
