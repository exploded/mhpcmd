package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const OnOffSwitches = 8 // Number of on/off switches
const DewHeaters = 4    // Number of dew heaters

func main() {

	// example usage:
	//
	// Turn switch #1 on
	// mhpcmd.exe switch -num=1 -state=1
	//
	// Turn switch #1 off
	// mhpcmd.exe switch -num=1 -state=0
	//
	// Turn switch #4 on
	// mhpcmd.exe switch -num=4 -state=1
	//
	// Turn on all switches:
	// mhpcmd.exe switch -num=9 -state=1
	//
	// Turn dew heater #1 to max (i.e. 100%)
	// mhpcmd.exe dew -num=1 -level=100
	//
	// Turn dew heater #1 to 75%
	// mhpcmd.exe dew -num=1 -level=75
	//
	// Turn dew heater #1 off (i.e. 0%)
	// mhpcmd.exe dew -num=1 -level=0
	//
	// Default switch is #1, default action is On. This command will turn on switch #1:
	// mhpcmd switch
	//
	// for dew heaters, default is dew heater #1, and default level is 100. This command will set dew heater #1 to 100%:
	// mhpcmd dew
	//
	// To get help:
	// mhpcmd switch -h
	// mhpcmd dew -h

	switchCmd := flag.NewFlagSet("switch", flag.ExitOnError)
	switchno := switchCmd.Int("no", 1, "Switch 1 to 8. Switch 9 is all switches.")
	state := switchCmd.Int("state", 1, "1 = to turn on, 0 = to turn off")

	dewCmd := flag.NewFlagSet("dew", flag.ExitOnError)
	dewno := dewCmd.Int("num", 1, "Switch 1 to 4")
	devlevel := dewCmd.Int("level", 100, "0 to 100. 0 = off")

	if len(os.Args) < 2 {
		fmt.Println("expected switch number and on or off")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "switch":
		switchCmd.Parse(os.Args[2:])

		if *switchno < 1 || *switchno > OnOffSwitches+1 {
			fmt.Println("Invalid switch number")
			os.Exit(1)
		}

		if *state < 0 || *state > 1 {
			fmt.Println("Invalid on off value")
			os.Exit(1)
		}

		err := MhpSetSwitch(*switchno, (*state == 1))
		if err != nil {
			fmt.Println("Unable to turn switch")
			os.Exit(1)
		}

	case "dew":

		dewCmd.Parse(os.Args[2:])

		if *dewno < 1 || *dewno > DewHeaters {
			fmt.Println("Invalid dew number number")
			os.Exit(1)
		}

		if *devlevel < 0 || *devlevel > 100 {
			fmt.Println("Invalid dew level")
			os.Exit(1)
		}

		err := MhpSetDewHeater(*dewno, *devlevel)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Println("expected switch or dew command")
		os.Exit(1)
	}

}

// Function send the command to turn the 8 on/off switches on or off. id is from 1 to 9.
// 9 is all switches
func MhpSetSwitch(id int, state bool) (err error) {
	// Examples fround from Wireshark:
	//
	// Switch #1 on 100  (0x64)
	// Switch #1 off 99	(0x63)
	// ...
	// Switch #8 on 86	(0x56)
	// Switch #8 off 85	(0x55)
	var command int
	if id < 1 || id > OnOffSwitches+1 {
		err = errors.New("invalid switch number")
		return
	}
	command = 85 + (OnOffSwitches-id)*2
	// If the switch is to be turned on, add 1
	if state {
		command++
	}
	err = hidSend(command)
	if err != nil {
		return err
	}
	return
}

// Function sends the command to set the 4 variable switches (i.e. dew heater controllers).
// id is from 1 to 4. range / value is from 0 to 100 (0x00 to 0x64)
func MhpSetDewHeater(id int, value int) (err error) {
	// Examples
	// Dew #1 to 0      4b 00	75xx
	// Dew #1 to 50     4b 32	75xx
	// Dew #1 to 100    4b 64	75xx
	// ...
	// Dew #4 to 06     48 00	72xx
	// Dew #4 off 85    48 64	72xx
	var command int
	if id < 1 || id > DewHeaters {
		err = errors.New("invalid dew heater number")
		return
	}

	if value < 0 || value > 100 {
		err = errors.New("invalid switch level")
		return
	}

	// Value is in the 2 most significant digits.
	// Switch number is 2 least significant digits.
	command = (value * 0x100) + (0x48 + DewHeaters - id)
	err = hidSend(command)
	if err != nil {
		return err
	}
	return
}
