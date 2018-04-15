package signatures

import (
  "github.com/kklash/gogpg/execution"
  "errors"
  "os"
)

const APP string = "gpg"

func writeTempFile(filename string, contents string) (string, error) {
  filepath := "/tmp/" + filename
  fh, err := os.Create(filepath)
  if err != nil {
    return "", errors.New("FileAccessError")
  }
  defer fh.Close()
  fh.WriteString(contents)
  return filepath, nil
}

func VerifyDetached(message string, signature string) (bool, error) {
  msg_file, err := writeTempFile("msg", message)
  sig_file, err := writeTempFile("sig", signature)
  if err != nil { return false, err }
  defer os.Remove(msg_file)
  defer os.Remove(sig_file)

  process := execution.Command {
    App:  APP,
    Args: []string { "--verify", sig_file, msg_file },
  }
  result := process.CheckSuccess()
  return result, nil
}

func Verify(signed_msg string) bool {
  process := execution.Command {
    App:  APP,
    Args: []string { "--verify" },
  }
  result := process.CheckSuccess(signed_msg)
  return result
}