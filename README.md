# JSON Unmarshal Memory Usage

This repository contains scripts to analyze memory usage of `json.Unmarshal`.

### To run benchmarks

```bash
go test -bench=. -memprofile=m.p json_bench_test.go
```

> To view memory profile run `go tool pprof  -png m.p > out.png`

`main.go` contains code to see the memory impact of `json.Unmarshal` on `string`, `slice` and `object` of varying lengths.

Results from running above tests can be found [here](https://docs.google.com/spreadsheets/d/1HlYeD-kYBoGrPjAMCXUu5ObC-VGHb5EX1Gh4QlInSpo/edit?usp=sharing).