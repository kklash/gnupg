package decryption

import (
  "os"
  "errors"
)

func DecryptFile(filepath string) (string, error) {
  fileh, err := os.Open(filepath)
  if err != nil {
    return "", errors.New("FileAccessError") 
  }
  stats, _ := fileh.Stat()
  var length int64 = stats.Size()
  data := make([]byte, length)
  fileh.Read(data)
  ciphertext := string(data)
  
  plaintext, err := Decrypt(ciphertext)
  if err != nil {
    return "", errors.New("DecryptionError")
  }
  return plaintext, nil
}