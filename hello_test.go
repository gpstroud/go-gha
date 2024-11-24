package main

import "testing"

func TestHelloWorld(t *testing.T) {
    expected := "Hello, World!"
    actual := myHelloWord()
    if actual != expected {
        t.Errorf("Expected %s but got %s", expected, actual)
    }
}
