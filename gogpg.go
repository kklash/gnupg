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

func DecryptFile(filepath string, output_path string) error {
  return decryption.DecryptFile(filepath, output_path)
}

func SignDetached(message string, key string) (string, error) {
  return signatures.SignDetached(message, key) 
}

func Sign(message string, key string) (string, error) {
  return signatures.Sign(message, key)
}

func SignFileDetached(filepath string, key string) (string, error) {
  return signatures.SignFileDetached(filepath, key)
}

func SignAndEncryptFile(filepath, key string) (string, error) {
  return signatures.SignAndEncryptFile(filepath, key)
}
  
func VerifyDetached(message string, signature string) (bool, error) {
  return signatures.VerifyDetached(message, signature)
}

func Verify(signed_msg string) bool {
  return signatures.Verify(signed_msg)
}

func VerifyFileDetached(src_file string, sig_file string) (bool, error) {
  return signatures.VerifyFileDetached(src_file, sig_file)
}

func DecryptAndVerifyFile(cipher_file, output_file, key string) (bool, error) {
  return signatures.DecryptAndVerifyFile(cipher_file, output_file, key)
}

func DecryptAndVerifyString(ciphertext, key string) (bool, string, error) {
  return signatures.DecryptAndVerifyString(ciphertext, key)
}

