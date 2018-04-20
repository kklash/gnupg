package signatures

import (
  "github.com/kklash/gogpg/execution"
  "errors"
  "io/ioutil"
  "os"
  "regexp"
)

const APP string = "gpg"

func VerifyDetached(message, signature string) (bool, error) {
  msg_file, msg_err := ioutil.TempFile("/tmp", "gogpg")
  if msg_err != nil {
    return false, errors.New("FileAccessError: Could not create msg temp file for verification") 
  }
  sig_file, sig_err := ioutil.TempFile("/tmp", "gogpg")
  if sig_err != nil { 
    return false, errors.New("FileAccessError: Could not create sig temp file for verification") 
  }
  defer os.Remove(msg_file.Name())
  defer os.Remove(sig_file.Name())
  
  msg_file.WriteString(message)
  sig_file.WriteString(signature)
  

  process := execution.Command {
    App:  APP,
    Args: []string { "--verify", sig_file.Name(), msg_file.Name() },
  }
  return process.CheckSuccess(), nil
}

func Verify(signed_msg string) bool {
  process := execution.Command {
    App:  APP,
    Args: []string { "--verify" },
  }
  result := process.CheckSuccess(signed_msg)
  return result
}

func DecryptAndVerifyString(ciphertext, key string) (bool, string, error) {
  tmpfile, err := ioutil.TempFile("/tmp", "gogpg")
  if err != nil {
    return false, "", errors.New("FileAccessError: Could create temp file for gpg output logging")
  }
  defer tmpfile.Close()
  defer os.Remove(tmpfile.Name())
  process := execution.Command {
    App:  APP,
    Args: []string { "--status-file", tmpfile.Name(), "-o", "-", "-d" },
  }
  output, err := process.Execute(ciphertext)
  if err != nil { 
    return false, "", errors.New("DecryptionError: Could not decrypt ciphertext string. Is the key available?")
  }
  tmpstats, _ := tmpfile.Stat()
  log := make([]byte, tmpstats.Size())
  tmpfile.Read(log)
  goodsig, _ := regexp.Match("GOODSIG \\S* "+key, log)
  if goodsig {
    return true, output, nil
  }
  return false, "", nil
}