# Go-Test-2

For question 3:
No my running of the two sorting methods does not follow the runtime described in the documentation. The stable sort was able to complete faster than the unstable sort. I'm frankly confused why it did not run with the given runtime as when using a slice of more than 100 elements, there have to be repeats as I limited the range of random integers to be 0 - 100. This would indicate to me that due to the difference in how the stable sort handles swaps is some how more efficient than the unstable sort, as the unstable sort does not give a runtime for its swap function. This may be due to the stable sort keeping the positions the same of two equal values.



To run program:

go build Gotest2Pg1.go
go build Gotest2Pg2.go
./Gotest2Pg1 [int]
./Gotest2Pg2 [int]

Gotest2Pg1 should print a similar output as the following:
Sum:  48533
Normal Runtime:  458ns
Sum:  48533
Parallel Runtime:  68.792µs

GOtest2Pg2 should print an output as such:
Slice Sort Runtime:  61.541µs
Slice Stable Runtime:  5.625µs
