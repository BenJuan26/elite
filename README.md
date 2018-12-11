# elite

Go API for Elite Dangerous

# Installation

```bash
go get github.com/BenJuan26/elite
```

# Example Usage

```go
import (
    "fmt"
 
    "github.com/BenJuan26/elite"
)


func main() {
    // Errors not handled here
    system, _ := elite.GetStarSystem()
    fmt.Println("Current star system is " + system)

    status, _ := elite.GetStatus()
    flags := status.ExpandFlags()
    if flags.Docked {
        fmt.Println("Ship is docked")
    } else {
        fmt.Println("Ship is not docked")
    }
}
```
