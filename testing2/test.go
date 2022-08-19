package testing2

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Age int
}

func (p *Person) IsChild() bool {
	return p.Age < 18
}

func updateXML(filename string) {
	// Open our xmlFile
	xmlFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened ", filename)
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)
	fmt.Println(string(byteValue))
}
