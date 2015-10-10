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

func TestTabluateSimple(t *testing.T) {
	fmt.Print(
		Tabulate(
			[][]string{
				[]string{"1", "crackcell"},
				[]string{"2", "crackcell2"},
				[]string{"3", "crackcell3", "redundant cell"},
				[]string{"4"},
				[]string{"5", "crackcell5"}},
			[]string{"id", "name"},
			"simple",
		))
}
