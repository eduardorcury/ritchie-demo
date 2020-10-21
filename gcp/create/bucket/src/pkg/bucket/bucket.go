package bucket

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
)

type Inputs struct {
	BucketName    string
	ProjectID     string
	Credentials   string
}

func (in Inputs) Run() {
	ctx := context.Background()
	creds, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		fmt.Printf("Failed to obtain credentials. Error: %v\n", err)
		return
	}
	in.createBucket(ctx)
}

func (in Inputs) createBucket(ctx context.Context) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("Failed to create client. Error: %v\n", err)
		return
	}
	if in.BucketName == "" {
		fmt.Println("Bucket name cannot be blank!")
		return
	}
	bucket := client.Bucket(in.BucketName)
	if err := bucket.Create(ctx, in.ProjectID, nil); err != nil {
		fmt.Printf("Failed to create bucket. Error: %v\n", err)
	}
	fmt.Printf("Created bucket %v in project %v", in.BucketName, in.ProjectID)

}
