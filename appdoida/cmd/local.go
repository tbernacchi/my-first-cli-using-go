package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

var mymachineCmd = &cobra.Command{
	Use:   "local",
	Short: "Vou copiar arquivos de um dir para outro com esse subcomando da appdoida junto com as flags",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Get values of flag arguments
		files, _ := cmd.Flags().GetString("files")
		dest, _ := cmd.Flags().GetString("dest")

		input := string(files)
		output := string(dest)

		//Return string from filepath.Walk to use on ioutil.ReadFile to copy the files.
		var justfilename []string

		err := filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}
			if info.IsDir() {
				return nil
			}
			justfilename = append(justfilename, info.Name()) //Just filename.
			return nil
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		total := 0

		for _, file := range justfilename {
			inputfiles := (input + "/" + file)
			read, err := ioutil.ReadFile(inputfiles)
			if err != nil {
				fmt.Println(err)
				return
			}
			checkDir(output)
			destfiles := (output + "/" + file)
			err = ioutil.WriteFile(destfiles, read, 0644) //Has to be the fullpath.
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Copying", inputfiles)
			total++
		}
		fmt.Println("Total files copied to", output, "=>", total)
	},
}

func checkDir(dirName string) bool {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(mymachineCmd)
	mymachineCmd.Flags().StringP("files", "f", "", "Files to copy.")
	mymachineCmd.Flags().StringP("dest", "d", "", "Destination of files.")
	mymachineCmd.MarkFlagRequired("files")
	mymachineCmd.MarkFlagRequired("dest")
}
