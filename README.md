# Go Kata
Collection of Go kata challenges

## Install test executable
```sh
go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
```

## Bootstrap kata
```sh
./scaffold.sh codewars get_sum
cd codewars/1_get_sum
ginkgo bootstrap
ginkgo generate
```

## Run tests
```sh
# Run all tests
ginkgo ./...

# Run specific tests matching description
ginkgo --focus "should reverse word" ./...

# Run specific tests matching filename
ginkgo --focus-file codewars/1_reverse_words/ReverseWord_test.go ./...
```