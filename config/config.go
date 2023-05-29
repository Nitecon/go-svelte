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

	c.SessionKey = c.getEnvKey("SESSION_KEY")
	c.SecureCookie = securecookie.New([]byte(c.SessionHash), []byte(c.SessionKey))
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
	/*c := Get()
	sgKey, err := GetGoogleSecret("sendgrid")
	if err != nil {
		log.Error().Err(err).Msg("failed to get sendgrid secret")
	}
	tsKey, err := GetGoogleSecret("typesense")
	if err != nil {
		log.Error().Err(err).Msg("failed to get typesense secret")
	}
	gmailPass, err := GetGoogleSecret("gmail")
	if err != nil {
		log.Error().Err(err).Msg("failed to get gmail secret")
	}
	c.SendGrid = sgKey
	c.TypeSenseKey = tsKey
	c.GmailPassword = gmailPass

	configLock.Lock()
	defer configLock.Unlock()
	config = c*/
}
