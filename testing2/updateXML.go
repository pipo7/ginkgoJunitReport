package testing2

import (
	"fmt"
	"io"
	"os"
	S "strings"
)

func ReadXML(filename string) []byte {

	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened ", filename)
	defer xmlFile.Close()
	byteValue, _ := io.ReadAll(xmlFile)
	return byteValue

}

func ModifyXML(filename string, byteValue []byte) {
	xmlString := string(byteValue)
	modifiedXML := S.ReplaceAll(xmlString, " classname", " epicid=JIRAID-1234567 classname")
	fmt.Println("Modified String", modifiedXML)

	err := os.WriteFile(filename, []byte(modifiedXML), 0644)
	if err != nil {
		panic(err)
	}
}
