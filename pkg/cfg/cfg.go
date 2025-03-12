package cfg

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func InitiateClient(r string) (aws.Config, error) {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(r),
	)

	if err != nil {
		log.Fatal("Failed to initialize AWS client.")
	}

	return cfg, err
}

func InitiateClientWithProfile(r, p string) (aws.Config, error) {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(r),
		config.WithSharedConfigProfile(p),
	)

	if err != nil {
		log.Fatal("Failed to initialize AWS client.")
	}

	return cfg, err
}

