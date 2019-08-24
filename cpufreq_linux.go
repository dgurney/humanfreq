package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func readCPUFreq(cpu int) (int, error) {
	path := "/sys/devices/system/cpu/cpu" + strconv.Itoa(cpu) + "/cpufreq/scaling_cur_freq"
	freq, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}
	final, err := strconv.Atoi(strings.TrimSuffix(string(freq), "\n"))
	if err != nil {
		return 0, err
	}
	return final, nil
}

func main() {
	v := flag.Bool("v", false, "Show version and exit")
	flag.Parse()
	if *v {
		fmt.Printf("humanfreq v%s by Daniel Gurney\n", version)
		return
	}
	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		for i := 0; i < runtime.NumCPU(); i++ {
			freq, err := readCPUFreq(i)
			if err != nil {
				fmt.Println("CPU frequency read error:", err)
				return
			}
			mhz := freq / 1000
			fmt.Printf("CPU%d: %d MHz\n", i, mhz)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
