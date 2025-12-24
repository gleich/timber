<div align="center">
    <h1>timber</h1>
    <a href="https://pkg.go.dev/go.mattglei.ch/timber"><img alt="Godoc Reference" src="https://pkg.go.dev/badge/go.mattglei.ch/timber?utm_source=godoc"></a>
    <img alt="lint workflow result" src="https://github.com/gleich/timber/actions/workflows/lint.yml/badge.svg">
    <img alt="test workflow result" src="https://github.com/gleich/timber/actions/workflows/test.yml/badge.svg">
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/gleich/timber">
    <img alt="Golang report card" src ="https://goreportcard.com/badge/go.mattglei.ch/timber">
    <br>
    <i>Easy to use & pretty logger for golang</i>
    <br>
    <br>
</div>

_The original package is [gleich/lumber](https://github.com/gleich/lumber)._

- [Install](#-install)
- [Logging Functions](#-logging-functions)
  - [`timber.Done()`](#timberDone)
  - [`timber.Info()`](#timberinfo)
  - [`timber.Debug()`](#timberdebug)
  - [`timber.Warning()`](#timberwarning)
  - [`timber.Error()`](#timbererror)
  - [`timber.ErrorMsg()`](#timbererrormsg)
  - [`timber.Fatal()`](#timberfatal)
  - [`timber.FatalMsg()`](#timberfatalmsg)
- [Customization](#️-customization)
- [Examples](#-examples)

## Install

Simply run the following from your project root:

```bash
go get -u go.mattglei.ch/timber
```

## Logging Functions

### [`timber.Done()`](https://pkg.go.dev/go.mattglei.ch/timber#Done)

Output a "DONE" log.

Demo:

```go
package main

import (
    "time"

    "go.mattglei.ch/timber"
)

func main() {
    timber.Done("loaded up the program!")
    time.Sleep(2 * time.Second)
    timber.Done("waited 2 seconds")
}
```

Outputs:

![Done output](images/done.png)

> [!NOTE]
> Like all other logging functions there is a `log.Printf()` version of this function, which in this case would be `timber.Donef()`. Example:

```go
timber.Donef("running for %d minutes", 21)
```

### [`timber.Info()`](https://pkg.go.dev/go.mattglei.ch/timber#Info)

Output a info log.

Demo:

```go
package main

import (
    "time"

    "go.mattglei.ch/timber"
)

func main() {
    timber.Info("getting the current year")
    now := time.Now()
    timber.Info("current year is", now.Year())
}
```

Outputs:

![info output](images/info.png)

> [!NOTE]
> Like all other logging functions there is a `log.Printf()` version of this function, which in this case would be `timber.Infof()`. Example:

```go
timber.Infof("I am %d years old", 21)
```

### [`timber.Debug()`](https://pkg.go.dev/go.mattglei.ch/timber#Debug)

Output a debug log.

Demo:

```go
package main

import (
    "os"

    "go.mattglei.ch/timber"
)

func main() {
    homeDir, _ := os.UserHomeDir()
    timber.Debug("user's home dir is", homeDir)
}
```

Outputs:

![debug output](images/debug.png)

> [!NOTE]
> Like all other logging functions there is a `log.Printf()` version of this function, which in this case would be `timber.Debugf()`. Example:

```go
timber.Debugf("The value of foo is %t", true)
```

### [`timber.Warning()`](https://pkg.go.dev/go.mattglei.ch/timber#Warning)

Output a warning log.

Demo:

```go
package main

import (
    "time"

    "go.mattglei.ch/timber"
)

func main() {
    now := time.Now()
    if now.Year() != 2004 {
        timber.Warning("current year isn't 2004")
    }
}
```

Outputs:

![warning output](images/warning.png)

> [!NOTE]
> Like all other logging functions there is a `log.Printf()` version of this function, which in this case would be `timber.Warningf()`. Example:

```go
timber.Warningf("Responded with a HTTP status code of %d.", 404)
```

### [`timber.Error()`](https://pkg.go.dev/go.mattglei.ch/timber#Error)

Output an error log with a stack trace.

Demo:

```go
package main

import (
    "os"

    "go.mattglei.ch/timber"
)

func main() {
    fname := "invisible-file.txt"
    _, err := os.ReadFile(fName)
    if err != nil {
        timber.Error(err, "failed to read from", fname)
    }
}
```

Outputs:

![error output](images/error.png)

> [!NOTE]
> Like all other logging functions there is a `log.Printf()` version of this function, which in this case would be `timber.Errorf()`. Example:

```go
timber.Errorf(err, "Failed to open file from the %s directory", "/usr/bin")
```

### [`timber.ErrorMsg()`](https://pkg.go.dev/go.mattglei.ch/timber#ErrorMsg)

Output an error message with a stack trace.

Demo:

```go
package main

import "go.mattglei.ch/timber"

func main() {
    timber.ErrorMsg("error message")
}
```

Outputs:

![errorMsg output](images/errorMsg.png)

> [!NOTE]
> Like all other logging functions there is a `log.Printf()` version of this function, which in this case would be `timber.ErrorMsgf()`. Example:

```go
timber.ErrorMsgf("Failed to open file from the %s directory", "/usr/bin")
```

### [`timber.Fatal()`](https://pkg.go.dev/go.mattglei.ch/timber#Fatal)

Output a fatal log with a stack trace.

Demo:

```go
package main

import (
    "os"

    "go.mattglei.ch/timber"
)

func main() {
    fName := "invisible-file.txt"
    _, err := os.ReadFile(fName)
    if err != nil {
        timber.Fatal(err, "failed to read from", fName)
    }
}
```

Outputs:

![fatal output](images/fatal.png)

> [!NOTE]
> Like all other logging functions there is a `log.Printf()` version of this function, which in this case would be `timber.Fatalf()`. Example:

```go
timber.Fatalf(err, "Failed to delete %s from the system", "/usr/bin/timber")
```

### [`timber.FatalMsg()`](https://pkg.go.dev/go.mattglei.ch/timber#FatalMsg)

Output a fatal message with a stack trace.

Demo:

```go
package main

import "go.mattglei.ch/timber"

func main() {
    timber.FatalMsg("fatal message")
}
```

Outputs:

![fatalMsg output](images/fatalMsg.png)

> [!NOTE]
> Like all other logging functions there is a `log.Printf()` version of this function, which in this case would be `timber.FatalMsgf()`. Example:

```go
timber.FatalMsgf("Failed to delete %s from the system", "/usr/bin/timber")
```

## Customization

You can customize a number of different features of timber. Below is an example of some of this customization:

```go
package main

import (
    "time"

    "go.mattglei.ch/timber"
)

func main() {
    timber.Timezone(time.Local)
    timber.TimeFormat("Mon Jan 2 15:04:05 MST 2006")
    timber.FatalExitCode(0)

    timber.Done("calling from custom logger")
}
```

Check the [godoc documentation](https://pkg.go.dev/go.mattglei.ch/timber) to see all the customization functions.

# Examples

See some examples in the [\_examples/](_examples/) folder.
