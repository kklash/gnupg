package signatures

import (
  "github.com/kklash/gogpg/execution"
  "errors"
)

const APP string = "gpg"

func SignDetached(message string, key string) (string, error) {
  process := execution.Command {
    App:  APP,
    Args: []string { "-a", "-b", "-u", key, "--sign" },
  }

  signature, err := process.Execute(message)
  if err != nil { 
    return "", errors.New("SignaturesError") 
  }
  return signature, nil
}

func Sign(message string, key string) (string, error) {
  process := execution.Command {
    App:  APP,
    Args: []string { "-a", "-u", key, "--sign" },
  }
  
  signed_msg, err := process.Execute(message)
  if err != nil { 
    return "", errors.New("SignaturesError") 
  }
  return signed_msg, nil
}