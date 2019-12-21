package wrand

import "testing"

func BenchmarkTypeRands_GetInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetInt(10000)
	}
}

func TestGetInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			for {
				GetInt(10000)
			}
		}()
	}
	for {
		GetInt(10000)
	}
}
