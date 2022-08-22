package testing2

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type CustomJUnitTestSuite struct {
	XMLName   xml.Name              `xml:"testsuite"`
	TestCases []CustomJUnitTestCase `xml:"testcase"`
	Name      string                `xml:"name,attr"`
	Tests     int                   `xml:"tests,attr"`
	Failures  int                   `xml:"failures,attr"`
	Errors    int                   `xml:"errors,attr"`
	Time      float64               `xml:"time,attr"`
}
type CustomJUnitTestCase struct {
	Name      string `xml:"name,attr"`
	ClassName string `xml:"classname,attr"`
	// Add EPICID as a non-mandatory attribute
	Epicid         string                     `xml:"epicid,attr,omitempty"`
	FailureMessage *CustomJUnitFailureMessage `xml:"failure,omitempty"`
	Skipped        *CustomJUnitSkipped        `xml:"skipped,omitempty"`
	Time           float64                    `xml:"time,attr"`
	SystemOut      string                     `xml:"system-out,omitempty"`
}

type CustomJUnitSkipped struct {
	Message string `xml:",chardata"`
}
type CustomJUnitFailureMessage struct {
	Type    string `xml:"type,attr"`
	Message string `xml:",chardata"`
}

func ReadTheXML(filename string) ([]byte, error) {

	xmlFile, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}
	fmt.Println("Successfully opened the file: ", filename)
	defer xmlFile.Close()
	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return []byte{}, err
	}
	fmt.Println("Successfully read the File as :", string(byteValue))
	return byteValue, nil
}

func ModifyTheXML(filename string, byteValue []byte, tagNewValue string) error {

	myForm := &CustomJUnitTestSuite{}
	err := xml.Unmarshal(byteValue, &myForm)
	if err != nil {
		return err
	} else {
		fmt.Println("Successfully unmarshalled the File in CustomJUnitTestSuite format")
	}

	// Use TestSuite ptr to access slice by index & change value by de-referencing
	for idx := range myForm.TestCases {
		(&myForm.TestCases[idx]).Epicid = tagNewValue
	}

	modifiedXML, err := xml.MarshalIndent(myForm, " ", "  ")
	if err != nil {
		return err
	} else {
		fmt.Println("Successfully marshalled the CustomJUnitTestSuite format to XML")
	}
	// Update modifiedXML with XMLheader
	byteString := []byte(xml.Header + string(modifiedXML))
	err = os.WriteFile(filename, byteString, 0644)
	if err != nil {
		return err
	} else {
		fmt.Println("Successfully saved the modified File")
	}
	return nil
}
