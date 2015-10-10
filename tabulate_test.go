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
 * @file tabulate_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Oct 10 11:20:17 2015
 *
 **/

package gotabulate

import (
	"fmt"
	"testing"
)

var tabulator = NewTabulator()

func printStyleTable(format string) {
	fmt.Printf("%s\n", format)
	tabulator.SetHeader([]string{"id", "name", "age"})
	tabulator.SetFormat(format)
	fmt.Print(
		tabulator.Tabulate(
			[][]string{
				[]string{"long long long id 1", "crackcell", "27"},
				[]string{"2", "crackcell2", "27"},
				[]string{"3", "crackcell3", "27", "redundant cell"},
				[]string{"4"},
				[]string{"5", "crackcell5"}},
		))
	fmt.Println()
}

func printStyleTableFirstRow(format string) {
	fmt.Printf("%s\n", format)
	tabulator.SetFirstRowHeader(true)
	tabulator.SetFormat(format)
	fmt.Print(
		tabulator.Tabulate(
			[][]string{
				[]string{"long long long id 1", "crackcell", "27"},
				[]string{"2", "crackcell2", "27"},
				[]string{"3", "crackcell3", "27", "redundant cell"},
				[]string{"4"},
				[]string{"5", "crackcell5"}},
		))
	fmt.Println()
}

func TestTabluateAllStyle(t *testing.T) {
	printStyleTable("simple")
	printStyleTableFirstRow("simple")
	printStyleTable("orgtbl")
	printStyleTableFirstRow("orgtbl")
	printStyleTable("pipe")
	printStyleTableFirstRow("pipe")
}
