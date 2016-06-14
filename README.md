# cyrlat

Library to translate the Cyrillic to the Latin alphabet for use in semantic URLs, file names, etc.

**Usage:**
```sh
go get github.com/njcm/cyrlat
```
```go
package example

import(
  "fmt"
  "github.com/njcm/cyrlat"
)

func main() {
  cyrillic := "Навзание файла"
  name := cyrlat.Latin(cyrillic, cyrlat.UnSpace)
  fmt.Println(name)
}
```
**Modificators:**
- `cyrlat.Lower` - all letters mapped to their lower case
- `cyrlat.Upper` - all letters mapped to their upper case
- `cyrlat.UnSpace` - all spaces replaced underscore symbol
