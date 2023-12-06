package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const base = "problems"
const bucket = 50

func main() {
	flagProblem := flag.Int("problem", 0, "problem number")
	flag.Parse()
	if flagProblem == nil || *flagProblem <= 0 {
		log.Println("invalid problem number")
		flag.Usage()
		return
	}

	sub := fmt.Sprintf("%03d", (*flagProblem/bucket+1)*bucket)
	project := fmt.Sprintf("%03d", *flagProblem)

	err := os.Mkdir(filepath.Join(base, sub), 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalln("create sub dir", err)
	}
	log.Println("create dir", filepath.Join(base, sub))

	dir := filepath.Join(base, sub, project)
	err = os.Mkdir(dir, 0755)
	if err != nil {
		log.Fatalln("create project dir", err)
	}
	log.Println("create dir", dir)

	version := runtime.Version()
	version = strings.Replace(version, "go", "", 1)

	err = os.WriteFile(filepath.Join(dir, "go.mod"), []byte(fmt.Sprintf(templateMod, *flagProblem, version)), 0664)
	if err != nil {
		log.Fatalln("write go.mod", err)
	}
	log.Println("write file", "go.mod")

	err = os.WriteFile(filepath.Join(dir, "main.go"), []byte(templateMain), 0664)
	if err != nil {
		log.Fatalln("write main.go", err)
	}
	log.Println("write file", "main.go")

	err = exec.Command("go", "work", "use", dir).Run()
	if err != nil {
		log.Fatalln("go work use", err)
	}
	log.Println("go work use", dir)
}

const (
	templateMod = "module projecteuler.net/problem/%d\n" +
		"\n" +
		"go %s"
	templateMain = `package main

func main() {
	
}`
)
