package encryption

import (
  "errors"
  "github.com/kklash/gogpg/execution"
)

func EncryptFile(filepath, output_path string, recipients ...string) error {
  if len(recipients) < 1 {
    return errors.New("EncryptionError: no recipients given to EncryptFile")
  }
  
  process := execution.Command {
    App:  APP,
    Args: make([]string, 0, len(recipients)*2 + 10),
  }
  process.AddArgs("-a", "--yes", "--always-trust", "-o", output_path)
  
  for i:=0; i<len(recipients); i++ {
    process.AddArgs("--recipient", recipients[i])
  }
  process.AddArgs("--encrypt", filepath)
  
  if process.CheckSuccess() {
    return nil
  } else {
    return errors.New("EncryptionError: Could not encrypt " + filepath)
  }
}