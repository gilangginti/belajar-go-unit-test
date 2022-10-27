package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE TEST")
	m.Run()
	fmt.Println("AFTER TEST")
}

func TestHelloWorldWindows(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run in windows")
	}
	result := HelloWorld("Salva")
	require.Equal(t, "Hello Salva", result)
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Salva")
	// assert = fail
	assert.Equal(t, "Hello Salva", result)
	// require = failnow
	require.Equal(t, "Hello Salva", result)

	fmt.Println("TEST DONE")
}

func TestHelloWorld(t *testing.T) {
	myFunc := HelloWorld("Gilang")
	if myFunc != "Hello Gilang" {
		// CODE SELANJUTNYA AKAN DIEKSEKUSI
		// t.Fail()
		// BISA MEMBERIKAN INFORMASI
		// t.Error("TEST MUST BE 'Hello Gilang'")

		// KODE SELANJUTNYA TIDAK DIEKSEKUSI
		// t.FailNow()
		// BISA MEMBERIKAN INFORMASI
		t.Fatal("TEST MUST BE 'Hello Gilang'")

	}
	fmt.Println("Test DOne")
}
