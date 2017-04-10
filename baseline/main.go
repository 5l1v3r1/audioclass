package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/unixpickle/anydiff"
	"github.com/unixpickle/anynet"
	"github.com/unixpickle/anyvec"
	"github.com/unixpickle/anyvec/anyvec64"
	"github.com/unixpickle/audioset"
	"github.com/unixpickle/essentials"
)

func main() {
	var csvPath string
	var wavDir string

	flag.StringVar(&csvPath, "csv", "", "path to segment CSV file")
	flag.StringVar(&wavDir, "dir", "", "path to sample download directory")
	flag.Parse()

	if csvPath == "" || wavDir == "" {
		essentials.Die("Required flags: -csv and -dir. See -help.")
	}

	samples, err := audioset.ReadSet(wavDir, csvPath)
	if err != nil {
		essentials.Die(err)
	}

	instanceCounts := map[string]int{}
	for _, sample := range samples {
		for _, class := range sample.Classes {
			instanceCounts[class]++
		}
	}

	classes := samples.Classes()
	netOuts := make([]float64, len(classes))
	for i, class := range classes {
		prob := float64(instanceCounts[class]) / float64(len(samples))
		netOuts[i] = inverseSigmoid(prob)
	}

	outRes := anydiff.NewConst(anyvec64.MakeVectorData(netOuts))
	var totalCost float64
	for _, sample := range samples {
		ohVec := make([]float64, len(classes))
		for i, class := range classes {
			for _, x := range sample.Classes {
				if x == class {
					ohVec[i] = 1
				}
			}
		}
		desired := anydiff.NewConst(anyvec64.MakeVectorData(ohVec))
		costRes := (&anynet.SigmoidCE{}).Cost(desired, outRes, 1)
		totalCost += anyvec.Sum(costRes.Output()).(float64)
	}

	fmt.Println("Total cost:", totalCost)
	fmt.Println("Mean cost:", totalCost/float64(len(samples)))
}

func inverseSigmoid(x float64) float64 {
	return math.Log(-x / (x - 1))
}
