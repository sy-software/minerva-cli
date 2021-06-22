package service

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/radovskyb/watcher"
)

type TestRunnerService struct {
}

func (runner *TestRunnerService) RunTests(path string, watch bool, args []string) error {
	if watch {
		return watchMode(path, args)
	} else {
		return executeTests(path, args)
	}
}

func (runner *TestRunnerService) DiscoverTests(path string) error {
	return nil
}

// executeTests Runs go test command and prints the output using terminal color
func executeTests(path string, args []string) error {
	fmt.Println("Running tests...")
	goPath, err := exec.LookPath("go")

	if err != nil {
		return errors.New("go command not found")
	}

	var buffer, errBuf bytes.Buffer

	bashCommand := &exec.Cmd{
		Path:   goPath,
		Args:   append([]string{goPath, "test", "-v", path}, args...),
		Stdout: &buffer,
		Stderr: &errBuf,
	}

	err = bashCommand.Run()

	passRegex := regexp.MustCompile("(--- )?(PASS)(:)?(.*)")
	errRegex := regexp.MustCompile("(--- )?(FAIL)(:)?(.*)")
	noTestRegex := regexp.MustCompile(`(.*)?(\[no test files\])`)

	outStr := buffer.String()
	outStr = passRegex.ReplaceAllString(outStr, "\033[1;32m$1$2$3$4\033[0m")
	outStr = errRegex.ReplaceAllString(outStr, "\033[1;31m$1$2$3$4\033[0m")
	outStr = noTestRegex.ReplaceAllString(outStr, "\033[1;33m$1$2\033[0m")

	fmt.Print(outStr)

	if err != nil {
		fmt.Print(errBuf.String())
	}

	return err
}

// watchMode Starts a file watcher and runs test when a change is detected
func watchMode(path string, args []string) error {
	w := watcher.New()

	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(1)

	// Only notify rename and move events.
	// w.FilterOps(watcher.Rename, watcher.Move)

	// Only files that match the regular expression during file listings
	// will be watched.
	r := regexp.MustCompile(`\.go$`)
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
				fmt.Println()
				executeTests(path, args)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch recursively for changes.
	if err := w.AddRecursive(strings.Replace(path, "...", "", -1)); err != nil {
		return err
	}

	fmt.Println()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		return err
	}

	return nil
}
