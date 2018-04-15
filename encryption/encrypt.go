package encryption

import (
  "github.com/kklash/gogpg/execution"
  "errors"
)

func Encrypt(input string, recipients ...string) (string, error) {
  CMD := execution.Command {
    App: APP,
    Args: make([]string, 0, len(recipients)+10),
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




