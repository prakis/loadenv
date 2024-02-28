# LoadEnv
Load Environment variables from .env files

A simple and lightweight Go lang module to load environment variables from .env or other files.

Github Source: [Repo](https://github.com/prakis/loadenv/)

Please note that this is still in beta, I highly recommend [godotenv](https://github.com/joho/godotenv) OR [dotenvx](https://github.com/dotenvx/dotenvx).

# Installation

```shell
go get github.com/prakis/loadenv
```

# Usage

Add your environment variables to `.env` file in the root folder.

```shell
S3_BUCKET=YOURS3BUCKET
SECRET_KEY=YOURSECRETKEY
```

```go
package main

import (
    "log"
    "os"

    "github.com/prakis/loadenv"
)

func main() {
  err := loadenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  s3Bucket := os.Getenv("S3_BUCKET")
  secretKey := os.Getenv("SECRET_KEY")

}
```
