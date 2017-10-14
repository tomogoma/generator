package generator_test

import (
	"strings"
	"testing"

	"fmt"

	"github.com/tomogoma/generator"
)

func ExampleCharSet_SecureRandomBytes() {
	charSet, _ := generator.NewCharSet(generator.AllChars)
	randChars, _ := charSet.SecureRandomBytes(15)
	fmt.Println(len(randChars))
	// Output: 15
}

func TestNewCharSet(t *testing.T) {
	type testcase struct {
		desc    string
		charSet string
		expErr  bool
	}

	testcases := []testcase{
		{
			desc:    "good characterset",
			charSet: generator.AllChars,
			expErr:  false,
		},
		{
			desc:    "empty characterset",
			charSet: "",
			expErr:  true,
		},
		{
			desc:    "too long a characterset",
			charSet: strings.Repeat("a", 257),
			expErr:  true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			g, err := generator.NewCharSet(tc.charSet)
			if tc.expErr {
				if err == nil {
					t.Fatal("Expected an error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Got error: %v", err)
			}
			if g == nil {
				t.Error("Got nil generator")
			}
		})
	}
}

func TestSecureRandomBytes(t *testing.T) {
	g, err := generator.NewCharSet(generator.AllChars)
	if err != nil {
		t.Fatalf("Error setting up: new char set: %v", err)
	}
	size := 36
	p, err := g.SecureRandomBytes(size)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(p) != size {
		t.Errorf("Expected random bytes of length %d, got %d", size, len(p))
	}
}
