package main

import (
	"errors"
	"os"
)

const (
	authIssuerEnvName   = "AUTH_ISSUER"
	authClientIDEnvName = "AUTH_CLIENTID"
)

type Configuration struct {
	Auth AuthConfiguration
	Port string
}

func GetConfigFromEnvironment() (Configuration, error) {
	authConfig, err := createAuthConfiguration()
	if err != nil {
		return Configuration{}, err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config := Configuration{Auth: authConfig, Port: port}

	return config, nil
}

type AuthConfiguration struct {
	Issuer   string
	ClientID string
}

func (config AuthConfiguration) validate() error {
	if config.Issuer == "" {
		return errors.New(authIssuerEnvName + " environment variable is not set")
	}

	if config.ClientID == "" {
		return errors.New(authClientIDEnvName + " environment variable is not set")
	}

	return nil
}

func createAuthConfiguration() (AuthConfiguration, error) {
	config := AuthConfiguration{}
	config.Issuer = os.Getenv(authIssuerEnvName)
	config.ClientID = os.Getenv(authClientIDEnvName)

	err := config.validate()
	if err != nil {
		return AuthConfiguration{}, err
	}

	return config, nil
}
