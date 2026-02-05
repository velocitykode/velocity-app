package config

import "github.com/velocitykode/velocity/pkg/config"

// GetCryptoKey returns the crypto key (read at call time, not init time)
func GetCryptoKey() string {
	return config.Get("CRYPTO_KEY", config.Get("APP_KEY", ""))
}

// GetCryptoCipher returns the crypto cipher
func GetCryptoCipher() string {
	return config.Get("CRYPTO_CIPHER", "AES-256-CBC")
}

// Legacy variables for backwards compatibility (set after godotenv loads)
var (
	CryptoKey    string
	CryptoCipher string
)

func InitCrypto() {
	CryptoKey = GetCryptoKey()
	CryptoCipher = GetCryptoCipher()
}
