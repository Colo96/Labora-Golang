package main

import "testing"

func TestRotateRight(t *testing.T) {
	tests := []struct {
		txt      string
		espected string
	}{
		{txt: "abc", espected: "cab"},
		{txt: "cab", espected: "bca"},
		{txt: "bac", espected: "cba"},
		{txt: "acb", espected: "bac"},
	}
	for _, tt := range tests {
		t.Run(tt.txt, func(t *testing.T) {
			txt := RotateRight(tt.txt)
			if txt != tt.espected {
				t.Errorf("Ambos textos no coinciden %s", txt)
			} else {
				t.Log("Los textos coinciden")
			}
		})
	}
}

func BenchmarkRotateRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateRight("ajdjdjfgdjdjsdjdjuuifuuussja")
	}
}
