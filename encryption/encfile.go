package encryption

import (
  "os"
  "errors"
)

func EncryptFile(filepath string, recipients ...string) (string, error) {
  fileh, err := os.Open(filepath)
  if err != nil { 
    return "", errors.New("FileAccessError")
  }
  stats, _ := fileh.Stat()
  var length int64 = stats.Size()
  data := make([]byte, length)
  fileh.Read(data)
  plaintext := string(data)
  
  ciphertext, err := Encrypt(plaintext, recipients...)
  if err != nil {
    return "", errors.New("EncryptionError") 
  }
  return ciphertext, nil
}