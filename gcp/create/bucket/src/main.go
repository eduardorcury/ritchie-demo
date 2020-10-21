package main

import (
	"gcp/create/bucket/pkg/bucket"
	"os"
	"strconv"
)

func main() {
	name        := os.Getenv("BUCKET")
	project     := os.Getenv("PROJECT")
	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	formula.Formula{
		BucketName:  name,
		ProjectID:   project,
		Credentials: credentials
	}.Run(os.Stdout)
}
