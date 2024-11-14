package dsm

import (
	"encoding/json"
	"fmt"
)

type ListSecretResponse struct {
	ID        string   `json:"id"`
	Signature string   `json:"signature"`
	Error     string   `json:"error"`
	Message   string   `json:"message"`
	Secrets   []Secret `json:"secrets"`
	Response  struct {
		Status    int    `json:"status"`
		Message   string `json:"message"`
		Error     bool   `json:"error"`
		ErrorCode int    `json:"error_code"`
	} `json:"response"`
}

type secrets []Secret

type Secret struct {
	SecretID       string              `json:"secret_id"`
	SecretName     string              `json:"secret_name"`
	Identity       string              `json:"identity"`
	Version        string              `json:"version"`
	ExpirationDate string              `json:"expiration_date"`
	Engine         string              `json:"engine"`
	Data           []map[string]string `json:"data"`
}

func (r *ListSecretResponse) Unmarshal(msg []byte) error {
	err := json.Unmarshal(msg, r)
	if err != nil {
		return err
	}

	return nil
}

/**
 * Validate the response of senhasegura server
 */
func (r *ListSecretResponse) Validate() error {
	if r.Error != "" {
		return fmt.Errorf(r.Message)
	}

	if r.Response.Error {
		return fmt.Errorf(r.Response.Message)
	}

	return nil
}

func (r *ListSecretResponse) GetError() string {
	return r.Error
}

func (r *ListSecretResponse) GetMessage() string {
	return r.Message
}

func (r *ListSecretResponse) GetAccessToken() string {
	return r.Message
}

func (r *ListSecretResponse) GetResponse() interface{} {
	return r.Response
}

func (r *ListSecretResponse) GetEntity() interface{} {
	return r.Response
}
