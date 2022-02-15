# Looker

> 🔍 find app path and print it.

<p align="center">
  <img src="./preview.gif" width="700px" />
</p>

**A alternative package to `exec.LookPath()`.**

## The problem

The following code, relatively common approach to running external commands has a vulnerability on Windows:

```go
import "os/exec"

func gitAdd() error {
    cmd := exec.Command("git", "add", "Formula")
    return cmd.Run()
}
```

Searching the current directory (surprising behavior) before searching folders listed in the PATH environment variable (expected behavior) seems to be intended in Go and unlikely to be changed [**go#38736**](https://github.com/golang/go/issues/38736)

## Install

```bash
go get -v github.com/abdfnx/looker@v0.1.0
```

## Example:

```go
import (
    "os/exec"

    "github.com/abdfnx/looker"
)

func gitAdd() error {
    gitPath, err := looker.LookPath("git")

    if err != nil {
        return err
    }

    cmd := exec.Command(gitPath, "add", "Formula")
    return cmd.Run()
}
```
