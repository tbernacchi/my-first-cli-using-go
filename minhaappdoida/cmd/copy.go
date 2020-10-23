package cmd

import (
	"fmt"
	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "It'll copy files locally",
	Run: func(cmd *cobra.Command, args []string) {

		// Get values of flag arguments
		files, _ := cmd.Flags().GetString("files")
		dest, _ := cmd.Flags().GetString("dest")

		dir := (files + "/")
		CheckDir(dir)

		if (len(os.Args)) == 3 {
			// check if the dir exist.
			src, err := os.Stat(dir)
			if err != nil {
				fmt.Println(err)
			}

			// check if the dir is indeed a directory, if it's copy.
			if src.IsDir() {
				err = copy.Copy(dir, dest)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("not a directory")
				os.Exit(1)
			}
		} else {
			//Check if the user pass wildcards, e.g: /tmp/dir/*.txt
			threearg := (os.Args[3:])                //From 3 to the end. (The args here it's counting until three.)
			matchfiles := threearg[:len(threearg)-2] //Remove the last two.

			for _, match := range matchfiles {
				dir := strings.Split(match, "/")      //Split the input with *
				mypath := (strings.Join(dir[:], "/")) //Convert to path.
				thenamefiles := dir[len(dir)-1]       //The last value, the name of files.

				//Check if there is dir in the directory when passing e.g: /tmp/texto/*
				src, err := os.Stat(match)
				if err != nil {
					fmt.Println(err)
				}

				if !src.IsDir() {
					//Check if the match iterate it's not a directory.
					myinputread, err := ioutil.ReadFile(mypath)
					if err != nil {
						fmt.Println(err)
						return
					}

					destfiles := (dest + "/" + thenamefiles)

					err = ioutil.WriteFile(destfiles, myinputread, 0644)
					if err != nil {
						fmt.Println("Error creating", match)
						fmt.Println(err)
						return
					}
					fmt.Println("Copying", match)

				} else {
					//Copy directories.
					subdirmatch := strings.Split(match, "/")
					subdirname := subdirmatch[len(subdirmatch)-1]
					path := filepath.Join(dest, subdirname)

					err = os.MkdirAll(path, 0755)
					if err != nil {
						fmt.Println(err)
					}

					//Copy dir and files in subdirectories.
					err = copy.Copy(match, path)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("Copying", match)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringP("dest", "d", "", "destination of files")
	copyCmd.MarkFlagRequired("dest")
}
