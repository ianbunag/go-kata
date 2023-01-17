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


## Time and space complexities

### Quick guide
| Complexity | Name        | Description                                         | Example                 |
|------------|-------------|-----------------------------------------------------|-------------------------|
| O(1)       | Constant    | Metric is the same regardless of input              | Indexed access          |
| O(log n)   | Logarithmic | Metric is lower with each additional step           | Binary trees            |
| O(n)       | Linear      | Metric depends linearly on the input                | For loop                |
| O(n^2)     | Quadratic   | Metric grows quadratically for each additional step | Nested for loop         |
| O(2^n)     | Exponential | Metric grows exponentially for each additional step | Fibonacci calculation   |
| O(n!)      | Factorial   | Metric grows factorially for each additional step   | Permutation calculation |

### Notes
- Constant complexities may be disregarded, for example:
  - O(2) is considered as O(1)
  - O(2n) is considered as O(n)
  - O(nk) is considered as O(n)
- Space complexity should include everything created in the lifecycle of the
  algorithm:
  - Auxiliary variables (space used while the algorithm is being executed)
  - Output
- Space complexity before `2023-01-17` includes everything in the lifecycle of
  the algorithm
  - Input
  - Auxiliary variables (space used while the algorithm is being executed)
  - Output

### Cheatsheet
[Cheatsheet](https://www.bigocheatsheet.com/)

## Kata websites
- https://www.codewars.com/
- https://neetcode.io/
- https://leetcode.com/
