# Test Harness

## Setup

From the root of repository, run (e.g. in Git Bash):

```sh
cd generator
go mod init gentests
go mod tidy
cd ../testharness
go mod init testharness
```

If you get some errors, might also need to run
```sh
go get gonum.org/v1/gonum/mat
```

## Build

From the root of repository, run:

```sh
cd generator
go build
cd ../testharness
go build
```

This will build your go programs in /gentests and /testharness folders as gentests(.exe) and testharness(.exe)

## Run

To generate **new matrices.go and results.go**, from the root of repository, run:

```sh
cd generator
./gentests.exe #'./gentests' on Linux
```

To run the test harness, from the root of repository, run:

```sh
cd testharness
./testharness.exe #./testharness on Linux
```

**NOTE:** Make sure you uncomment the Naive implementation of multiply() in testHarness.go before build/run (and comment the empty one)

## GenTests

This script generates the matrices that we're going to use for the competition.
On the day of the competition, we will run the script once, and run all of our algorithms on the same matrices.

The script generates two files in testharness/matrices folder:
- matrices.go which contains 12 matrices (M0 to M11)
- results.go which contains 6 matrices (R01 to R1011)

RAB is the result of MA x MB
We need the result to check that our algorithms produce the correct result

There are 3 constants at the beginning of the script: small, medium and large.
These are integers. They have to be even.
We need to agree on 3 values for these constants so that all of our tests run in about 2-3mins

The matrices are generated as follows:

M0 is of size (small\*2 x small/2)<br/>
M1 is of size (small/2 x small\*2)<br/>
M2 is of size (small x small)<br/>
M3 is of size (small x small)<br/>

M4 is of size (medium\*2 x medium/2)<br/>
M5 is of size (medium/2 x medium\*2)<br/>
M6 is of size (medium x medium)<br/>
M7 is of size (medium x medium)<br/>

M8 is of size (large\*2 x large/2)<br/>
M9 is of size (large/2 x large\*2)<br/>
M10 is of size (large x large)<br/>
M11 is of size (large x large)<br/>

Repository contains some pre-generated matrices.go and results.go for convenience

## TestHarness

This script will run 6 tests, each test will be **repeated 10 times** and the **average time** will be calculated:

M0 x M1<br/>
M2 x M3<br/>
M4 x M5<br/>
M6 x M7<br/>
M8 x M9<br/>
M10 x M11<br/>

It will output the average time for each test as well as the total time (sum of all the average times of each test)
If the result produced by your algorithm does not match the result matrix, it will tell you which test failed

**--- PUT YOUR ALGORITHM IN THE multiply() FUNCTION ---**

You can create and call any additional function you'd like, but the time is taken before and after multiply()
