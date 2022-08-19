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
	var r []Reporter
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
	RunSpecsWithDefaultAndCustomReporters(t, "Kubernete-E2E-suite", r)

	testing2.UpdateXML("outputDir/TestReportPrefix_junit_01.xml")
}

var _ = Describe("1PersonIsChild()", func() {
	Context("When the person is child then it", func() {
		It("returns True", func() {
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
		It("returns False", func() {
			person := testing2.Person{
				Age: 20,
			}
			response := person.IsChild()
			Expect(response).To(BeFalse())
		})
		fmt.Println("Current spec report example2:", CurrentGinkgoTestDescription())
	})
})
