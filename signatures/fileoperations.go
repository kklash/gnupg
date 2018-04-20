package signatures

import (
  "os"
  "io/ioutil"
  "regexp"
  "errors"
  "github.com/kklash/gogpg/execution"
)

func DecryptAndVerifyFile(cipher_file, output_file, key string) (bool, error) {
  tmpfile, err := ioutil.TempFile("/tmp", "gogpg")
  if err != nil {
    return false, errors.New("FileAccessError: Could not make temp file for gpg logging")
  }
  defer tmpfile.Close()
  defer os.Remove(tmpfile.Name())

  
  process := execution.Command {
    App:  APP,
    Args: []string { "--status-file", tmpfile.Name(), "-o", output_file, "-d", cipher_file },
  }
  success := process.CheckSuccess()
  tmpstats, _ := tmpfile.Stat()
  log := make([]byte, tmpstats.Size())
  tmpfile.Read(log)
  goodsig, _ := regexp.Match("GOODSIG \\S* " + key, log)
  if goodsig && success {
    return true, nil
  }
  return false, nil
}

func VerifyFileDetached(src_file, sig_file string) (bool, error) {
  _, src_err := os.Stat(src_file)
  if os.IsNotExist(src_err)  {
    return false, errors.New("FileAccessError: could not find " + src_file)
  }
  _, sig_err := os.Stat(sig_file)
  if os.IsNotExist(sig_err) {
    return false, errors.New("FileAccessError: could not find " + sig_file)
  }
  
  process := execution.Command {
    App:  APP,
    Args: []string { "--verify", sig_file, src_file },
  }
  return process.CheckSuccess(), nil
}

func SignAndEncryptFile(filepath, key string) (string, error) {
  cipher_file := filepath + ".gpg"
  process := execution.Command {
    App:  APP,
    Args: []string { "-a", "-u", key, "-r", key, "-o", cipher_file, "-s", "-e", filepath },
  }
  if ! process.CheckSuccess() {
    return "", errors.New("EncryptionError: Could not sign and encrypt " + filepath)
  }
  return cipher_file, nil
}