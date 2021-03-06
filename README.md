# URL Marshaler

:warning: THIS LIB HAS BEEN INTEGRATED INTO https://github.com/xescugc/marshaler USE IT INSTEAD :warning:

[![GoDoc](https://godoc.org/github.com/xescugc/url-marshaler?status.svg)](https://godoc.org/github.com/xescugc/url-marshaler)

It's a simple type that implements the [Marshaler](https://golang.org/pkg/encoding/json/#Marshaler) and [Unmarshaler](https://golang.org/pkg/encoding/json/#Unmarshaler) interfaces for a URL transforming it to an string.

## Installation

```
$> go get github.com/xescugc/url-marshaler
```

## Usage

```go
import (
	"encoding/json"
	"fmt"
	"net/url"

	marshaler "github.com/xescugc/url-marshaler"
)

type User struct {
  Name  string        `json:"name"`
  URL   marshaler.URL `json:"url"`
}

func main() {
  u, _ := url.Parse("http://example.com")
  usr := User{
    Name: "Pepito",
    URL:  marshaler.URL{
      URL: u,
    },
  }

  b, _ := json.Marshal(usr)

  fmt.Println(string(b))
  // { "name": "Pepito", "url": "http://example.com" }
}
```
