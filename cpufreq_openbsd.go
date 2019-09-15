package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func readCPUFreq() (uint32, error) {
	call, err := syscall.Sysctl("hw.cpuspeed")
	if err != nil {
		return 0, err
	}
	freq := append([]byte(call), 0x00)
	fmt.Println()
	return binary.LittleEndian.Uint32(freq), nil
}

func main() {
	g := flag.Bool("g", false, "Use GHz instead of MHz.")
	v := flag.Bool("v", false, "Show version and exit.")
	flag.Parse()
	if *v {
		fmt.Printf("humanfreq v%s by Daniel Gurney\n", version)
		return
	}
	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		freq, err := readCPUFreq()
		if err != nil {
			fmt.Println("CPU frequency read error:", err)
			return
		}
		var conv float32
		switch {
		default:
			fmt.Printf("CPU: %d MHz\n", freq)
		case *g:
			conv = float32(freq) / 1000 // GHz
			fmt.Printf("CPU: %.2f GHz\n", conv)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
