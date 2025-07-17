package dsm

import (
	"fmt"
	"net/url"
	"os"

	sdk "github.com/senhasegura/dsmcli/sdk/iso"
)

type DsmClient struct {
	client      *sdk.Client
	name        string
	system      string
	environment string
}

/**
 * Constructor for DsmClient
 */
func NewDsmClient(client *sdk.Client, name string, environment string, system string) DsmClient {
	if string(name) == "" {
		fmt.Println("Error: Application name must be defined")
		os.Exit(1)
	}

	if string(environment) == "" {
		fmt.Println("Error: Environment must be defined")
		os.Exit(1)
	}

	if string(system) == "" {
		fmt.Println("Error: System must be defined")
		os.Exit(1)
	}

	a := DsmClient{
		name:        name,
		environment: environment,
		system:      system,
		client:      client,
	}

	return a
}

/**
 * RegisterApplication a new authorization for this application on senhasegura using the iso endpoint
 * "POST /iso/dapp/Application"
 */
func (a *DsmClient) RegisterApplication() (ApplicationResponse, error) {
	a.client.V("Registering Application on DevSecOps\n")

	a.client.Authenticate()

	data := url.Values{
		"application": {a.name},
		"environment": {a.environment},
		"system":      {a.system},
	}

	var appResp ApplicationResponse
	err := a.client.Post("/iso/dapp/Application", data, &appResp)
	if err != nil {
		return ApplicationResponse{}, err
	}
	a.client.V("Application register success\n")

	return appResp, nil
}

/**
 * Authenticate on senhasegura with credentials provided from application api response
 */
func (a *DsmClient) DefineCredentialsByApplication(application ApplicationResponse) error {
	return a.client.DefineNewCredentials(application.ID, application.Signature)
}

/**
 * Makes requests for /iso/dapp/Application
 * to get Application
 */
func (a *DsmClient) GetApplication() (ApplicationResponse, error) {
	a.client.Authenticate()

	var appResp ApplicationResponse
	err := a.client.Get("/iso/dapp/Application", url.Values{}, &appResp)
	if err != nil {
		return ApplicationResponse{}, err
	}

	return appResp, nil
}

/**
 * Makes requests for /iso/dapp/Application
 * to get secrets of Application
 */
func (a *DsmClient) GetApplicationSecrets() (secrets, error) {
	a.client.V("Finding secrets from application\n")

	app, err := a.GetApplication()
	if err != nil {
		return nil, err
	}

	return app.Application.Secrets, nil
}

/**
 * Makes requests for /iso/sctm/secret
 * to get secrets of Current authorization
 */
func (a *DsmClient) ListSecrets() (ListSecretResponse, error) {
	a.client.V("Finding secrets from application\n")
	a.client.Authenticate()

	var resp ListSecretResponse
	err := a.client.Get("/iso/sctm/secret", url.Values{}, &resp)
	if err != nil {
		return ListSecretResponse{}, err
	}

	return resp, nil
}

func (a *DsmClient) GetClient() *sdk.Client {
	return a.client
}
