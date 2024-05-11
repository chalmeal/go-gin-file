package common

import (
	"context"
	cfg "go-gin-file/config"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func AwsS3Session(ctx context.Context) aws.Config {
	conf := map[string]string{
		"ENDPOINT":   cfg.GetInitKey("aws", "AWS_ENDPOINT"),
		"REGION":     cfg.GetInitKey("aws", "AWS_REGION"),
		"ACCESS_KEY": cfg.GetInitKey("aws", "AWS_ACCESS_KEY"),
		"SECRET_KEY": cfg.GetInitKey("aws", "AWS_SECRET_KEY"),
	}
	cred := credentials.NewStaticCredentialsProvider(conf["ACCESS_KEY"], conf["SECRET_KEY"], "")

	c, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(cred),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: conf["ENDPOINT"],
			}, nil
		})))
	if err != nil {
		log.Println(err)
		return aws.Config{}
	}

	return c
}

func FileUpload(p *s3.PutObjectInput) error {
	ctx := context.Background()

	client := s3.NewFromConfig(AwsS3Session(ctx), func(options *s3.Options) {
		options.UsePathStyle = true
	})

	if _, err := client.PutObject(ctx, p); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
