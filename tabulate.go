/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the BSD licence
 *
 **************************************************************/

/**
 *
 *
 * @file tabulate.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Oct  8 23:58:11 2015
 *
 **/

package gotabulate

import (
	"fmt"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type Table struct {
	Headers    []string
	Data       [][]string
	ColumnSize int
	CellWidth  []int
}

var formatters = map[string]Formatter{
	"simple": NewSimpleFormatter(),
}

func Tabulate(data [][]string, headers []string, tablefmt string) string {
	headerLen := len(headers)
	table := &Table{
		Headers:    headers,
		Data:       data,
		ColumnSize: headerLen,
		CellWidth:  make([]int, headerLen),
	}
	table.UpdateCellWidth(headers)
	for _, row := range table.Data {
		table.UpdateCellWidth(row)
	}
	if formatter, ok := formatters[strings.ToLower(tablefmt)]; ok {
		return formatter.Format(table)
	} else {
		panic(fmt.Errorf("unsupported format: %s", tablefmt))
	}
}

//===================================================================
// Private
//===================================================================

func (this *Table) UpdateCellWidth(row []string) {
	rowLen := len(row)
	var colSize int
	if rowLen < this.ColumnSize {
		colSize = rowLen
	} else {
		colSize = this.ColumnSize
	}
	for i := 0; i < colSize; i++ {
		cellSize := len(row[i])
		if cellSize > this.CellWidth[i] {
			this.CellWidth[i] = cellSize
		}
	}
}
