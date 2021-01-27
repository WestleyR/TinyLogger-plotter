// Created by WestleyR on 2021-01-27
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

package csvParse

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

// GetXYDataFromFile will return the X Y data from a csv file path. Returns
// an error if one occures.
func GetXYDataFromFile(fileName string) ([]float64, []float64, error) {
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
