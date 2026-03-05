package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	APITokenPrefix  = "mo-"
	APITokenLength  = 32
	APITokenCharset = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Maker is an interface for managing tokens
type Maker interface {
	// GetValue returns raw token
	GetValue() string

	// GetHash returns token hash
	GetHash() string

	// GetPrefix extracts display prefix from full token (e.g., "mo-a3K9m")
	GetPrefix() string

	// VerifyToken checks if the token is valid or not
	VerifyToken() error
}

type APIToken struct {
	token string
}

func NewAPIToken(token string) *APIToken {
	if token == "" {
		token = generateAPIToken()
	}
	return &APIToken{
		token: token,
	}
}

// generateAPIToken creates a new API token with mo- prefix
func generateAPIToken() string {
	b := make([]byte, APITokenLength)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	result := make([]byte, APITokenLength)
	for i := range b {
		result[i] = APITokenCharset[b[i]%byte(len(APITokenCharset))]
	}

	return APITokenPrefix + string(result)
}

// GetValue returns raw token value
func (t APIToken) GetValue() string {
	return t.token
}

// GetHash returns SHA256 hash of the token for storage
func (t APIToken) GetHash() string {
	hash := sha256.Sum256([]byte(t.token))
	return hex.EncodeToString(hash[:])
}

// GetPrefix extracts display prefix from full token (e.g., "mo-a3K9m")
func (t APIToken) GetPrefix() string {
	if len(t.token) < 8 {
		return t.token
	}
	return t.token[:8] // "mo-" + 5 chars
}

func (t APIToken) VerifyToken() error {
	// Check token format
	if !strings.HasPrefix(t.token, APITokenPrefix) {
		return fmt.Errorf("invalid API key format")
	}

	return nil
}
