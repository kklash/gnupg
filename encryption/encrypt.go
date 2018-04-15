package encryption

import (
  "github.com/kklash/gogpg/execution"
  "errors"
)

const APP string = "gpg"

func Encrypt(input string, recipients ...string) (string, error) {
  if len(recipients) < 1 {
    return "", errors.New("EncryptionError")
  }
  
  CMD := execution.Command {
    App: APP,
    Args: make([]string, 0, len(recipients)*2 + 10),
  }
  CMD.AddArgs("--armor", "--always-trust", "--output", "-")
  
  for i:=0; i<len(recipients); i++ {
    CMD.AddArgs("--recipient", recipients[i])
  }
  CMD.AddArgs("--encrypt", "-")
  ciphertext, err := CMD.Execute(input)
  if err != nil { 
    return "", errors.New("EncryptionError") 
  }
  return ciphertext, nil
}




