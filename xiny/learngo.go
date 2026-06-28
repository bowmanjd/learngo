/* Go was created out of the need to get work done. It's not the latest trend
in programming language theory, but it is a highly effective way to solve
real-world problems. It draws concepts from imperative languages with static
typing. It's incredibly fast to compile and execute, adds clean concurrency
to leverage multi-core systems, and scales effortlessly to massive codebases.

Go comes with a robust standard library and a highly active community.
*/

// Single line comment
/* Multi-line
   comment */

/* A build tag is a compiler directive starting with //go:build. Build tags must
be placed before the package clause, followed by a blank line. Example:

    //go:build prod || dev || test

Build with: go build -tags="prod" to include files with that tag.
*/

// A package clause starts every source file.
// "main" is a special name declaring an executable program rather than a library.
package main

// Import declaration declares standard library or third-party packages referenced in this file.
import (
	"cmp"      // Ordered comparison types (Go 1.21+)
	"context"  // Cancellation and timeout propagation
	"errors"   // Error wrapping and inspection
	"fmt"      // Formatting and I/O standard package
	"io"       // Core I/O interfaces (replaces deprecated io/ioutil)
	"iter"     // Custom range-over-func iterators (Go 1.23+)
	"log/slog" // Structured logging (Go 1.21+)
	m "math"   // Math library with a local alias 'm'
	"net/http" // Production-ready web server
	"os"       // Operating system interface
	"regexp"   // Regular expressions for input validation
	"slices"   // Type-safe slice utility functions (Go 1.21+)
	"strconv"  // String conversions
	"time"     // Time and duration utilities
)

// A function definition. "main" is the entry point of the executable.
func main() {
	// Println outputs a line to stdout.
	fmt.Println("Hello world!")

	// Call another function within this package.
	beyondHello()
}

// Functions declare parameters in parentheses.
func beyondHello() {
	var x int // Variable declaration. Variables must be declared before use.
	x = 3     // Variable assignment.

	// "Short" declarations use := to infer type, declare, and assign.
	y := 4
	sum, prod := learnMultiple(x, y) // Functions can return multiple values.
	fmt.Println("sum:", sum, "prod:", prod)
	fmt.Println("named return:", learnNamedReturns(x, y))

	learnTypes()
}

/*
Go functions can have multiple parameters and multiple return values.

	In this case, x and y are input parameters, sum and prod are named return values.
*/
func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // Return statement returns multiple values.
}

// Named return values allow a bare "return" statement.
// Use sparingly — it can reduce clarity in longer functions.
func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return // z is implicit here because we named it earlier.
}

// Built-in types and basic constructs.
func learnTypes() {
	// Short declarations are the standard convention.
	str := "Learn Go!" // string type.
	s2 := `A "raw" string
literal can include line breaks.` // Same string type.

	// Non-ASCII literals are fully supported. Go source is UTF-8.
	g := 'Σ' // rune type, which is an alias for int32 (holds a Unicode code point).

	f := 3.14159 // float64 (IEEE-754 64-bit float).
	c := 3 + 4i  // complex128 (represented internally by two float64s).

	// var syntax with explicit initialization.
	var u uint = 7 // Unsigned integer.
	var pi float32 = 22. / 7

	// Conversion syntax.
	n := byte('\n') // byte is an alias for uint8.

	// Arrays have a fixed size defined at compile-time.
	var a4 [4]int           // An array of 4 ints, initialized to zero value (0).
	a3 := [...]int{3, 1, 5} // An array initialized with a fixed size of 3 elements.

	// Arrays have value semantics — assignment copies the entire array.
	a4Copy := a4
	a4Copy[0] = 25
	fmt.Println("arrays copied by value:", a4Copy[0] == a4[0]) // false

	// Slices have dynamic size. Arrays and slices have distinct use-cases,
	// but slices are much more common.
	s3 := []int{4, 5, 9}    // A slice literal has no size inside the brackets.
	s4 := make([]int, 4)    // Allocates a slice of 4 ints initialized to 0.
	var d2 [][]float64      // Declaration only, d2 is currently nil.
	bs := []byte("a slice") // Type conversion to byte slice.

	// Slices have reference semantics — they share underlying storage.
	s3Copy := s3
	s3Copy[0] = 0
	fmt.Println("slices share storage:", s3Copy[0] == s3[0]) // true
	s3[0] = 4                                                // Restore original value for later use.

	// Because slices are dynamic, elements can be appended:
	s3 = append(s3, 12, 13) // Append elements and reassign to s3.
	fmt.Println(s3)

	// Append another slice using the spread operator (...).
	s3 = append(s3, []int{20, 21}...)
	fmt.Println("after spread append:", s3)

	// Slices can also be manipulated using the standard 'slices' package (Go 1.21+).
	if slices.Contains(s3, 9) {
		fmt.Println("s3 contains 9!")
	}

	// Under Go's compilation rules, unused variables are compile errors.
	// Use the blank identifier '_' to discard unused variables.
	_, _, _, _, _, _, _, _, _ = str, s2, g, f, c, u, pi, n, a3

	// To satisfy compiling constraints, print our arrays and slices:
	fmt.Println(a4, s4, d2, bs)

	learnFlowControl()
}

func expensiveComputation() int {
	return int(m.Exp(10))
}

// Go features automatic garbage collection. It has pointers, but no pointer arithmetic.
// Unlike C, returning the address of a local variable is safe — Go manages the lifetime.
func learnMemory() (*int, *int) {
	// new(Type) allocates zero-initialized memory and returns a pointer.
	p := new(int)
	*p = 7 // Dereference to assign.

	// Go 1.26+ supports expression-based new(). Pass a value directly
	// to allocate and initialize a pointer in a single expression.
	q := new(-2)

	// make() allocates slices, maps, and channels (not pointers).
	s := make([]int, 20)
	s[0] = 99

	// '&' takes the address of any addressable value.
	fmt.Println("slice element address:", &s[0])

	return p, q
}

func learnFlowControl() {
	// If-statements require braces, but do not require parentheses.
	if true {
		fmt.Println("Always prints")
	}

	// Code formatting is strictly standardized via the "go fmt" tool.
	// You can execute a variable assignment inside the "if" conditional block.
	if y := expensiveComputation(); y > 100 {
		fmt.Println("Computation was large:", y)
	}

	// Switch statements do not need manual "break" statements; they break automatically.
	// Use the "fallthrough" keyword if you want code to cascade.
	switch "Go" {
	case "Rust":
		fmt.Println("Not Go")
	case "Go":
		fmt.Println("Found Go!")
	default:
		fmt.Println("Fallback case")
	}

	// 'for' is the only loop statement in Go, but it has several forms.
	// 1. Classical C-style loop:
	for i := 0; i < 3; i++ {
		fmt.Println("Iteration:", i)
	}

	// 2. Loop over integer ranges (Go 1.22+ syntax):
	for i := range 3 {
		fmt.Println("Modern range loop index:", i) // Prints 0, 1, 2
	}

	// 3. While-loop style (condition only):
	counter := 0
	for counter < 3 {
		fmt.Println("While-loop equivalent count:", counter)
		counter++
	}

	// 4. Infinite loop:
	for {
		break // Exits loop
	}

	// 5. Range loop over collections:
	items := []string{"apple", "banana"}
	for index, value := range items {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// 6. Range-over-func Custom Iterators (Go 1.23+):
	// Custom yield-based iterators can be cleanly evaluated directly in 'for range'.
	for num := range countdown(3) {
		fmt.Println("Countdown:", num)
	}

	// Goroutine variable capture inside loops (Go 1.22+ Scoping):
	// Historically, referencing a loop variable directly inside a goroutine was a bug.
	// As of Go 1.22, loop variables are allocated per-iteration, making this completely safe:
	for i := range 3 {
		go func() {
			// This is safe! Each iteration has its own independent "i" instance.
			fmt.Println("Async index:", i)
		}()
	}

	// Function literals are closures.
	x := 10
	closure := func() int {
		return x * 2 // Captures variable x from outer scope.
	}
	fmt.Println("Closure result:", closure())

	// Function literals can be defined and called inline.
	fmt.Println("Add + double two numbers:",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2)) // Called immediately with args 10 and 2.

	// Demonstrate learning memory features
	p, q := learnMemory()
	fmt.Println("Pointers values:", *p, *q)

	learnFunctionFactory()
}

func learnFunctionFactory() {
	// Functions can return functions (higher-order functions).
	fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

	d := sentenceFactory("summer")
	fmt.Println(d("A beautiful", "day!"))
	fmt.Println(d("A lazy", "afternoon!"))

	learnDefer()
}

// sentenceFactory returns a closure that captures mystring.
func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after)
	}
}

func learnDefer() {
	// defer pushes a function call onto a stack executed AFTER the function returns.
	// Deferred calls run in LIFO (reverse) order.
	defer fmt.Println("deferred: runs third (LIFO order)")
	defer fmt.Println("deferred: runs second")
	defer fmt.Println("deferred: runs first")

	// Common use: closing resources near where they're opened.
	fmt.Println("normal execution: runs before all defers")

	learnInterfaces()
}

// Custom Range-over-func Iterator (Go 1.23+).
// Returns an iter.Seq[int] which can be evaluated directly in a 'for range' loop.
func countdown(start int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i > 0; i-- {
			if !yield(i) {
				return // Consumer requested early termination.
			}
		}
	}
}

// Generics (Go 1.18+) allow writing code with parametric types.
// T is a type parameter restricted by cmp.Ordered (any type that supports <, >, ==).
func learnGenerics[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Struct declarations define structured, user-defined types.
type pair struct {
	x, y int
}

// Define a method on the custom struct type.
// This is a value receiver method.
func (p pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

// pair can also implement http.Handler, making it a web server endpoint.
// This demonstrates how one type can satisfy multiple interfaces.
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Pair values: %s", p.String())))
}

func learnInterfaces() {
	// Initializing a struct literal using brackets.
	p := pair{3, 4}
	fmt.Println(p.String())

	// pair implements fmt.Stringer, so it can be assigned to that interface.
	// Interfaces are satisfied implicitly — no "implements" keyword needed.
	var s fmt.Stringer = p
	fmt.Println("Interface output:", s.String())

	// fmt.Println automatically calls String() on types that implement Stringer.
	fmt.Println("fmt sees Stringer:", p)

	// Type switch allows switching on the type of an interface value.
	var data any = "gopher" // 'any' is the modern spelling of interface{}.
	switch v := data.(type) {
	case string:
		fmt.Println(v, "is a string")
	case int64:
		fmt.Printf("%d is an int64\n", v)
	default:
		fmt.Println("unknown type")
	}

	// Testing generics function
	fmt.Println("Generics Max Int:", learnGenerics(10, 20))
	fmt.Println("Generics Max String:", learnGenerics("apple", "banana"))

	learnVariadicParams("great", "learning", "here!")
}

// Functions can have variadic parameters (variable number of arguments).
func learnVariadicParams(myStrings ...any) {
	for _, param := range myStrings {
		fmt.Println("variadic param:", param)
	}

	// Pass variadic values to another variadic function using ...
	fmt.Println("spread to Sprintln:", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	// Maps return a second "ok" boolean for key existence checks.
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok {
		fmt.Println("no one there")
	} else {
		fmt.Println(x)
	}

	// Most functions that can fail return an error as the last return value.
	if _, err := strconv.Atoi("non-int"); err != nil {
		fmt.Println("Atoi error:", err)
	}

	// Use errors.Is to check for specific error types (Go 1.13+).
	if _, err := os.ReadFile("definitely-missing.txt"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("missing file detected with errors.Is")
		}
	}

	learnConcurrency()
}

// inc demonstrates a send-only channel parameter.
func inc(i int, c chan<- int) {
	c <- i + 1
}

func learnConcurrency() {
	// Channels coordinate concurrent execution.
	// make() allocates and initializes slices, maps, and channels.
	// The second argument is the buffer size (0 = unbuffered).
	c := make(chan int, 3)

	// Launch concurrent goroutines using the 'go' statement.
	// As of Go 1.22, loop variables are allocated per-iteration, making this safe.
	for _, n := range []int{0, 10, -805} {
		go inc(n, c)
	}

	// Receive values from the channel (blocks until senders are ready).
	fmt.Println("concurrent results:", <-c, <-c, <-c)

	// Select waits on multiple channel operations.
	// It picks randomly from cases that are ready to communicate.
	cs := make(chan string)
	ccs := make(chan chan string) // Channels can hold any type, including other channels.
	go func() { c <- 84 }()
	go func() { cs <- "wordy" }()

	select {
	case i := <-c:
		fmt.Printf("received int: %d\n", i)
	case s := <-cs:
		fmt.Printf("received string: %s\n", s)
	case <-ccs:
		fmt.Println("didn't happen — channel is empty")
	}

	learnWebProgramming()
}

func learnWebProgramming() {
	logger := slog.Default()

	// Modern HTTP Routing (Go 1.22+ wildcard routing).
	mux := http.NewServeMux()

	// Method and path pattern with parameter extraction.
	var validName = regexp.MustCompile(`^[\p{L}\p{N} .'-]{1,64}$`)

	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		if !validName.MatchString(name) {
			http.Error(w, "invalid name", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = io.WriteString(w, fmt.Sprintf("Hello, %s!\n", name))
	})

	// Any type implementing http.Handler can be used directly.
	// pair implements ServeHTTP, so it satisfies http.Handler.
	mux.Handle("GET /pair", pair{3, 4})

	// Configure server with timeouts (security best practice).
	srv := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	errCh := make(chan error, 1)

	go func() {
		logger.Info("starting demo server", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	// Make a request to demonstrate the server works.
	requestServer()

	// Graceful shutdown with context timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("shutdown error:", err)
	}

	if err, ok := <-errCh; ok && err != nil {
		logger.Error("server error", "err", err)
	}
}

func requestServer() {
	client := &http.Client{Timeout: 500 * time.Millisecond}

	// Retry loop waiting for server to start.
	for range 20 {
		resp, err := client.Get("http://127.0.0.1:8080/hello/Gopher")
		if err != nil {
			time.Sleep(25 * time.Millisecond)
			continue
		}

		// Read and close body explicitly — don't use defer inside a loop.
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println("read failed:", err)
			return
		}

		fmt.Printf("web server said: %q\n", string(body))
		return
	}

	fmt.Println("server did not start in time")
}
