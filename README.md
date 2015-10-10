# gotabulate
Golang clone of Python [tabulate](https://pypi.python.org/pypi/tabulate).

## Installation

```
go get github.com/crackcell/gotabulate
```

## Library usage

```go
package main

import "github.com/crackcell/gotabulate"

func main() {
	tabulator := gotabulate.NewTabulator()
	tabulator.SetFirstRowHeader(true)
	tabulator.SetFormat("orgtbl")
	fmt.Print(
		tabulator.Tabulate(
			[][]string{
				[]string{"1", "crackcell"},
				[]string{"2", "crackcell2"},
				[]string{"3", "crackcell3", "redundant cell"},
				[]string{"4"},
				[]string{"5", "crackcell5"}},
		))
}
```
Output:

```
| long long long id 1 |  crackcell |
|---------------------+------------|
| 2                   | crackcell2 |
| 3                   | crackcell3 |
| 4                   |            |
| 5                   | crackcell5 |
```

### Headers

### Table format

~~not supported yet~~

- ~~plain~~
- simple
- ~~grid~~
- ~~fancy_grid~~
- pipe
- orgtbl
- ~~rst~~
- ~~mediawiki~~
- ~~html~~
- ~~latex~~
- ~~latex_booktabs~~

#### simple

```
id                         name
-------------------  ----------
long long long id 1   crackcell
2                    crackcell2
3                    crackcell3
4
5                    crackcell5
```

#### orgtbl

```
| id                  |       name |
|---------------------+------------|
| long long long id 1 |  crackcell |
| 2                   | crackcell2 |
| 3                   | crackcell3 |
| 4                   |            |
| 5                   | crackcell5 |
```

#### pipe

```
| id                  |       name | age |
|:--------------------|-----------:|----:|
| long long long id 1 |  crackcell |  27 |
| 2                   | crackcell2 |  27 |
| 3                   | crackcell3 |  27 |
| 4                   |            |     |
| 5                   | crackcell5 |     |
```

### Column alignment

## Contributors

- Menglong TAN <tanmenglong@gmail.com>
