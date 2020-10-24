package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "appdoida",
	Short: "This is my first CLI using Go",
	Long: `This is my first app CLI using Go and cobra, I'm very happy to make work as I expect. 
It's a simple app that list/copy files between local filesystem and also send it to a S3.

It's very easy to use, it has three subcommands, ls, copy and s3.
You have to use the flags in order to specify which files you want to handle.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand()
	rootCmd.PersistentFlags().StringP("files", "f", "", "files to list/copy")
	rootCmd.SetHelpCommand(&cobra.Command{ //Remove help on subcommands available.
		Use:    "no-help",
		Hidden: true,
	})
}
