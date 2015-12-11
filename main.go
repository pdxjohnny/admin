package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/admin/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "admin"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
