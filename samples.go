package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/unixpickle/audioset"
	"github.com/unixpickle/essentials"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var csvPath string
	var wavDir string
	var align int

	flag.StringVar(&csvPath, "csv", "", "path to segment CSV file")
	flag.StringVar(&wavDir, "dir", "", "path to sample download directory")
	flag.IntVar(&align, "align", 512, "PCM sample count alignment")
	flag.Parse()

	if csvPath == "" || wavDir == "" {
		essentials.Die("Required flags: -csv and -dir. See -help.")
	}

	samples, err := audioset.ReadSet(wavDir, csvPath)
	if err != nil {
		essentials.Die(err)
	}

	classes := samples.Classes()

	for {
		perm := rand.Perm(len(samples))
		for _, i := range perm {
			sample := samples[i]
			data, err := sample.Read()
			if err != nil {
				essentials.Die(err)
			}
			if len(data)%align != 0 {
				padding := make([]float64, align-(len(data)%align))
				data = append(data, padding...)
			}
			printPCM(data)
			printClasses(classes, sample)
		}
	}
}

func printPCM(data []float64) {
	var parts []string
	for _, x := range data {
		parts = append(parts, strconv.FormatFloat(x, 'f', -1, 32))
	}
	fmt.Println(strings.Join(parts, " "))
}

func printClasses(classes []string, sample *audioset.Sample) {
	var vec []string
	for _, class := range classes {
		var present bool
		for _, x := range sample.Classes {
			if x == class {
				present = true
				break
			}
		}
		if present {
			vec = append(vec, "1")
		} else {
			vec = append(vec, "0")
		}
	}
	fmt.Println(strings.Join(vec, " "))
}
