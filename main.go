package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("должно быть 2 аргумента: путь до каталога и имя программы")
	}

	path := os.Args[1]
	progname := os.Args[2]

	cmd := exec.Command(progname)
	cmd.Env = env(path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Запуск программы завершился с ошибкой: %v", err)
	}
}

func env(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var env []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := filepath.Join(path, file.Name())
		data, err := ioutil.ReadFile(name)
		if err != nil {
			continue
		}

		var builder strings.Builder
		builder.WriteString(file.Name())
		builder.WriteRune('=')
		builder.WriteString(string(data))
		env = append(env, builder.String())
	}

	return env
}
