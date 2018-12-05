# ci-cw-go
This project is an implementation of two Computational Intelligence algorithms to help solve a generic problem of Market Pricing.

NOTE: This is a coursework project.

PricingProblem was supplied by the University in Java, I have translated this code into Go, testing it alongside the Java version for consistency.

## Set Up
This project was developed using Go v1.11.2

The easiest way to run this project is to clone the repository into: 
```bash 
%GOPATH%/src/github.com/aagoldingay/ci-cw-go
```

Then run the code:

```bash 
cd %GOPATH%/src/github.com/aagoldingay/ci-cw-go
go run main.go
```

### Run without output to xlsx
To run this project without xlsx output, ensure lines 71-72 of `main.go` are commented out.
```go
// xlsx output
// xlsxhandler.WriteXLSXParams(finalRevenues, psoPopulation, aisPopulation, aisReplacement, aisClonesFactor)
xlsxhandler.WriteXLSXRevenues(seeds, randomRevenues, psoRevenues, aisRevenues)
```
It is also useful to discard the second return object from each algorithm on lines 57, 61 and 65 as well,  and change the bool value to false, as below,:
```go
finalRevenues[0][i], _ = algorithms.RandomSearch(numGoods, false, &p)
// ...
finalRevenues[1][i], _ = algorithms.PSOSearch(numGoods, psoPopulation, false, &p)
// ...
finalRevenues[2][i], _ = algorithms.AISSearch(numGoods, aisPopulation, aisReplacement, aisClonesFactor, false, &p)
```

### Run on single seed
To run the algorithms on one seed, alter the seeds variable in `main.go`, line 14 to suit your needs.
```go
seeds := []int64{0, 38, 113}
```

### Run one algorithm
To run one algorithm, change lines 16 and 17 of `main.go` as required.
```go
runSingle(numGoods, seeds[0])
// runAll(numGoods, seeds)
```

### Run as steps instead of time limits
This will need more changes to the code, however within `algorithms/algoriths.go`, each algorithm has a for loop as well as more information about the changes required
1. Uncomment the for loop condition
2. Comment out the `timeout` and `tick` variables
3. Comment out switch statement, and move `<-timeout:` code block outside of the loop
4. Move `if trace` condition to the end of the for loop