package signatures

import (
  "github.com/kklash/gogpg/execution"
  "errors"
)

func SignDetached(message, key string) (string, error) {
  process := execution.Command {
    App:  APP,
    Args: []string { "-a", "-b", "-u", key, "--sign" },
  }

  signature, err := process.Execute(message)
  if err != nil { 
    return "", errors.New("SignaturesError: Could not detached-sign string with key " + key) 
  }
  return signature, nil
}

func Sign(message, key string) (string, error) {
  process := execution.Command {
    App:  APP,
    Args: []string { "-a", "-u", key, "--sign" },
  }
  
  signed_msg, err := process.Execute(message)
  if err != nil { 
    return "", errors.New("SignaturesError: Could not sign string with key " + key) 
  }
  return signed_msg, nil
}

func SignFileDetached(filepath, key string) (string, error) {
  process := execution.Command {
    App:  APP,
    Args: []string { "-a", "-b", "-u", key, "-o", "-", "--sign", filepath },
  }
  signature, err := process.Execute()
  if err != nil { 
    return "", errors.New("SignaturesError: Could not detached-sign " + filepath + " with key " + key) 
  }
  return signature, nil
}

