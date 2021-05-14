package main

import (
	"regexp"
	"strings"
	"testing"
)

var (
	browser = "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
	re      = regexp.MustCompile("Android")
)

// regexp.MatchString компилирует регулярку каждый раз
func BenchmarkRegExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = regexp.MatchString("Android", browser)
	}
}

// используем прекомпилированную регулярку
func BenchmarkRegCompiled(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re.MatchString(browser)
	}
}

// просто ищем вхождение подстроки
func BenchmarkStrContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Contains(browser, "Android")
	}
}
