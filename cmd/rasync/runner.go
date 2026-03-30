package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func download(local string, remote string) {
	var err error

	err = os.MkdirAll(filepath.Dir(local), os.ModePerm)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] %s -> %s (os.MkdirAll: %v)\n", remote, local, err)
		return
	}

	// Download file
	resp, err := http.Get(remote)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("[FAIL] %s -> %s (http.Get: %v)\n", remote, local, err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(local)
	if err != nil {
		fmt.Printf("[FAIL] %s -> %s (os.Create: %v)\n", remote, local, err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("[FAIL] %s -> %s (write error: %v)\n", remote, local, err)
		return
	}

	fmt.Printf("[OK] %s -> %s\n", remote, local)
}

func main() {
	var err error
	var wg sync.WaitGroup

	file, err := os.Open(".rasync")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open .rasync (%v)\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber += 1
		line := strings.TrimSpace(scanner.Text())

		if line == "" || line[0] == '#' {
			continue
		}

		parts := strings.SplitN(line, " ", 3)

		if len(parts) < 2 || (len(parts) == 3 && parts[2][0] != '#') {
			fmt.Fprintf(os.Stderr, ".rasync:%d - invalid line format (expected: [local] [remote] (# comment))\n", lineNumber)
			fmt.Fprintf(os.Stderr, "             [] represent required tokens\n")
			fmt.Fprintf(os.Stderr, "             () represent optional tokens\n")
			os.Exit(1)
		}

		local, remote := parts[0], parts[1]

		wg.Go(func() {
			download(local, remote)
		})
	}

	err = scanner.Err()

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to scan .rasync (%v)\n", err)
		os.Exit(1)
	}

	wg.Wait()
}
