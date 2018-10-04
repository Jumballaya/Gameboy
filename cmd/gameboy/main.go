package main

import (
	"fmt"
	"io/ioutil"
	"os"

	emulator "github.com/jumballaya/gameboy/emulator/core"
)

func Run(rom []byte) {
	emu := emulator.New()
	err := emu.Load(rom)
	if err != nil {
		fmt.Println(err)
		return
	}
	emu.Start()
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("You must provide a path to a valid gameboy ROM")
		return
	}

	rom, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Println(err)
	}

	Run(rom)
}
