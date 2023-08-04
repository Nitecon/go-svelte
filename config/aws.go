package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/pkg/errors"
)

type SecretManager struct {
	service *secretsmanager.SecretsManager
}

func NewSecretManager() (*SecretManager, error) {
	// AWS SDK will use AWS_REGION environment variable if available
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create session")
	}

	svc := secretsmanager.New(sess)

	return &SecretManager{
		service: svc,
	}, nil
}

func (s *SecretManager) GetAWSSecret(secretName string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := s.service.GetSecretValue(input)
	if err != nil {
		return "", errors.Wrap(err, "failed to get secret value")
	}

	if result.SecretString != nil {
		return *result.SecretString, nil
	}

	return "", errors.New("unable to decode secret value")
}
