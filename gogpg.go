package gogpg

import (
  "github.com/kklash/gogpg/encryption"
  "github.com/kklash/gogpg/decryption"
  "github.com/kklash/gogpg/signatures"
)

func Encrypt(input string, recipients ...string) (string, error) {
  return encryption.Encrypt(input, recipients...)
}

func EncryptFile(filepath string, output_path string, recipients ...string) error {
  return encryption.EncryptFile(filepath, output_path, recipients...)
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

func SignFile(filepath string, key string) (string, error) {
  return signatures.SignFile(filepath, key)
}

func VerifyDetached(message string, signature string) (bool, error) {
  return signatures.VerifyDetached(message, signature)
}

func Verify(signed_msg string) bool {
  return signatures.Verify(signed_msg)
}

func VerifyFile(src_file string, sig_file string) (bool, error) {
  return signatures.VerifyFile(src_file, sig_file)
}