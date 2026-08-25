// Package fixture is a minimal Go module used by go-generate-check's own
// CI to exercise the pass and fail paths of the action against a real
// //go:generate directive.
package fixture

//go:generate go run ./gen

// Greeting is a hand-written constant, unrelated to code generation.
const Greeting = "hello"
