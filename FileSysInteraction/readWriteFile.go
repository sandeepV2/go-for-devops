package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type User struct{
    Name string
    ID int
    err error
}

func getUser(s string) (User, error) {
    sp := strings.Split(s, ":")
    if len(sp) != 2 {
        return User{}, fmt.Errorf("record(%s) was not in the correct format", s)
    }

    id, err := strconv.Atoi(sp[1])
    if err != nil {
        return User{}, fmt.Errorf("record(%s) had non-numeric ID", s)
    }
    return User{Name: strings.TrimSpace(sp[0]), ID: id}, nil
}

func main() {
    d1 := []byte("Text in file\n")
    if err := os.WriteFile("/tmp/tempfile", d1, 0644); err != nil {
        panic(err)
    }
	if data, err := os.ReadFile("/tmp/tempfile"); err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}

	f, err := os.Open("/tmp/tempfile")
	if err != nil {
    	panic(err)
	}
	if bytes, err := io.Copy(os.Stdout, f); err != nil {
    	panic(err)
	} else {
		fmt.Println(bytes)
	}
}