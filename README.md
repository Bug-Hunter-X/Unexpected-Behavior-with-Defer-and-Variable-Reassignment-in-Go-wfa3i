# Go Defer Statement and Variable Reassignment

This example demonstrates the potential for unexpected behavior when using the `defer` keyword in Go along with variable reassignments within the function. The deferred print statement will use the final values of x and y, not the initial values.

## Bug

The `bug.go` file shows how `defer`'s execution is tied to the function's return, not the moment of declaration.  This leads to the incorrect sum being printed first.

## Solution

The `bugSolution.go` file illustrates a way to achieve the intended behavior, capturing the initial values using an immediately-invoked function expression (IIFE) to preserve the original state of the variables.
