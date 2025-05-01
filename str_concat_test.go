package bench

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var result string

func BenchmarkConcatStrings_Plus(b *testing.B) {
	var r string
	a, c := "foo", "bar"
	for i := 0; i < b.N; i++ {
		r = a + c
	}
	result = r
}

func BenchmarkConcatStrings_Sprintf(b *testing.B) {
	var r string
	a, c := "foo", "bar"
	for i := 0; i < b.N; i++ {
		r = fmt.Sprintf("%s%s", a, c)
	}
	result = r
}

func BenchmarkConcatStrings_StringsBuilder(b *testing.B) {
	var r string
	a, c := "foo", "bar"
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.WriteString(a)
		sb.WriteString(c)
		r = sb.String()
	}
	result = r
}

func BenchmarkConcatStrings_Join(b *testing.B) {
	var r string
	a, c := "foo", "bar"
	for i := 0; i < b.N; i++ {
		r = strings.Join([]string{a, c}, "")
	}
	result = r
}

func BenchmarkConcatStringAndInt_Plus(b *testing.B) {
	var r string
	a := "foo"
	n := 42
	for i := 0; i < b.N; i++ {
		r = a + strconv.Itoa(n)
	}
	result = r
}

func BenchmarkConcatStringAndInt_Sprintf(b *testing.B) {
	var r string
	a := "foo"
	n := 42
	for i := 0; i < b.N; i++ {
		r = fmt.Sprintf("%s%d", a, n)
	}
	result = r
}

func BenchmarkConcatStringAndInt_StringsBuilder(b *testing.B) {
	var r string
	a := "foo"
	n := 42
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.WriteString(a)
		sb.WriteString(strconv.Itoa(n))
		r = sb.String()
	}
	result = r
}

func BenchmarkConcatStringAndInt_Join(b *testing.B) {
	var r string
	a := "foo"
	n := 42
	for i := 0; i < b.N; i++ {
		r = strings.Join([]string{a, strconv.Itoa(n)}, "")
	}
	result = r
}
