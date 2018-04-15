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
  msg_file, msg_err := writeTempFile("msg", message)
  sig_file, sig_err := writeTempFile("sig", signature)
  if (msg_err != nil || sig_err != nil) { 
    return false, errors.New("FileAccessError") 
  }
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

func VerifyFile(src_file string, sig_file string) (bool, error) {
  _, src_err := os.Stat(src_file)
  _, sig_err := os.Stat(sig_file)
  if (os.IsNotExist(src_err) || os.IsNotExist(sig_err)) {
    return false, errors.New("FileAccessError")
  }
  
  process := execution.Command {
    App:  APP,
    Args: []string { "--verify", sig_file, src_file },
  }
  return process.CheckSuccess(), nil
  // if err == nil {
  //   return true, nil
  // } else {
  //   return false, nil
  // }
}