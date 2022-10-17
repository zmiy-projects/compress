# compress


# install

```bash
go get -u github.com/zmiy-projects/compress
```

# examples

```go
package main

import (
	"fmt"
	"github.com/zmiy-projects/compress"
)

func main() {
	result := compress.Compress("aaaabbbcccccaaaa")
	fmt.Println(result) // a8b3c5
}
```