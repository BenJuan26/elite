# elite

Go API for Elite Dangerous

# Installation

```bash
go get github.com/BenJuan26/elite
```

# Example Usage

```go
system, err := elite.GetStarSystem()
if err != nil {
    panic("Couldn't get star system: " + err.Error())
}
fmt.Println("Current star system is " + system)
```
