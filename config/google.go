package config

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"fmt"
	"google.golang.org/api/option"
)

func GetGoogleSecret(keyName string) (string, error) {
	ctx := context.Background()

	// Replace <PROJECT_ID> with your project ID.
	projectID := Get().GoogleProjectID

	// Create the Secret Manager client.
	client, err := secretmanager.NewClient(ctx, option.WithCredentialsFile(Get().GoogleAppCredentials))
	if err != nil {
		return "", fmt.Errorf("failed to create secret manager client: %v", err)
	}

	// Build the resource name of the secret.
	name := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, keyName)

	// Retrieve the secret value.
	resp, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	})
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}

	// Convert the secret value to a string and return it.
	return string(resp.Payload.Data), nil
}
