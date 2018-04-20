package decryption

import (
  "github.com/kklash/gogpg/execution"
  "os"
  "errors"
)

func DecryptFile(filepath string, output_path string) error {
  if _, err := os.Stat(filepath); os.IsNotExist(err) {
    return errors.New("FileAccessError: "+filepath+" path not found")
  }
  
  process := execution.Command {
    App:  APP,
    Args: []string { "--output", output_path, "--decrypt", filepath },
  }
  
  success := process.CheckSuccess()
  if success {
    return nil
  } else {
    return errors.New("DecryptionError: could not decrypt " + filepath)
  }
}