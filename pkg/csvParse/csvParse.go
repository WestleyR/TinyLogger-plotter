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
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetXYDataFromFile will return the X Y data from a csv file path. Returns
// an error if one occures. X and Y columns will be determined automatically
// based on the csv header.
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

	// The default column position
	xColumn := 1
	yColumn := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if lineCount == 0 {
			// This is first line in the file, so we should check if the
			// first column is X or Y.

			dataArr := strings.Split(scanner.Text(), ",")
			if len(dataArr) != 2 {
				return nil, nil, fmt.Errorf("file: %s needs to start with 'x,y'", fileName)
			}

			if (dataArr[0] != "x" && dataArr[0] != "y") || (dataArr[1] != "x" && dataArr[1] != "y") {
				return nil, nil, fmt.Errorf("file: %s needs to start with 'x,y' only", fileName)
			}

			if dataArr[0] == dataArr[1] {
				return nil, nil, fmt.Errorf("invalid csv header; got: %s; expecting: x,y", scanner.Text())
			}

			if dataArr[0] == "x" {
				xColumn = 0
				yColumn = 1
			}
			if dataArr[0] == "y" {
				yColumn = 0
				xColumn = 1
			}

			lineCount++
			continue
		}
		lineValue := scanner.Text()

		dataArr := strings.Split(lineValue, ",")
		// Make sure we have the two data values
		if len(dataArr) != 2 {
			return nil, nil, fmt.Errorf("missing value in: %s at line: %d", fileName, lineCount)
		}

		// Convert the string to float
		y, err := strconv.ParseFloat(dataArr[yColumn], 64)
		if err != nil {
			return nil, nil, err
		}
		x, err := strconv.ParseFloat(dataArr[xColumn], 64)
		if err != nil {
			return nil, nil, err
		}

		xData = append(xData, x)
		yData = append(yData, y)

		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return xData, yData, nil
}
