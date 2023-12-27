package parse

import (
	"os"
	"testing"
)

func TestSplit(t *testing.T) {
	expectedResult := []string{"url", "its.me.again"}
	output := split("url=its.me.again")

	if expectedResult[0] != output[0] && expectedResult[1] != output[1] {
		t.Errorf("got %s, wanted %s", output, expectedResult)
	}
}

func TestSplitWithEqual(t *testing.T) {
	expectedResult := []string{"password", "n8^gdhT++="}
	output := split("password=n8^gdhT++=")

	if expectedResult[0] != output[0] && expectedResult[1] != output[1] {
		t.Errorf("got %s, wanted %s", output, expectedResult)
	}
}

func TestSplitWithEqualAtStarting(t *testing.T) {
	expectedResult := []string{"password", "=n8^gdhT++="}
	output := split("password==n8^gdhT++=")

	if expectedResult[0] != output[0] || expectedResult[1] != output[1] {
		t.Errorf("got %s, wanted %s", output, expectedResult)
	}
}

func TestSplitWithEqualEmpty(t *testing.T) {
	output := split("")
	if len(output) != 0 {
		t.Error("output should be empty")
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		split("password=n8^gdhT++=")
	}
}

func TestParse(t *testing.T) {
	expectedResult := make(map[string]string)
	expectedResult["url"] = "test.postgres.com"
	expectedResult["port"] = "5432"
	expectedResult["password"] = "test"

	byt, err := os.ReadFile("config_test.ini")
	if err != nil {
		t.Errorf("Fail to read the ini file")
	}
	output := parse(byt)

	if len(expectedResult) != len(output) {
		t.Errorf("got %s, wanted %s", output, expectedResult)
	}
}

func BenchmarkParse(b *testing.B) {
	byt, err := os.ReadFile("config_test.ini")
	if err != nil {
		b.Errorf("Fail to read the ini file")
	}
	for i := 0; i < b.N; i++ {
		parse(byt)
	}
}
