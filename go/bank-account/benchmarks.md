# without lock balance 

BenchmarkAccountOperations
BenchmarkAccountOperations-12                   35433702                28.43 ns/op            0 B/op          0 allocs/op
BenchmarkAccountOperationsParallel
BenchmarkAccountOperationsParallel-12           10816755               111.1 ns/op             0 B/op          0 allocs/op




# lock balance

goos: darwin
goarch: amd64
pkg: account
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkAccountOperations
BenchmarkAccountOperations-12                   35469026                28.66 ns/op            0 B/op          0 allocs/op
BenchmarkAccountOperationsParallel
BenchmarkAccountOperationsParallel-12           10406052               113.3 ns/op             0 B/op          0 allocs/op
