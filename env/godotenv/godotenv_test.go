package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func Test_normal(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("age: ", os.Getenv("age"))
}

func Test_Lazy(t *testing.T) {
	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("age: ", os.Getenv("age"))
}

func Test_Alias(t *testing.T) {
	err := godotenv.Load("common", "dev.env")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("version: ", os.Getenv("version"))
	fmt.Println("database: ", os.Getenv("database"))
}

func Test_Yaml(t *testing.T) {
	err := godotenv.Load("k8s.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("version: ", os.Getenv("version"))
}

func Test_ReadWithoutEnv(t *testing.T) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])
}

func Test_String(t *testing.T) {
	content := `
name: awesome web
version: 0.0.1
  `
	myEnv, err := godotenv.Unmarshal(content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])
}

func Test_Reader(t *testing.T) {
	file, _ := os.OpenFile(".env", os.O_RDONLY, 0666)
	myEnv, err := godotenv.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])

	buf := &bytes.Buffer{}
	buf.WriteString("name: awesome web @buffer")
	buf.Write([]byte{'\n'})
	buf.WriteString("version: 0.0.1")
	myEnv, err = godotenv.Parse(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])
}
