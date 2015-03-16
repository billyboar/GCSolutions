package drum

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"testing"
)

func TestDecodeFile(t *testing.T) {
	tData := []struct {
		path   string
		output string
	}{
		{"pattern_1.splice",
			`Saved with HW Version: 0.808-alpha
Tempo: 120
(0) kick	|x---|x---|x---|x---|
(1) snare	|----|x---|----|x---|
(2) clap	|----|x-x-|----|----|
(3) hh-open	|--x-|--x-|x-x-|--x-|
(4) hh-close	|x---|x---|----|x--x|
(5) cowbell	|----|----|--x-|----|
`,
		},
		{"pattern_2.splice",
			`Saved with HW Version: 0.808-alpha
Tempo: 98.4
(0) kick	|x---|----|x---|----|
(1) snare	|----|x---|----|x---|
(3) hh-open	|--x-|--x-|x-x-|--x-|
(5) cowbell	|----|----|x---|----|
`,
		},
		{"pattern_3.splice",
			`Saved with HW Version: 0.808-alpha
Tempo: 118
(40) kick	|x---|----|x---|----|
(1) clap	|----|x---|----|x---|
(3) hh-open	|--x-|--x-|x-x-|--x-|
(5) low-tom	|----|---x|----|----|
(12) mid-tom	|----|----|x---|----|
(9) hi-tom	|----|----|-x--|----|
`,
		},
		{"pattern_4.splice",
			`Saved with HW Version: 0.909
Tempo: 240
(0) SubKick	|----|----|----|----|
(1) Kick	|x---|----|x---|----|
(99) Maracas	|x-x-|x-x-|x-x-|x-x-|
(255) Low Conga	|----|x---|----|x---|
`,
		},
		{"pattern_5.splice",
			`Saved with HW Version: 0.708-alpha
Tempo: 999
(1) Kick	|x---|----|x---|----|
(2) HiHat	|x-x-|x-x-|x-x-|x-x-|
`,
		},
	}

	for _, exp := range tData {
		decoded, err := DecodeFile(path.Join("fixtures", exp.path))
		if err != nil {
			t.Fatalf("something went wrong decoding %s - %v", exp.path, err)
		}
		if fmt.Sprint(decoded) != exp.output {
			t.Logf("decoded:\n%#v\n", fmt.Sprint(decoded))
			t.Logf("expected:\n%#v\n", exp.output)
			t.Fatalf("%s wasn't decoded as expect.\nGot:\n%s\nExpected:\n%s",
				exp.path, decoded, exp.output)
		}
	}
}

func BenchmarkDecodeM(b *testing.B) {
	// Read the entire file into memory so that we don't benchmark the
	// disk IO speed.
	data, err := ioutil.ReadFile("fixtures/pattern_1.splice")
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		_, err = Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
	}
}

var n int = 3

func benchmarkDecode(s string, i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		DecodeFile(s)
	}
}
func BenchmarkDecode1(b *testing.B) {
	benchmarkDecode(path.Join("fixtures", "pattern_1.splice"), n, b)
}

func BenchmarkDecode2(b *testing.B) {
	benchmarkDecode(path.Join("fixtures", "pattern_2.splice"), n, b)
}

func BenchmarkDecode3(b *testing.B) {
	benchmarkDecode(path.Join("fixtures", "pattern_3.splice"), n, b)
}

func BenchmarkDecode4(b *testing.B) {
	benchmarkDecode(path.Join("fixtures", "pattern_4.splice"), n, b)
}

func BenchmarkDecode5(b *testing.B) {
	benchmarkDecode(path.Join("fixtures", "pattern_5.splice"), n, b)
}

/*
func BenchmarkPatternString(b *testing.B) {
	p, err := DecodeFile("fixtures/pattern_1.splice")
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		p.String()
	}
}
*/
