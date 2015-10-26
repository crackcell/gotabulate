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

type TableInfo struct {
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
	info := &TableInfo{
		Headers:        this.Headers,
		FirstRowHeader: this.FirstRowHeader,
		Data:           data,
		ColumnSize:     0,
		CellWidth:      []int{},
	}
	if !this.FirstRowHeader {
		info.Update(info.Headers)
	}
	for _, row := range info.Data {
		info.Update(row)
	}
	if formatter, ok := formatters[strings.ToLower(this.Format)]; ok {
		return formatter.Format(info)
	} else {
		panic(fmt.Errorf("unsupported format: %s", this.Format))
	}
}

//===================================================================
// Private
//===================================================================

func (this *TableInfo) Update(row []string) {
	rowLen := len(row)
	if rowLen > this.ColumnSize {
		for i := 0; i < rowLen-this.ColumnSize; i++ {
			this.CellWidth = append(this.CellWidth, 0)
		}
		this.ColumnSize = rowLen
	}
	for i := 0; i < rowLen; i++ {
		l := len(row[i])
		if l > this.CellWidth[i] {
			this.CellWidth[i] = l
		}
	}
}
