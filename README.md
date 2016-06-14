# cyrlat
Library to translate the Cyrillic to the Latin alphabet for use in semantic URLs, file names, etc.
Usage:
go get github.com/njcm/cyrlat
import "github.com/njcm/cyrlat"
name := cyrlat.Latin(cyrillic string[, modificators ...int])
Modificators:
cyrlat.Lower - all letters mapped to their lower case
cyrlat.Upper -all letters mapped to their upper case
cyrlat.UnSpace - all spaces replaced underscore symbol
$ cat README.md                                                     
# cyrlat
### Library to translate the Cyrillic to the Latin alphabet for use in semantic URLs, file names, etc.
**Usage:**
```
go get github.com/njcm/cyrlat
import "github.com/njcm/cyrlat"
name := cyrlat.Latin(cyrillic string[, modificators ...int])
```
**Modificators:**
- `cyrlat.Lower` - all letters mapped to their lower case
- `cyrlat.Upper` - all letters mapped to their upper case
- `cyrlat.UnSpace` - all spaces replaced underscore symbol
