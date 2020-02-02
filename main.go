package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	defaultEditor = "vim"
	mainFileName  = "main.go"
	modFileName   = "go.mod"

	mainBase = `package main

func main() {
}`

	modBase = `module %s

go %s`
)

type cli struct {
	stdout io.Writer
	stderr io.Writer
}

func main() {
	c := &cli{
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	os.Exit(c.run())
}

func (c *cli) run() int {
	// Creates needed files.
	tmpDir, err := ioutil.TempDir("", "gotmp")
	if err != nil {
		fmt.Fprintln(c.stderr, err)
		return 1
	}
	defer os.RemoveAll(tmpDir)
	mainfile, err := create(filepath.Join(tmpDir, mainFileName), mainBase)
	if err != nil {
		fmt.Fprintln(c.stderr, err)
		return 1
	}
	// TODO: Make sure to detect go version and package name
	_, err = create(filepath.Join(tmpDir, modFileName), fmt.Sprintf(modBase, "gotmp", "1.13"))
	if err != nil {
		fmt.Fprintln(c.stderr, err)
		return 1
	}

	// TODO: Make it possible to change the editor.
	// Edit with the preferred editor.
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = defaultEditor
	}
	executable, err := exec.LookPath(editor)
	if err != nil {
		fmt.Fprintln(c.stderr, err)
		return 1
	}
	cmd := exec.Command(executable, mainfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(c.stderr, err)
		return 1
	}

	return 0
}

func create(path, data string) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if _, err = f.Write([]byte(data)); err != nil {
		return nil, err
	}
	return f, nil
}
