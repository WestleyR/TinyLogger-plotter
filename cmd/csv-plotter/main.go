// Created by WestleyR on 2021-01-26
// Source code: https://github.com/WestleyR/csv-plotter
// Last modified data: 2021-01-27
//
// This file is licensed under the terms of
//
// The Clear BSD License
//
// Copyright (c) 2021 WestleyR
// All rights reserved.
//

package main

import (
	"os"
	"io"
	"fmt"

	"github.com/WestleyR/csv-plotter/pkg/csvParse"
	flag "github.com/spf13/pflag"
	chart "github.com/wcharczuk/go-chart/v2"
)

func renderGraph(fp io.Writer, xData []float64, yData []float64, outputType string) error {
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "Minutes",
		},
		YAxis: chart.YAxis{
			Name: "Volts",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xData,
				YValues: yData,
			},
		},
	}

	if outputType == "png" {
		graph.Render(chart.PNG, fp)
	} else if outputType == "svg" {
		graph.Render(chart.SVG, fp)
	} else {
		return fmt.Errorf("invalid output type: %s", outputType)
	}

	return nil
}

func main() {

	helpFlag := flag.BoolP("help", "h", false, "Print this help output.")
	versionFlag := flag.BoolP("version", "V", false, "print srm version.")

	inputFileNameFlag := flag.StringP("input", "i", "", "input file name")
	outputFileNameFlag := flag.StringP("output", "o", "", "output file name")
	formatFlag := flag.StringP("format", "f", "png", "output image format")

	flag.Parse()

	// Help flag
	if *helpFlag {
		fmt.Printf("Copyright (c) 2021 WestleyR. All rights reserved.\n")
		fmt.Printf("This software is licensed under the terms of The Clear BSD License.\n")
		fmt.Printf("Source code: https://github.com/WestleyR/csv-plotter\n")
		fmt.Printf("\n")
		flag.Usage()
		os.Exit(0)
	}

	if *versionFlag {
		// TODO
		os.Exit(0)
	}

	if *inputFileNameFlag == "" {
		fmt.Fprintf(os.Stderr, "Need an input file. Use '-i=input.csv'\n")
		os.Exit(1)
	}

	if *outputFileNameFlag == "" {
		fmt.Fprintf(os.Stderr, "Need an output file. Use '-o=output.png'\n")
		os.Exit(1)
	}

	f, _ := os.Create(*outputFileNameFlag)
	defer f.Close()

	x, y, err := csvParse.GetXYDataFromFile(*inputFileNameFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting values from %s: %s\n", *inputFileNameFlag, err)
		os.Exit(1)
	}

	fmt.Printf("plotted %d data points to: %s\n", len(x), *outputFileNameFlag)

	renderGraph(f, x, y, *formatFlag)
}

