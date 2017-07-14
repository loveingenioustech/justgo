package main

import (
	"os"
	"fmt"
)

func main() {
	demoDir()

	//demoFileWrite()

	//demoFileRead()

	demoFileRemove()
}

func demoDir(){
	os.Mkdir("text/robin", 0777)
	os.MkdirAll("text/robin/test1/test2", 0777)
	err := os.Remove("text/robin")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("text/robin")
}

func demoFileWrite(){
	userFile := "text/robin.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}
}

func demoFileRead()  {
	userFile := "text/robin.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func demoFileRemove()  {
	userFile := "text/robin.txt"
	err := os.Remove(userFile)

	if err != nil {
		fmt.Println(userFile, err)
		return
	}
}
