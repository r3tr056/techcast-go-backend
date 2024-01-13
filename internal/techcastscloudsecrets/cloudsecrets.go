
import (
	"context"
	"fmt"
	"hash/crc32"
	"io"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func createSecret(w io.Writer, parent, id string) error {
	// parent := "projects/my-project"
	// id := "my-secret"

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create secretmanager client : %v", err)
	}
	defer client.Close()

	// Build the request
	req := &secretmanagerpb.CreateSecretRequest{
		Parent: parent,
		SecretId: id,
		Secret: &secretmanagerpb.Secret {
			Replication: &secretmanagerpb.Replication {
				Replication: &secretmanagerpb.Replication_Automatic_ {
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}

	// Call the API
	result, err := client.CreateSecret(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to create secret: %v", err)
	}

	fmt.Fprintf(w, "Created secret : %s\n", result.Name)
	return nil
}

func addSecretVersion(w io.Writer, parent string) error {
	payload := []byte("my super secret data")
	crc32c := crc32.MakeTable(crc32.Castagnoli)
	checksum := int64(crc32.Checksum(payload, crc32c))

	// Create the client
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create secretmanager client : %v", err)
	}
	defer client.Close()

	// Build the request
	req := &secretmanagerpb.AddSecretVersionRequest{
		Parent: parent,
		Payload: &secretmanagerpb.SecretPayload {
			Data: payload,
			DataCrc32C: &checksum,
		},
	}

	result, err := client.AddSecretVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to add secret version : %v", err)
	}

	fmt.Fprintf(w, "Added secret version : %s\n", result.Name)
	return nil
}