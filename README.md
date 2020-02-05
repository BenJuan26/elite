# elite

[![GoDoc reference](https://godoc.org/github.com/BenJuan26/elite?status.svg)](https://godoc.org/github.com/BenJuan26/elite)

Go API for [Elite Dangerous](https://elitedangerous.com)

## Description

This API can be used to extract data from Elite Dangerous while the game is running. The data is obtained from the journal files written by the game, and therefore is only limited by what those files provide.

At the moment that includes things like:

* The status of many ship properties, such as night vision, landing gear, headlights, and [many more](https://godoc.org/github.com/BenJuan26/elite/flags).
* The current star system.
* Information about the ship, such as hull, shields, jump range, and modules.
* Players stats regarding things like combat, mining, exploration, and trading.

For a more complete picture of what can be obtained from the API, [see the documentation](https://godoc.org/github.com/BenJuan26/elite).

## Motivation

I wrote this API in order to pass the data to [a custom controller](https://github.com/BenJuan26/edca) that I'm making for the game. I kept it as a separate package so that it could be consumed by anybody that might find the data useful. I think it could also be useful for a dashboard-type project, and in fact, the controller will have a display on it that will provide some of the info supplied by this API.

## Installation

```bash
go get github.com/BenJuan26/elite
```

## Example Usage

```go
import (
    "fmt"

    "github.com/BenJuan26/elite"
)

func main() {
    system, _ := elite.GetStarSystem()
    fmt.Println("Current star system is " + system)

    status, _ := elite.GetStatus()
    if status.Flags.Docked {
        fmt.Println("Ship is docked")
    } else {
        fmt.Println("Ship is not docked")
    }
}
```
