package app

import (
	"fmt"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp()

	fmt.Printf("app=>[%+v]", app)

	if app == nil {
		t.Error("create app failed")
	}
}

func TestNewRoute(t *testing.T) {
	app := NewApp()

	fmt.Printf("app=>[%+v]", app)

	if app == nil {
		t.Error("create app failed")
	}

	route := NewRoute(app)

	if route == nil {
		t.Error("create route failed")
	}
}
