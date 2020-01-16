package seqgen

import "testing"

func BenchmarkNewSeqGenerator(b *testing.B) {
	sg, _ := NewSeqGenerator(1, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sg.NextID()
	}
}
