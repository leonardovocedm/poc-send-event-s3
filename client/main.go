package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		fmt.Println("Error load default config AWS: ", err)
	}

	client := s3.NewFromConfig(sdkConfig)

	file, err := os.Open("../obs.txt")
	if err != nil {
		fmt.Println("Error load file: ", err)
	}

	fmt.Println("Putting file in S3")
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("test-securitization"),
		Key:    aws.String("Test" + "REMESSA_DMPAG_" + "TESTEATARDE" + ".txt"),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Failed to put object in S3 due: %+v\n", err)
	}

	fmt.Println("sign file in AWS")
	presignClient := s3.NewPresignClient(client)
	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("test-securitization"),
		Key:    aws.String("Test" + "REMESSA_DMPAG_" + "TESTEATARDE" + ".txt"),
	},
		func(opts *s3.PresignOptions) {
			opts.Expires = time.Duration(84600 * int64(time.Second))
		})
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
	}

	fmt.Println("AWS KEY: ", aws.String("securitization/remessa/"+"REMESSA_DMPAG_"+"TESTE"+".txt"))
	fmt.Println("URL: ", req.URL)

	if err != nil {
		fmt.Println("Error insert object in bucket: ", err)
	}
}
