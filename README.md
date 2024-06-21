### README for LostPapyrus

# LostPapyrus

LostPapyrus is a minimalist and fast web framework for Go (Golang), inspired by Fiber. It aims to be easy to use, fast, and feature-rich, providing a broad range of functionalities for building web applications.

## Features

- Simple and intuitive API
- Supports GET and POST routes
- Middleware support
- JSON response handling
- Route parameters

## Installation

To install LostPapyrus, use the following command:

```bash
go get -u github.com/simplyYan/LostPapyrus
```

## Quick Start

Here is a simple example to get you started with LostPapyrus:

```go
package main

import (
    "fmt"
    "github.com/simplyYan/LostPapyrus"
    "net/http"
)

func main() {
    app := lostpapyrus.New()

    app.Use(func(ctx *lostpapyrus.Context, next lostpapyrus.HandlerFunc) {
        fmt.Println("Middleware before")
        next(ctx)
        fmt.Println("Middleware after")
    })

    app.Get("/", func(ctx *lostpapyrus.Context) {
        ctx.Send("Welcome to LostPapyrus!")
    })

    app.Get("/hello/:name", func(ctx *lostpapyrus.Context) {
        name := ctx.Params["name"]
        ctx.Send(fmt.Sprintf("Hello, %s!", name))
    })

    app.Post("/echo", func(ctx *lostpapyrus.Context) {
        var data map[string]interface{}
        if err := ctx.BindJSON(&data); err != nil {
            ctx.Status(http.StatusBadRequest).Send("Invalid JSON")
            return
        }
        ctx.JSON(data)
    })

    app.Listen(":3000")
}
```

## Contributing

We welcome contributions from the community. To contribute to LostPapyrus, you can:

1. **Open an Issue**: If you find a bug or have a feature request, please open an issue on the GitHub repository.

2. **Submit a Pull Request**: If you want to contribute code, fork the repository, make your changes, and submit a pull request. Please ensure your code follows the project's coding standards and includes appropriate tests.

## License

LostPapyrus is released under the BSD-3-Clause license. See the LICENSE file for more details.

---

Thank you for using LostPapyrus! If you have any questions or need further assistance, feel free to open an issue on the GitHub repository.
