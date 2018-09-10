package main

import (
  "log"
  "os"
)

func main(){
  cli := pload.New()
  if err := cli.Run(); err != nil {
    log.Fatal(err)
	}
  os.Exit(0)
}
