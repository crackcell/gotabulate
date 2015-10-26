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
 * @file pipe.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Oct 10 16:35:28 2015
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

type PipeFormatter struct{}

func NewPipeFormatter() *PipeFormatter {
	return &PipeFormatter{}
}

func (this *PipeFormatter) Format(info *TableInfo) string {
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

func (this *PipeFormatter) formatLine(row []string, info *TableInfo) string {
	str := ""
	l := len(row)
	str += fmt.Sprintf("| %s%s ", row[0], strings.Repeat(" ", info.CellWidth[0]-len(row[0])))
	for i := 1; i < info.ColumnSize; i++ {
		v := ""
		if i < l {
			v = row[i]
		}
		str += fmt.Sprintf("| %s%s ", strings.Repeat(" ", info.CellWidth[i]-len(v)), v)
	}
	str += "|\n"
	return str
}

func (this *PipeFormatter) formatHeaderSep(info *TableInfo) string {
	str := fmt.Sprintf("|:%s", strings.Repeat("-", info.CellWidth[0]+1))
	for i := 1; i < info.ColumnSize; i++ {
		str += fmt.Sprintf("|%s:", strings.Repeat("-", info.CellWidth[i]+1))
	}
	str += "|\n"
	return str
}
