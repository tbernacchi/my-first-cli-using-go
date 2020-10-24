package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var mylsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list files",
	Run: func(cmd *cobra.Command, args []string) {

		// Get values of flag arguments
		files, _ := cmd.Flags().GetString("files")

		if (len(os.Args)) == 4 {
			last1 := files[len(files)-1:]
			if last1 != "/" {
				files = (files + "/")
			}
			CheckDir(files)
			PrintDir(files)
		} else {
			fourarg := (os.Args[4:])

			for _, match := range fourarg {
				//Check if there is dir in the directory when passing e.g: /tmp/texto/*
				src, err := os.Stat(match)
				if err != nil {
					fmt.Println(err)
				}

				if !src.IsDir() {
					//Check if the match iterate it's not a directory.
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(match)
				} else {
					PrintDir(match + "/")
				}
			}
		}
	},
}

func PrintDir(dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(dir + file.Name())
	}
	return false
}

func CheckDir(dirName string) bool {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(mylsCmd)
	mylsCmd.Flags().StringP("files", "f", "", "files to list/copy")
	mylsCmd.MarkFlagRequired("files")
}
