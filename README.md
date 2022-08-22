## Run ginkgo
// READ on https://github.com/onsi/ginkgo
and https://golang.hotexamples.com/examples/github.com.onsi.ginkgo/-/RunSpecsWithDefaultAndCustomReporters/golang-runspecswithdefaultandcustomreporters-function-examples.html


# go to testing directory and run
cd testing &&  ginkgo bootstrap

## To generate tests use
cd testing && ginkgo generate set

# You can run tests from vscode or in terminal enter command
ginkgo -v
# using slowSpecThreshold
ginkgo -slowSpecThreshold='5.0' -v  

# Shows example of 
generate custom report using RunSpecsWithDefaultAndCustomReporters
