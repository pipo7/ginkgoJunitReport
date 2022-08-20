package testing2

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type JUnitTestSuite struct {
	XMLName   xml.Name        `xml:"testsuite"`
	TestCases []JUnitTestCase `xml:"testcase"`
	Name      string          `xml:"name,attr"`
	Tests     int             `xml:"tests,attr"`
	Failures  int             `xml:"failures,attr"`
	Errors    int             `xml:"errors,attr"`
	Time      float64         `xml:"time,attr"`
}
type JUnitTestCase struct {
	Name           string               `xml:"name,attr"`
	ClassName      string               `xml:"classname,attr"`
	Epicid         string               `xml:"epicid,attr"`
	FailureMessage *JUnitFailureMessage `xml:"failure,omitempty"`
	Skipped        *JUnitSkipped        `xml:"skipped,omitempty"`
	Time           float64              `xml:"time,attr"`
	SystemOut      string               `xml:"system-out,omitempty"`
}
type JUnitSkipped struct {
	Message string `xml:",chardata"`
}
type JUnitFailureMessage struct {
	Type    string `xml:"type,attr"`
	Message string `xml:",chardata"`
}

func ReadTheXML(filename string) []byte {

	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened ", filename)
	defer xmlFile.Close()
	byteValue, _ := io.ReadAll(xmlFile)
	return byteValue

}

/*
func ModifyTheXML(filename string, byteValue []byte) {
	xmlString := string(byteValue)
	modifiedXML := S.ReplaceAll(xmlString, " classname", " epicid=JIRAID-1234567 classname")
	fmt.Println("Modified String", modifiedXML)

	err := os.WriteFile(filename, []byte(modifiedXML), 0644)
	if err != nil {
		panic(err)
	}
} */
