package generator_test

import (
	"testing"

	"fmt"

	"github.com/tomogoma/generator"
)

func ExampleRandom_GenerateAllChars() {
	r := generator.Random{}
	randChars, _ := r.GenerateAllChars(6)
	fmt.Println(len(randChars))
	// Output: 6
}

func TestGenerateLowerCaseChars(t *testing.T) {
	g := generator.Random{}
	size := 36
	p, err := g.GenerateLowerCaseChars(size)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(p) != size {
		t.Errorf("Expected random bytes of length %d, got %d", size, len(p))
	}
}

func TestGenerateUpperCaseChars(t *testing.T) {
	g := generator.Random{}
	size := 36
	p, err := g.GenerateUpperCaseChars(size)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(p) != size {
		t.Errorf("Expected random bytes of length %d, got %d", size, len(p))
	}
}

func TestGenerateNumberChars(t *testing.T) {
	g := generator.Random{}
	size := 36
	p, err := g.GenerateNumberChars(size)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(p) != size {
		t.Errorf("Expected random bytes of length %d, got %d", size, len(p))
	}
}

func TestGenerateSpecialChars(t *testing.T) {
	g := generator.Random{}
	size := 36
	p, err := g.GenerateSpecialChars(size)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(p) != size {
		t.Errorf("Expected random bytes of length %d, got %d", size, len(p))
	}
}

func TestGenerateAlphabetChars(t *testing.T) {
	g := generator.Random{}
	size := 36
	p, err := g.GenerateAlphabetChars(size)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(p) != size {
		t.Errorf("Expected random bytes of length %d, got %d", size, len(p))
	}
}

func TestGenerateAlphaNumericChars(t *testing.T) {
	g := generator.Random{}
	size := 36
	p, err := g.GenerateAlphaNumericChars(size)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(p) != size {
		t.Errorf("Expected random bytes of length %d, got %d", size, len(p))
	}
}

func TestGenerateAllChars(t *testing.T) {
	g := generator.Random{}
	size := 36
	p, err := g.GenerateAllChars(size)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(p) != size {
		t.Errorf("Expected random bytes of length %d, got %d", size, len(p))
	}
}
