package gogpg

import (
  "github.com/kklash/gogpg/encryption"
  "github.com/kklash/gogpg/decryption"
  "github.com/kklash/gogpg/signatures"
)

func Encrypt(input string, recipients ...string) (string, error) {
  return encryption.Encrypt(input, recipients...)
}

func EncryptFile(filepath string, recipients ...string) (string, error) {
  return encryption.EncryptFile(filepath, recipients...)
}

func Decrypt(ciphertext string) (string, error) {
  return decryption.Decrypt(ciphertext)
}

func DecryptFile(filepath string) (string, error) {
  return decryption.DecryptFile(filepath)
}

func SignDetached(message string, key string) (string, error) {
  return signatures.SignDetached(message, key) 
}

func Sign(message string, key string) (string, error) {
  return signatures.Sign(message, key)
}

func VerifyDetached(message string, signature string) (bool, error) {
  return signatures.VerifyDetached(message, signature)
}

func Verify(signed_msg string) bool {
  return signatures.Verify(signed_msg)
}