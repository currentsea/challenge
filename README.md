# Coding Challenge - Uphold

### Compiling and Building the project 
You can compile and run the project with the following: 

```
$ go build 
$ go run CodingChallenge
James Bond 007
PI=03.14
It's 03:13pm
It's 12:13pm
000099UR001337
```

### Runtime Complexity
When referencing my following function `codingChallenge` as shown below, the runtime can be analyzed as follows:

```go
func codingChallenge(input string, padding int) string {
    re := regexp.MustCompile("[0-9]+")
    numbers := re.FindAllString(input, -1)
    for i := 0; i < len(numbers); i++ {
        newString := numbers[i]
        if padding > len(numbers[i]) {
            diff := padding - len(numbers[i])
            for x := 0; x < diff; x++ {
                newString = "0" + newString
            }
        }
        input = strings.Replace(input, numbers[i], newString, 1)
    }
    return input
}
```

* Finding numbers in the input string: The regular expression "[0-9]+" is used to find all numbers in the input string. Finding all occurrences of numbers in a string using regular expressions typically has a linear time complexity O(n), where n is the length of the input string.

* Looping over each found number: The function then iterates over each found number in the numbers slice. In the worst case, if all characters in the input string are numbers, the loop will iterate through all characters once. Therefore, the time complexity of this part is also O(n), where n is the length of the input string.

* Padding numbers: If the padding value is greater than the length of the current number, the function pads the number with leading zeros. The padding operation takes O(padding) time.

* Replacing numbers in the input string: For each number found, the function replaces it in the input string with the padded version. This replacement operation may involve traversing the input string to find the occurrence of the number, and replacing it. In the worst case, if each number is unique in the input string, this operation can take O(n) time.

* Combining all these operations, the overall time complexity of the function is dominated by the regular expression matching and the loop iteration, resulting in O(n), where n is the length of the input string.

### Unit Tests
You can run the unit tests with the following command: 

```
$ go test -v ./.
$ go test -v ./.
=== RUN   TestProgramCase1
--- PASS: TestProgramCase1 (0.00s)
=== RUN   TestProgramCase2
--- PASS: TestProgramCase2 (0.00s)
=== RUN   TestProgramCase3
--- PASS: TestProgramCase3 (0.00s)
=== RUN   TestProgramCase4
--- PASS: TestProgramCase4 (0.00s)
=== RUN   TestProgramCase5
--- PASS: TestProgramCase5 (0.00s)
PASS
ok      awesomeProject  0.087s
```