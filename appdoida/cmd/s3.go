package cmd

import (
	//"fmt"
	"github.com/spf13/cobra"
)

var bucketname string

var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Vou copiar arquivos locais do meu sistema para um S3 da AWS com esse subcomando da appdoida junto com flags",
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("s3 called")
	//},
}

func init() {
	rootCmd.AddCommand(s3Cmd)
	s3Cmd.Flags().StringVarP(&bucketname, "bucket", "b", "", "The name of the bucket.")
}
