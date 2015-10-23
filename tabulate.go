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
	Headers        []string
	FirstRowHeader bool
	Data           [][]string
	ColumnSize     int
	CellWidth      []int
}

type Tabulator struct {
	Headers        []string
	FirstRowHeader bool
	Format         string
}

const (
	defaultFormat = "simple"
)

var formatters = map[string]Formatter{
	"simple": NewSimpleFormatter(),
	"orgtbl": NewOrgtblFormatter(),
	"pipe":   NewPipeFormatter(),
	"plain":  NewPlainFormatter(),
	"psql":   NewPsqlFormatter(),
	"grid":   NewGridFormatter(),
}

func NewTabulator() *Tabulator {
	return &Tabulator{
		FirstRowHeader: true,
		Format:         defaultFormat,
	}
}

func (this *Tabulator) SetHeader(header []string) {
	this.Headers = header
	this.FirstRowHeader = false
}

func (this *Tabulator) SetFirstRowHeader(use bool) {
	this.FirstRowHeader = use
}

func (this *Tabulator) SetFormat(format string) {
	this.Format = format
}

func (this *Tabulator) Tabulate(data [][]string) string {
	if len(data) == 0 {
		return ""
	}
	headerLen := 0
	if this.FirstRowHeader {
		headerLen = len(data[0])
	} else {
		headerLen = len(this.Headers)
	}
	if headerLen == 0 {
		return ""
	}

	table := &Table{
		Headers:        this.Headers,
		FirstRowHeader: this.FirstRowHeader,
		Data:           data,
		ColumnSize:     headerLen,
		CellWidth:      make([]int, headerLen),
	}
	table.UpdateCellWidth(table.Headers)
	for _, row := range table.Data {
		table.UpdateCellWidth(row)
	}
	if formatter, ok := formatters[strings.ToLower(this.Format)]; ok {
		return formatter.Format(table)
	} else {
		panic(fmt.Errorf("unsupported format: %s", this.Format))
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
