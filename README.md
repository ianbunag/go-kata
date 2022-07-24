# Go Kata
Collection of Go kata challenges

## Run tests
```sh
# Run all tests
ginkgo ./...

# Run specific tests matching description
ginkgo --focus "should reverse word" ./...

# Run specific tests matching filename
ginkgo --focus-file codewars/1_reverse_words/ReverseWord_test.go ./...
```