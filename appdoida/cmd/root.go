package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	//"strings"
)

//var CopyFiles string

// rootCmd represents the base command when called without any subcommands without Run just print the app function.
var rootCmd = &cobra.Command{
	Use:   "appdoida",
	Short: "Essa app doida vai tentar copiar arquivos tanto local qto para um S3.",
	Long: `Essa app eh muita doida que eu vou fazer.
  
	E ainda printa em outra linha esse echo.
    `,
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand()
	//rootCmd.PersistentFlags().StringVarP(&CopyFiles, "files", "f", "", "Files to copy.") //Essa flag é global vou usar no arquivo local.go e s3.go!
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
