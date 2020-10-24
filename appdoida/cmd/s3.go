package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

// s3Cmd represents the s3 command
var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "upload files",
	Run: func(cmd *cobra.Command, args []string) {
		//Check AWS keys/env.
		checkAwsEnv()

		// Get values of flag arguments
		files, _ := cmd.Flags().GetString("files")
		bucket, _ := cmd.Flags().GetString("bucket")

		dir := (files + "/")
		CheckDir(dir)

		//Dont copy if it's a dir.
		mydir := (os.Args[3])
		length := len(mydir)
		last1 := mydir[length-1 : length]

		if len(os.Args) == 6 {
			checkdir := os.Args[3]
			src, err := os.Stat(checkdir)
			if err != nil {
				fmt.Println(err)
			}

			if src.IsDir() {
				checkdirerr := (checkdir + ":")
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("caused by:", "read", checkdirerr, "is a directory")
				os.Exit(1)
			}
		}

		//If it not a dir, copy files according the input of the user, e.g: /tmp/test/*.files
		if last1 == "/" {
			mydirerr := (mydir + ":")
			fmt.Println("caused by:", "read", mydirerr, "is a directory")
			os.Exit(1)
		} else {
			threeargs3 := os.Args[3:]
			filename := threeargs3[:len(threeargs3)-2] //Remove the last two.

			for _, f := range filename {

				src, err := os.Stat(f)
				if err != nil {
					fmt.Println(err)
				}

				// check if the file iterative it's a dir.
				if src.IsDir() {
					path, dir := filepath.Split(f)
					mypath := path + dir

					arquivos, err := ioutil.ReadDir(mypath) //pq que com .txt funcionou???
					if err != nil {
						fmt.Println(err)
					}

					for _, arq := range arquivos {
						subdirfiles := (dir + "/" + arq.Name())
						pathprint := (path + subdirfiles)

						subdir, err := os.Open(pathprint) //file *File (do os.stat)
						if err != nil {
							fmt.Println(err)
						}

						// Initialize a session , the SDK will load credentials from  ~/.aws/credentials.
						sess, err := session.NewSession(&aws.Config{
							//Region: aws.String("sa-east-1")},
							Region: aws.String(os.Getenv("AWS_REGION"))},
						)

						// Setup the S3 Upload Manager - http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
						uploader := s3manager.NewUploader(sess)

						// Upload the file's body to S3 bucket as an object with the key being dir1/file.txt
						_, err = uploader.Upload(&s3manager.UploadInput{
							Bucket: aws.String(bucket),
							Key:    aws.String(subdirfiles), //In this case I want dir/file.txt
							Body:   subdir,                  //Must be absolute path.
							ACL:    aws.String("private"),
						})

						if err != nil {
							// Print the error and exit.
							exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
						}
						fmt.Println("Uploading", pathprint)
					}

				} else {

					file, err := os.Open(f)
					if err != nil {
						fmt.Println(err)
					}

					// Initialize a session in the specific region that was set. It'll load credentials from  ~/.aws/credentials.
					sess, err := session.NewSession(&aws.Config{
						//Region: aws.String("sa-east-1")},
						Region: aws.String(os.Getenv("AWS_REGION"))},
					)

					// Setup the S3 Upload Manager - http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
					uploader := s3manager.NewUploader(sess)

					// Upload the file's body to S3 bucket as an object with the key being just the filename.
					_, err = uploader.Upload(&s3manager.UploadInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(filepath.Base(f)), //Relative path, just the name of file.
						Body:   file,                         //Must be absolute path.
						ACL:    aws.String("private"),
					})

					if err != nil {
						exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
					}
					fmt.Println("Uploading", f)
				}
			}
		}
	},
}

func checkAwsEnv() {
	keys := []string{"AWS_SECRET_ACCESS_KEY", "AWS_ACCESS_KEY_ID", "AWS_REGION"}
	for _, item := range keys {
		_, present := os.LookupEnv(item)
		if !present {
			fmt.Println("AWS keys/env not found", item)
			os.Exit(1)
		}
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func init() {
	rootCmd.AddCommand(s3Cmd)
	s3Cmd.Flags().StringP("bucket", "b", "", "bucketname")
	s3Cmd.MarkFlagRequired("bucket")
}
