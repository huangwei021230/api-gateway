package performance_test

import (
	"bytes"
	"net/http"
	"runtime"
	"testing"
)

func BenchmarkDivision(b *testing.B) {
	destUrl := "http://127.0.0.1:8888/div"
	contentType := "application/json"
	reqBody := `{"FirstNum":"100","SecondNum":"20"}`
	for i := 0; i < b.N; i++ {
		_, err := http.Post(destUrl, contentType, bytes.NewBufferString(reqBody))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDivisionParallel(b *testing.B) {
	destUrl := "http://127.0.0.1:8888/div"
	contentType := "application/json"
	reqBody := `{"FirstNum":"100","SecondNum":"20"}`
	runtime.GOMAXPROCS(8)
	for i := 0; i < b.N; i++ {
		_, err := http.Post(destUrl, contentType, bytes.NewBufferString(reqBody))
		if err != nil {
			b.Fatal(err)
		}
	}
}
