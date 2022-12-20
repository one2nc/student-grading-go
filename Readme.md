# Problem statement
This repo contains an exercise related to grading the students from various universities. Your job is to fork the repository,
set it up locally, and make the test cases pass in your own repo.

Here are some more details about the problem statement.

The `student-grading-go/grades.csv` contains student records. For each student, we know the following
- first name of the student (column `FirstName`)
- last name of the student (column `LastName`)
- university of the student (column `University`)
- student's score in first test (column `Test1`)
- student's score in second test (column `Test2`)
- student's score in third test (column `Test3`)
- student's score in fourth test (column `Test4`)

There are a total of 30 students in the csv. The first row in the csv file is a header row indicating the order of the fields in the csv.

We need to calculate the final score and grade for each of the students. The scoring is done as follows:

- Final score is the average of all test scores. i.e. if test1=50, test2=60, test3=65 and test4=45, then final score = (50 + 60 + 65 + 45)/4, i.e. 55
- Grade is based on the final score.
    - If final score is < 35, then student is graded as F (failed)
    - If final score is >= 35 and < 50, then student is graded as C
    - If final score is >= 50 and < 70, then student is graded as B
    - If final score is >= 70, then student is graded as A
- Thus, for a student with final score as 55, the grade will be B

As part of this exercise, you need to find out answers to the three questions (written as test cases in `main_test.go` file). These questions are:
1. Calculate each student's grade and final score
2. Find out overall topper student
3. Find out topper student for each university

For this, first, you must read and parse the `grades.csv` file into a `[]student`. This is mentioned in the first test case in the `main_test.go` - `TestParseCSV`.

# Instructions
1. You are expected to fork this repo
2. Set up the repo locally (check the pre-requisites section for software). Run `make setup` to download the dependencies. Run `make test` to run the test cases.
3. Write code in `main.go` file such that each of the test cases pass. You are NOT supposed to modify any code in test case files (i.e. `main_test.go` file)
4. You are not supposed to modify the grades.csv file. The program should return appropriate results so that the test cases pass.
5. Once you write the code and all the test cases pass, push your code to your forked repo on GitHub. Reply to the email you have received from One2N with the following:

   i) Link to GitHub URL of your forked repo

   ii) A screenshot of the test cases passing (you can take the screenshot from an IDE or a command line)
6. For writing code, please use Go 1.19 or above. Write idiomatic Go code. Handle errors and edge cases appropriately. (You will get extra points for this.)
7. The test cases use `github.com/stretchr/testify/assert` library for assertions. Please read the documentation of this library in case you are not familiar.
8. If something is unclear, try to figure out the solution, given the test cases. If something is still unclear, assume what you think is right and make a note in the readme.

# Prerequisites
Make your you have following prerequisite software installed on your laptop.

1. Go 1.19 or above
2. Please use any good code editor (VS Code, vim, etc) or an IDE (GoLand) for working on this codebase
3. Ensure you have GNU `make` installed.
4. Ensure you run gofmt or appropriate code formatting and linting tools to make the code legible.

# Evaluation Criteria
1. All the test cases must pass (no modification done to the test cases code). Share a screenshot of a gif or a video demo (< 1min) showing that the test cases pass.
2. The code follows Golang coding naming and formatting conventions.
3. Bonus points if you can handle error and edge cases.
4. Your code is well formatted and documented. If there are any assumptions, please add them to the project readme.

# Running the code
1. Clone this repository on your local machine
2. `cd` into the root directory of the project (i.e.`student-grading-go` directory)
3. Run `go mod tidy`. This will download the required packages as dependencies.
4. You should be able to run the test cases using `go test .`. This should generate an output similar to below (output below is redacted)
   ```shell
    --- FAIL: TestParseCSV (0.00s)
    main_test.go:12:
    Error Trace:	/Users/chinmay/work/one2n/gitrepos/student-grading-go/main_test.go:12
    Error:      	Not equal:
    expected: 30
    actual  : 0
    Test:       	TestParseCSV
    Messages:   	Student list size should be 30
    panic: runtime error: index out of range [0] with length 0 [recovered]
    panic: runtime error: index out of range [0] with length 0
    
    goroutine 19 [running]:
    testing.tRunner.func1.2({0x102517e20, 0x1400012a270})
    /opt/homebrew/Cellar/go/1.19.3/libexec/src/testing/testing.go:1396 +0x1c8
    testing.tRunner.func1()
    /opt/homebrew/Cellar/go/1.19.3/libexec/src/testing/testing.go:1399 +0x378
    panic({0x102517e20, 0x1400012a270})
    /opt/homebrew/Cellar/go/1.19.3/libexec/src/runtime/panic.go:884 +0x204
    github.com/one2nc/student-grading-go.TestParseCSV(0x0?)
    /Users/chinmay/work/one2n/gitrepos/student-grading-go/main_test.go:15 +0xe0
    testing.tRunner(0x140001071e0, 0x102529100)
    /opt/homebrew/Cellar/go/1.19.3/libexec/src/testing/testing.go:1446 +0x10c
    created by testing.(*T).Run
    /opt/homebrew/Cellar/go/1.19.3/libexec/src/testing/testing.go:1493 +0x300
    FAIL	github.com/one2nc/student-grading-go	0.343s
    FAIL
   ```

# FAQs
1. I am unable to set up the project locally. What do I do?

You should check the pre-requisite section and ensure you have the necessary software installed and configured properly. If you are stuck in installing or configuring Go, VSCode, or GoLand, please refer to official documentation.

2. How do I run tests?

You can use `go test .` command from the project root directory

3. I forked the repo and set it up locally, but I am getting compile time error. What do I do?

You need to identify and fix the compiler errors. After fixing all the compiler errors, when you try running the test cases for the first time using `go test .`, the tests will fail too. This means, you are on right track and you need to figure out how to write the code to make the test cases pass.

4. All my test cases are passing, how do I submit the solution to this exercise?

Create a short video demo (< 1min) demonstrating your code and the passing tests. Push the demo video in your forked repository and send us a link to your solution by email.