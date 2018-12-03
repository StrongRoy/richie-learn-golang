package tests

import "testing"

func BenchmarkAll(b *testing.B)  {
	for n:=0;n < b.N;n++{
		Print1to20()
	}
}