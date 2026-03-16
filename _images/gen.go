package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"slices"
	"time"

	"go.mattglei.ch/timber"
)

func main() {
	timber.FatalExitCode(0)

	logs := map[string]func(){
		"debug": func() {
			home, _ := os.UserHomeDir()
			timber.Debug("loaded home dir", timber.A("path", home))
		},
		"info": func() {
			timber.Info("server listening", timber.A("port", 8080))
		},
		"done": func() {
			sum := 2 + 2
			timber.Done("computed the sum of 2 and 2", timber.A("sum", sum))
		},
		"warning": func() {
			year := time.Now().Year()
			if year != 2004 {
				timber.Warning("it is not 2004")
			}
		},
		"error": func() {
			filename := "foo.txt"
			_, err := os.ReadFile(filename)
			if err != nil {
				timber.Error(err, "failed to read file", timber.A("filename", filename))
			}
		},
		"errorMsg": func() {
			age := 21
			if age != 22 {
				timber.ErrorMsg("user is not 22")
			}
		},
		"fatal": func() {
			_, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				timber.Fatal(err, "failed to connect to server")
			}
		},
		"fatalMsg": func() {
			if os.Getenv("API_KEY") == "" {
				timber.FatalMsg("API_KEY environment variable is not set")
			}
		},
	}

	if len(os.Args) <= 1 {
		timber.Timezone(time.Local)
		timber.TimeFormat("03:04:05")

		for name := range logs {
			generateImage(name, false)
			generateImage(name, true)
		}
	} else {
		if slices.Contains(os.Args, "--structured") {
			timber.Structured(true)
		}
		logs[os.Args[1]]()
	}
}

func generateImage(name string, structured bool) {
	var (
		filename string
		cmd      = fmt.Sprintf("go run gen.go %s", name)
	)
	if structured {
		filename = fmt.Sprintf("%s-structured.png", name)
		cmd += " --structured"
	} else {
		filename = fmt.Sprintf("%s.png", name)
	}
	err := exec.Command("freeze",
		"--output", filename,
		"--font.family", "psudoFont Liga Mono",
		"--execute", cmd,
		"--border.radius", "8",
	).Run()
	if err != nil {
		timber.Fatal(err, "failed to generate image", timber.A("name", name))
	}
	timber.Done("generated " + filename)

}
