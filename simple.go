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
 * @file simple.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Oct 10 15:57:57 2015
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

type SimpleFormatter struct{}

func NewSimpleFormatter() *SimpleFormatter {
	return &SimpleFormatter{}
}

func (this *SimpleFormatter) Format(table *Table) string {
	str := ""
	var header []string
	if table.FirstRowHeader && len(table.Data) > 0 {
		header = table.Data[0]
	} else {
		header = table.Headers
	}
	str += this.formatLine(header, table)
	str += this.formatHeaderSep(table)
	rowStart := 0
	if table.FirstRowHeader {
		rowStart = 1
	}
	for i := rowStart; i < len(table.Data); i++ {
		str += this.formatLine(table.Data[i], table)
	}
	return str
}

//===================================================================
// Private
//===================================================================

func (this *SimpleFormatter) formatLine(row []string, table *Table) string {
	str := ""
	l := len(row)
	str += fmt.Sprintf("%s%s  ", row[0], strings.Repeat(" ", table.CellWidth[0]-len(row[0])))
	for i := 1; i < l && i < table.ColumnSize; i++ {
		str += fmt.Sprintf("%s%s  ", strings.Repeat(" ", table.CellWidth[i]-len(row[i])), row[i])
	}
	str += "\n"
	return str
}

func (this *SimpleFormatter) formatHeaderSep(table *Table) string {
	str := ""
	headerSep := make([]string, table.ColumnSize)
	for i, _ := range headerSep {
		headerSep[i] = strings.Repeat("-", table.CellWidth[i])
	}
	str += this.formatLine(headerSep, table)
	return str
}
