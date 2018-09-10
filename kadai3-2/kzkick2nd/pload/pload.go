package pload

import (
  "flags"
  "runtime"
)

type Pload struct {
	TargetDir string
	URL       string
	Procs     int
	timeout   int
}

func New() *Pload {
  // NOTE 位置的等、CPU数並列、タイムアウト10秒でとりあえずやってみようか
	return &Pload{
    TargetDir: "./",
		Procs:   runtime.NumCPU(),
		timeout: 10,
	}
}

func (pload *Pload) Run() error {
  if err := pload.Set(); err != nil {
  		return err
	}
  if err := pload.Load(); err != nil {
  		return err
	}
  if err := pload.Merge(); err != nil {
  		return err
	}
  return nil
}

func (pload *Pload) Set() error {
  return nil
}

func (pload *Pload) Load() error {
  return nil
}

func (pload *Pload) Merge() error {
  return nil
}
