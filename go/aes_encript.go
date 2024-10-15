package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// Initialize vector, which is the random bytes
var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// Keep this secret key with you.
//const secretKey string = "abc&1*~#^2^#s0^=)^^7%b34"
const secretKey string = "ozMjQ1MjM1MzJnaw"

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	texto := encodeBase64(cipherText)

	fmt.Println("encodeBase64: ", texto)

	return texto, nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	cipherText := decodeBase64(text)

	fmt.Println("decodeBase64: ", cipherText)

	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

func main() {
	fmt.Println("Hello rodrigo!")

	phrase := "Senha descriptografada"

	fmt.Println("Senha: ", phrase)

	fmt.Println("secretKey: ", secretKey)

	encText, err := Encrypt(phrase, secretKey)
	if err != nil {
		// itrlog.Fatalw("error encrypting your classified text: ", err)
		//
	}
	//fmt.Println("encrypted text: ", encText)

	// To decrypt the original phrase
	//decText, err := Decrypt("IyVL7wcMdPcTbsuy", secretKey)
	decText, err := Decrypt(encText, secretKey)
	if err != nil {
		// itrlog.Fatalw("error decrypting your encrypted text: ", err)
	}
	fmt.Println("decrypted text: ", decText)

	fmt.Println("decrypted text: ", "dzn/e0hiX/pog49BZxMhqkScBuhzdpn5zZprtmwizJo=")
	decText2, err2 := Decrypt("dzn/e0hiX/pog49BZxMhqkScBuhzdpn5zZprtmwizJo=", secretKey)
	if err2 != nil {
		// itrlog.Fatalw("error decrypting your encrypted text: ", err)
	}
	fmt.Println("decrypted text: ", decText2)
}
