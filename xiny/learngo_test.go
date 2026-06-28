package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLearnMultiple(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		wantSum  int
		wantProd int
	}{
		{"positive values", 3, 4, 7, 12},
		{"zero values", 0, 0, 0, 0},
		{"negative values", -2, -3, -5, 6},
		{"mixed signs", -5, 3, -2, -15},
		{"one zero", 7, 0, 7, 0},
		{"identity multiply", 1, 99, 100, 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, prod := learnMultiple(tt.x, tt.y)
			if sum != tt.wantSum {
				t.Errorf("learnMultiple(%d, %d) sum = %d, want %d", tt.x, tt.y, sum, tt.wantSum)
			}
			if prod != tt.wantProd {
				t.Errorf("learnMultiple(%d, %d) prod = %d, want %d", tt.x, tt.y, prod, tt.wantProd)
			}
		})
	}
}

func TestLearnNamedReturns(t *testing.T) {
	tests := []struct {
		name string
		x, y int
		want int
	}{
		{"positive values", 3, 4, 12},
		{"zero product", 5, 0, 0},
		{"negative values", -2, -3, 6},
		{"mixed signs", -4, 5, -20},
		{"identity", 1, 42, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := learnNamedReturns(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("learnNamedReturns(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestExpensiveComputation(t *testing.T) {
	result := expensiveComputation()
	if result <= 0 {
		t.Errorf("expensiveComputation() = %d, expected positive value", result)
	}
	// math.Exp(10) ≈ 22026.47, int truncation gives 22026
	expected := 22026
	if result != expected {
		t.Errorf("expensiveComputation() = %d, want %d", result, expected)
	}
}

func TestLearnMemory(t *testing.T) {
	p, q := learnMemory()
	if p == nil {
		t.Fatal("learnMemory() first pointer is nil")
	}
	if q == nil {
		t.Fatal("learnMemory() second pointer is nil")
	}
	if *p != 7 {
		t.Errorf("learnMemory() *p = %d, want 7", *p)
	}
	if *q != -2 {
		t.Errorf("learnMemory() *q = %d, want -2", *q)
	}
}

func TestSentenceFactory(t *testing.T) {
	tests := []struct {
		name           string
		middle         string
		before, after  string
		want           string
	}{
		{"basic sentence", "summer", "A beautiful", "day!", "A beautiful summer day!"},
		{"empty middle", "", "Hello", "world", "Hello  world"},
		{"single words", "bright", "very", "morning", "very bright morning"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := sentenceFactory(tt.middle)
			got := fn(tt.before, tt.after)
			if got != tt.want {
				t.Errorf("sentenceFactory(%q)(%q, %q) = %q, want %q",
					tt.middle, tt.before, tt.after, got, tt.want)
			}
		})
	}
}

func TestSentenceFactoryClosure(t *testing.T) {
	// Verify the factory returns a reusable closure that captures state.
	factory := sentenceFactory("winter")
	r1 := factory("A cold", "night")
	r2 := factory("A snowy", "morning")
	if r1 != "A cold winter night" {
		t.Errorf("first call = %q, want %q", r1, "A cold winter night")
	}
	if r2 != "A snowy winter morning" {
		t.Errorf("second call = %q, want %q", r2, "A snowy winter morning")
	}
}

func TestCountdown(t *testing.T) {
	var got []int
	for n := range countdown(5) {
		got = append(got, n)
	}
	want := []int{5, 4, 3, 2, 1}
	if len(got) != len(want) {
		t.Fatalf("countdown(5) yielded %d values, want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("countdown(5)[%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestCountdownZero(t *testing.T) {
	var got []int
	for n := range countdown(0) {
		got = append(got, n)
	}
	if len(got) != 0 {
		t.Errorf("countdown(0) yielded %d values, want 0", len(got))
	}
}

func TestCountdownEarlyBreak(t *testing.T) {
	var got []int
	for n := range countdown(10) {
		got = append(got, n)
		if n == 8 {
			break
		}
	}
	want := []int{10, 9, 8}
	if len(got) != len(want) {
		t.Fatalf("early break yielded %d values, want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("early break [%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestLearnGenerics(t *testing.T) {
	t.Run("int max", func(t *testing.T) {
		if got := learnGenerics(10, 20); got != 20 {
			t.Errorf("learnGenerics(10, 20) = %d, want 20", got)
		}
	})
	t.Run("int equal", func(t *testing.T) {
		if got := learnGenerics(5, 5); got != 5 {
			t.Errorf("learnGenerics(5, 5) = %d, want 5", got)
		}
	})
	t.Run("int first larger", func(t *testing.T) {
		if got := learnGenerics(99, 1); got != 99 {
			t.Errorf("learnGenerics(99, 1) = %d, want 99", got)
		}
	})
	t.Run("negative ints", func(t *testing.T) {
		if got := learnGenerics(-1, -10); got != -1 {
			t.Errorf("learnGenerics(-1, -10) = %d, want -1", got)
		}
	})
	t.Run("string comparison", func(t *testing.T) {
		if got := learnGenerics("apple", "banana"); got != "banana" {
			t.Errorf("learnGenerics(apple, banana) = %q, want banana", got)
		}
	})
	t.Run("float64 comparison", func(t *testing.T) {
		if got := learnGenerics(1.5, 2.5); got != 2.5 {
			t.Errorf("learnGenerics(1.5, 2.5) = %f, want 2.5", got)
		}
	})
}

func TestPairString(t *testing.T) {
	tests := []struct {
		name string
		p    pair
		want string
	}{
		{"positive pair", pair{3, 4}, "(3, 4)"},
		{"zero pair", pair{0, 0}, "(0, 0)"},
		{"negative pair", pair{-1, -2}, "(-1, -2)"},
		{"mixed pair", pair{-5, 10}, "(-5, 10)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("pair%v.String() = %q, want %q", tt.p, got, tt.want)
			}
		})
	}
}

func TestPairServeHTTP(t *testing.T) {
	p := pair{3, 4}
	req := httptest.NewRequest(http.MethodGet, "/pair", nil)
	w := httptest.NewRecorder()

	p.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("ServeHTTP status = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	body := w.Body.String()
	expected := "Pair values: (3, 4)"
	if body != expected {
		t.Errorf("ServeHTTP body = %q, want %q", body, expected)
	}
}

func TestPairServeHTTPDifferentValues(t *testing.T) {
	p := pair{-7, 42}
	req := httptest.NewRequest(http.MethodGet, "/pair", nil)
	w := httptest.NewRecorder()

	p.ServeHTTP(w, req)

	body := w.Body.String()
	expected := "Pair values: (-7, 42)"
	if body != expected {
		t.Errorf("ServeHTTP body = %q, want %q", body, expected)
	}
}

func TestInc(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{"zero", 0, 1},
		{"positive", 5, 6},
		{"negative", -3, -2},
		{"large", 999999, 1000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := make(chan int, 1)
			inc(tt.in, c)
			got := <-c
			if got != tt.want {
				t.Errorf("inc(%d) sent %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
