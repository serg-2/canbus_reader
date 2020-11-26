package main

import (
	"./canlib"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(),
		"Usage send: %s [interface] send [arbid]_[data]\nUsage receive: %s [interface] receive\nUsage dump file: cat file | %s dump\n",
		os.Args[0], os.Args[0], os.Args[0])
	flag.PrintDefaults()
}


func main() {

	flag.Usage = usage
	flag.Parse()
	if !(len(flag.Args()) == 2 || len(flag.Args()) == 3) {
		usage()
		os.Exit(1)
	}

	canIf := flag.Args()[0]

	if canIf == "dump" {

		scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

		for scanner.Scan() {
			var frame canlib.CanFrame
			if strings.HasPrefix(scanner.Text(), "#") {
				continue
			}

			if strings.HasPrefix(scanner.Text(), "<") {
				frame = canlib.ParseFrameFromStringCanBoat(scanner.Text())
			} else {
				if !strings.HasPrefix(scanner.Text(), " ") {
					continue
				}
				frame = canlib.ParseFrameFromStringUnix(scanner.Text())
			}

			canlib.PrintFrameCanBoatStyle(frame)

			if len(scanner.Text()) < 2 {
				os.Exit(0)
			}

		}
		os.Exit(0)
	}

	// DEVICE PART
	var device canlib.Interface

	device, err := canlib.NewRawInterface(canIf)
	if err != nil {
		fmt.Printf("could not open interface %s: %v\n",
			canIf, err)
		os.Exit(1)
	}
	defer device.Close()

	if flag.Args()[1] == "send" {
		// SEND
		frameStr := flag.Args()[2]
		frame, err := canlib.ParseFrame(frameStr)
		if err != nil {
			fmt.Printf("could not parse frame: %v\n", err)
			os.Exit(1)
		}

		device.SendFrame(frame)
	} else if flag.Args()[1] == "receive" {
		// RECEIVE until Ctrl+C
		for {
			frame, err := device.RecvFrame()
			if err != nil {
				fmt.Printf("error receiving frame: %v", err)
				os.Exit(1)
			}

			// canlib.PrintFrameUnixStyle(device.IfName, frame)
			canlib.PrintFrameCanBoatStyle(frame)

		}
	}
}

