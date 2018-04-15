package decryption

import (
  "github.com/kklash/gogpg/execution"
  "errors"
)

const APP string = "gpg"

func Decrypt(ciphertext string) (string, error) {
  process := execution.Command {
    App: APP,
    Args: []string {"-q", "--allow-multiple-messages", "--decrypt"},
  }
  
  plaintext, err := process.Execute(ciphertext)
  if err != nil {
    return "", errors.New("DecryptionError")
  }
  
  return plaintext, nil
}
