package atenea

import (
    "crypto/aes"
    "crypto/cipher"
    cryptoRand "crypto/rand"
    "fmt"
    "io"
)

// Encrypt cifra data con key (AES-GCM) - ← Mayúscula (aunque no se usa fuera del paquete)
func encrypt(data []byte, key []byte) ([]byte, error) {  // ← minúscula (función interna)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(cryptoRand.Reader, nonce); err != nil{  // ← Usar cryptoRand
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Decrypt descifra data con key (AES-GCM)
func decrypt(data []byte, key []byte) ([]byte, error) {  // ← minúscula (función interna)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize() 
	if len(data) < nonceSize{
		return nil, fmt.Errorf("Data corrupted or incomplete")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil{
		return nil, fmt.Errorf("password incorrect or corrupted files")
	}

	return plaintext, nil
}