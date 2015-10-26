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

func (this *SimpleFormatter) Format(info *TableInfo) string {
	str := ""
	var header []string
	if info.FirstRowHeader && len(info.Data) > 0 {
		header = info.Data[0]
	} else {
		header = info.Headers
	}
	str += this.formatLine(header, info)
	str += this.formatHeaderSep(info)
	rowStart := 0
	if info.FirstRowHeader {
		rowStart = 1
	}
	for i := rowStart; i < len(info.Data); i++ {
		str += this.formatLine(info.Data[i], info)
	}
	return str
}

//===================================================================
// Private
//===================================================================

func (this *SimpleFormatter) formatLine(row []string, info *TableInfo) string {
	str := ""
	l := len(row)
	str += fmt.Sprintf("%s%s  ", row[0], strings.Repeat(" ", info.CellWidth[0]-len(row[0])))
	for i := 1; i < l && i < info.ColumnSize; i++ {
		str += fmt.Sprintf("%s%s  ", strings.Repeat(" ", info.CellWidth[i]-len(row[i])), row[i])
	}
	str += "\n"
	return str
}

func (this *SimpleFormatter) formatHeaderSep(info *TableInfo) string {
	str := ""
	headerSep := make([]string, info.ColumnSize)
	for i, _ := range headerSep {
		headerSep[i] = strings.Repeat("-", info.CellWidth[i])
	}
	str += this.formatLine(headerSep, info)
	return str
}
