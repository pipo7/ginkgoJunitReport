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
	fmt.Println("ReadXML as :", string(byteValue))
	return byteValue

}

func ModifyXML(filename string, byteValue []byte, tagNewValue string) {
	xmlString := string(byteValue)

	replacewith := fmt.Sprintf(" epicid=\"%s\" classname", tagNewValue)
	modifiedXML := S.ReplaceAll(xmlString, " classname", replacewith)

	err := os.WriteFile(filename, []byte(modifiedXML), 0644)
	if err != nil {
		panic(err)
	}
}
