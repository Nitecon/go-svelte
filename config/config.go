package config

import (
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/securecookie"
)

var (
	config     *AppConfig
	configLock = new(sync.RWMutex)
)

// AppConfig is the General Configuration for the application.
type AppConfig struct {
	SessionHash          string
	CloudProvider        string
	GoogleProjectID      string
	GoogleAppCredentials string
	SessionKey           string
	Debug                bool
	SecureCookie         *securecookie.SecureCookie
	OAuth                OauthCli
}

// OauthCli is a struct that holds google oauth related configurations
type OauthCli struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

// Get returns a config key whether stored in chamber of secrets or in clear text as envar.
func (c *AppConfig) getEnvKey(key string) string {
	envKey := os.Getenv(key)
	return envKey
}

// Get function returns the currently set active configuration to be used.
func Get() *AppConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

// InitConfig uses a mutex to quickly store all the initial configurations that is used within the application that provides faster access.
func InitConfig() {
	c := &AppConfig{}
	c.GoogleProjectID = strings.TrimSuffix(c.getEnvKey("GOOGLE_PROJECT_ID"), "\n")
	c.GoogleAppCredentials = strings.TrimSuffix(c.getEnvKey("GOOGLE_APPLICATION_CREDENTIALS"), "\n")
	c.SessionKey = c.getEnvKey("SESSION_KEY")
	c.SecureCookie = securecookie.New([]byte(c.SessionHash), []byte(c.SessionKey))
	c.CloudProvider = c.getEnvKey("CLOUD_PROVIDER")
	c.OAuth.RedirectURI = c.getEnvKey("OAUTH_REDIRECT_URI")
	c.OAuth.ClientID = c.getEnvKey("OAUTH_CLIENT_ID")
	if c.CloudProvider == "" {
		c.CloudProvider = "gcp"
	}
	debugVal := strings.ToLower(c.getEnvKey("DEBUG"))

	if debugVal == "true" || debugVal == "enabled" {
		c.Debug = true
	}
	log.Info().Msgf("Configuration loaded, and google project id is: %s", c.GoogleProjectID)

	configLock.Lock()
	defer configLock.Unlock()
	config = c
}

func InitSecrets() {
	c := Get()
	if c.CloudProvider == "gcp" {
		oSec, err := GetGoogleSecret("OAUTH_CLIENT_SECRET")
		if err != nil {
			log.Error().Err(err).Msg("failed to get OAUTH secret")
		}
		c.OAuth.ClientSecret = oSec
	} else if c.CloudProvider == "aws" {
		s, err := NewSecretManager()
		if err != nil {
			log.Error().Err(err).Msg("failed to create secret manager client")
			return
		}
		/* it's session based, so we can retrieve all the following keys with a single aws session */
		oSec, err := s.GetAWSSecret("OAUTH_CLIENT_SECRET")
		if err != nil {
			log.Error().Err(err).Msg("failed to get OAUTH secret")
			return
		}
		c.OAuth.ClientSecret = oSec
	} else {
		log.Error().Msg("No known cloud provider is set, retrieving from environment variables")
		c.OAuth.ClientSecret = c.getEnvKey("OAUTH_CLIENT_SECRET")
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = c
}
