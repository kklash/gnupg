package execution

import (
  "os/exec"
  "io"
)

const APP string = "gpg"

type Command struct {
  App string
  Args []string
}
func (CMD *Command) AddArgs(new_args ...string) {
  for _, a := range new_args {
    CMD.Args = append(CMD.Args, a)
  }
}

func (CMD *Command) Execute(stdin ...string) (string, error) {
  var process *exec.Cmd = exec.Command(CMD.App, CMD.Args...)
  if stdin != nil { 
    pipe, _ := process.StdinPipe()
    for _, input_str := range stdin {
      io.WriteString(pipe, input_str)
    }
    pipe.Close()
  }
  output, err :=  process.Output()
  if err != nil { return "", err }
  return string(output), nil
}

func (CMD *Command) CheckSuccess(stdin ...string) bool {
  var process *exec.Cmd = exec.Command(CMD.App, CMD.Args...)
  if stdin != nil { 
    pipe, _ := process.StdinPipe()
    for _, input_str := range stdin {
      io.WriteString(pipe, input_str)
    }
    pipe.Close()
  }
  err :=  process.Run()
  if err == nil { 
    return true
  } else {
    return false
  }
}