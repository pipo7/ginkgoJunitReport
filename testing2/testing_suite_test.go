package testing2_test

import (
	"fmt"
	"log"
	"os"
	"path"

	// using old version as ginkgo.
	g "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	testing2 "ginkgotutorialOldVersion/testing2"
	"testing"
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
			r = append(r, reporters.NewJUnitReporter(path.Join(ReportDir, fmt.Sprintf("junit_%v%02d.xml", ReportPrefix, config.GinkgoConfig.ParallelNode))))
		}
	}
	log.Printf("\nStarting e2e run %q on Ginkgo node %d", "Run1", config.GinkgoConfig.ParallelNode)
	g.RunSpecsWithDefaultAndCustomReporters(t, "Kubernete-E2E-suite", r)
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
