# gotabulate
Golang clone of Python [tabulate](https://pypi.python.org/pypi/tabulate).

## Installation

```
go install github.com/crackcell/gotabulate
```

## Library usage

```go
package main

import "github.com/crackcell/gotabulate"

func main() {
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
```

### Headers

### Table format

- plain
- simple
- grid
- fancy_grid
- pipe
- orgtbl
- rst
- mediawiki
- html
- latex
- latex_booktabs

### Column alignment

## Contributors

- Menglong TAN <tanmenglong@gmail.com>
