package docker

import (
	"kim/config"
	"kim/util"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/fatih/color"
)

func awsImageRepositoryCreator(config config.Configstructure, repoPath string) string {
	if config.LoadAwsCredentialsFromConfig {
		os.Setenv("AWS_SDK_LOAD_CONFIG", string(1))
	}

	svc := ecr.New(session.New())

	repositoryURI := ""

	input := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(repoPath),
	}

	result, err := svc.CreateRepository(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			util.Throw(aerr)
		} else {
			util.Throw(err)
		}
		return repositoryURI
	}

	repositoryURI = aws.StringValue(result.Repository.RepositoryUri)

	util.PrintTabbed("Created Repository with name", repositoryURI, color.FgCyan)

	return repositoryURI
}
