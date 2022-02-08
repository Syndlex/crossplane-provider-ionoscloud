package ionosclients

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	ionoscloud "github.com/ionos-cloud/sdk-go-dbaas-postgres"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	kubeclient "sigs.k8s.io/controller-runtime/pkg/client"

	apisv1alpha1 "github.com/ionos-cloud/crossplane-provider-ionoscloud/apis/v1alpha1"
)

const (
	// UserAgent is the user agent addition that identifies the Crossplane IONOS Cloud Clients
	UserAgent = "crossplane-provider-ionoscloud"
)

const (
	errTrackPCUsage = "cannot track ProviderConfig usage"
	errGetPC        = "cannot get ProviderConfig"
	errGetCreds     = "cannot get credentials"
	errNewClient    = "cannot create new Service"
)

// IonosServices contains ionos clients
type IonosServices struct {
	DBaaSPostgresClient *ionoscloud.APIClient
}

// credentials specify how to authenticate with the IONOS API
type credentials struct {
	// Username to use
	User string `json:"user"`
	// Password to use. The password must be base64 encoded to prevent parsing anc escaping issues with special characters.
	Password string `json:"password"`
	// Token is used instead of username and password if both are empty
	Token string `json:"token"`
	// HostURL is the baseURL of the IONOS API. Optional.
	HostURL string `json:"host_url"`
}

// NewIonosClients creates a IonosService from the given data. The data must be a json struct with the fields `User` and
// `Password`. Both fields must be a string value. The password string must be base64 encoded.
func NewIonosClients(data []byte) (*IonosServices, error) {
	creds := credentials{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, fmt.Errorf("failed to decode credentials: %w", err)
	}
	decodedPW := []byte("")
	var err error
	if creds.Password != "" {
		decodedPW, err = base64.StdEncoding.DecodeString(creds.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to decode password: %w", err)
		}
	}
	config := ionoscloud.NewConfiguration(creds.User, string(decodedPW), creds.Token, creds.HostURL)
	config.UserAgent = fmt.Sprintf("%v_%v", UserAgent, config.UserAgent)
	client := ionoscloud.NewAPIClient(config)
	return &IonosServices{
		DBaaSPostgresClient: client,
	}, nil
}

// ConnectForCRD resolves the referenced ProviderConfig and extracts the connection secret from that ProviderConfig.
// After that an ionos client is setup with those credentials.
func ConnectForCRD(ctx context.Context, mg resource.Managed, client kubeclient.Client, t resource.Tracker) (*IonosServices, error) {
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackPCUsage)
	}

	pc := &apisv1alpha1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetPC)
	}

	cd := pc.Spec.Credentials
	data, err := resource.CommonCredentialExtractor(ctx, cd.Source, client, cd.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, errGetCreds)
	}

	svc, err := NewIonosClients(data)
	if err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}
	return svc, nil
}
