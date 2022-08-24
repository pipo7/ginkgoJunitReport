package testing2_test

import (
	"fmt"
	"log"
	"os"
	"path"

	// using old version as ginkgo.

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"

	//"github.com/pipo7/ginkgoJunitReport/customreporters"

	testing2 "github.com/pipo7/ginkgoJunitReport/testing2"
)

// used for custom reporters when using old code.
const (
	ReportDir    = "outputDir"
	ReportPrefix = "TestReportPrefix"
)

func TestTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing Suite1")
	// Use below code with old version of ginkgo
	// When using ginkgo reporters
	var r []Reporter
	// when using customReporter
	//var r []customreporters.Reporter
	var filePath string
	var nodenum int
	if ReportDir != "" {
		// Create directory and save junit file at this path
		if err := os.MkdirAll(ReportDir, 0755); err != nil {
			log.Fatalf("Failed creating report directory: %v", err)
		} else {
			nodenum = 0
			if config.GinkgoConfig.ParallelTotal > 1 {
				nodenum = config.GinkgoConfig.ParallelNode
			}

			// Use filePath as variable, which can be used later to updates in file.
			filePath = path.Join(ReportDir, fmt.Sprintf("%v_junit_%02d.xml", ReportPrefix, nodenum))
			r = append(r, reporters.NewJUnitReporter(filePath))
			// using customReporter
			//r = append(r, customreporters.NewJUnitReporter(filePath))
		}
	}
	log.Printf("\nStarting E2E run %q on Ginkgo node %d", "Run1", nodenum)
	RunSpecsWithDefaultAndCustomReporters(t, "Kubernetes-E2E-suite", r)

	/*
		// Method1
		filebytes := testing2.ReadXML(filePath)
		testing2.ModifyXML(filePath, filebytes, "JIRA-321") */

	// Method2
	filebytes, err := testing2.ReadTheXML(filePath)
	if err != nil {
		log.Fatalf("Failed reading the XML: %v", err)
	}

	err = testing2.ModifyTheXML(filePath, filebytes)
	if err != nil {
		log.Fatalf("Failed to modify the XML: %v", err)
	}

}

var _ = Describe("Test1 IsPersonAChild()", func() {
	Context("When the person is child then it", func() {
		It("returns True  EPICid JIRAID-1001 xxx", func() {
			person := testing2.Person{
				Age: 10,
			}
			response := person.IsChild()
			fmt.Println("Response is ", response)
			Expect(response).To(BeTrue()) // note here we can also write Expect(person.IsChild()).To(BeTrue())
			// So IsChild mostly user defined method in the file being tested
		})
		fmt.Println("Current spec report example1:", CurrentGinkgoTestDescription())
	})
	Context("When the person is NOT a child then it", func() {
		It("returns False epICID JIRAID-1002  ", func() {
			person := testing2.Person{
				Age: 20,
			}
			response := person.IsChild()
			Expect(response).To(BeFalse())
		})
		fmt.Println("Current spec report example2:", CurrentGinkgoTestDescription())
	})
})
