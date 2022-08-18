package testing2_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	// using old version as ginkgo.
	g "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"

	testing2 "github.com/pipo7/ginkgoJunitReport/testing2"
)

// used for custom reporters when using old code.
const (
	ReportDir    = "outputDir"
	ReportPrefix = "TestReportPrefix"
)

func TestTesting(t *testing.T) {
	RegisterFailHandler(g.Fail)
	g.RunSpecs(t, "Testing Suite1")
	// Use below code with old version of ginkgo
	var r []g.Reporter
	if ReportDir != "" {
		// TODO: we should probably only be trying to create this directory once
		// rather than once-per-Ginkgo-node.
		if err := os.MkdirAll(ReportDir, 0755); err != nil {
			log.Fatalf("Failed creating report directory: %v", err)
		} else {
			r = append(r, reporters.NewJUnitReporter(path.Join(ReportDir, fmt.Sprintf("%v_junit_%02d.xml", ReportPrefix, config.GinkgoConfig.ParallelNode))))
		}
	}
	log.Printf("\nStarting e2e run %q on Ginkgo node %d", "Run1", config.GinkgoConfig.ParallelNode)
	g.RunSpecsWithDefaultAndCustomReporters(t, "Kubernete-E2E-suite", r)

	// Open our xmlFile
	xmlFile, err := os.Open("outputDir/TestReportPrefix_junit_01.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened junit_TestReportPrefix01.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)
	fmt.Println(string(byteValue))
}

// ReportAfterSuite is called exactly once at the end of the suite after any AfterSuite nodes have run
// IMPORTANT ReportAfterSuite only runs on process #1 and receives a Report that aggregates the SpecReports from all processes.
var _ = g.AfterSuite(func() {
	// process report
	log.Println("Printing All the reports after suite")
})

var _ = g.Describe("1PersonIsChild()", func() {
	g.Context("When the person is child", func() {
		g.It("returns True", func() {
			person := testing2.Person{
				Age: 10,
			}
			response := person.IsChild()
			fmt.Println("response is ", response)
			Expect(response).To(BeTrue()) // note here we can also write Expect(person.IsChild()).To(BeTrue())
			// So IsChild mostly user defined method in the file being tested
		})
		fmt.Println("Current spec report example:", g.CurrentGinkgoTestDescription())
	})
	g.Context("When the person is NOT a child", func() {
		g.It("returns True", func() {
			person := testing2.Person{
				Age: 20,
			}
			response := person.IsChild()
			Expect(response).To(BeFalse())
		})
		fmt.Println("Current spec report example:", g.CurrentGinkgoTestDescription())
	})
})
