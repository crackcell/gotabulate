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
	str += this.formatLine(table.Headers, table)
	headerSep := make([]string, table.ColumnSize)
	for i, _ := range headerSep {
		headerSep[i] = strings.Repeat("-", table.CellWidth[i])
	}
	str += this.formatLine(headerSep, table)
	for _, row := range table.Data {
		str += this.formatLine(row, table)
	}
	return str
}

//===================================================================
// Private
//===================================================================

func (this *SimpleFormatter) formatLine(row []string, table *Table) string {
	str := ""
	l := len(row)
	str += fmt.Sprintf("%s%s  ", strings.Repeat(" ", table.CellWidth[0]-len(row[0])), row[0])
	for i := 1; i < l && i < table.ColumnSize; i++ {
		str += fmt.Sprintf("%s%s  ", row[i], strings.Repeat(" ", table.CellWidth[i]-len(row[i])))
	}
	str += "\n"
	return str
}
