package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	env, _ := ReadDir("./config")

	RunCmd(os.Args, env)
}

func ReadDir(dir string) ([]string, error) {
	var env []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range files {
		file, err := os.Open("./config/" + fileInfo.Name())

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			env = append(env, scanner.Text())
		}

		_ = file.Close()
	}

	return env, nil
}

func RunCmd(cmd []string, env []string) int {
	path := cmd[1]
	args := cmd[2:]

	command := exec.Command(path, args...)
	command.Env = append(os.Environ(), env...)
	command.Stdout = os.Stdout

	if err := command.Run(); err != nil {
		log.Fatal(err)

		return 1
	}

	return 0
}
