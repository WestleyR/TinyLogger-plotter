// Created by WestleyR on 2021-01-26
// Source code: https://github.com/WestleyR/csv-plotter
// Last modified data: 2021-01-26
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
	"bufio"
	"strings"
	"strconv"

	//chart "github.com/wcharczuk/go-chart/v2"
	chart "github.com/wcharczuk/go-chart"
)

func renderGraph(fp io.Writer, xData []float64, yData []float64) error {
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

	graph.Render(chart.PNG, fp)
	//graph.Render(chart.SVG, fp)

	return nil
}

func getXYDataFromFile(fileName string) ([]float64, []float64, error) {
	var xData []float64
	var yData []float64

	file, err := os.Open(fileName)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()

	// Loop thought the file line-by-line
	lineCount := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		// Skip the first line
		// TODO: figure out if the first column is x or y
		if lineCount == 0 {
			lineCount++;
			continue
		}
		lineValue := scanner.Text()

		dataArr := strings.Split(lineValue, ",") 
		// TODO: make sure theres two slices

		// TODO: for now, y is the first column, and x is the second
		y, err := strconv.ParseFloat(dataArr[0], 64)
		if err != nil {
			return nil, nil, err
		}
		x, err := strconv.ParseFloat(dataArr[1], 64)
		if err != nil {
			return nil, nil, err
		}

		xData = append(xData, x)
		yData = append(yData, y)

		lineCount++;
    }

    if err := scanner.Err(); err != nil {
        return nil, nil, err
    }

	return xData, yData, nil 
}

func main() {
	f, _ := os.Create("output.png")
	defer f.Close()

	x, y, err := getXYDataFromFile("data.csv")

	fmt.Printf("ERROR: %v\n", err)

	fmt.Printf("LEN: X=%v Y=%v\n", len(x), len(y))

	renderGraph(f, x, y)
}
