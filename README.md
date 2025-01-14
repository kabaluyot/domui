# Domui
A DOM UI framework for Go

## Features

* Pure go code compiled to wasm
  + Utilizing Go's excellent support for static typing and dynamic reflection
  + Runs on any WebAssembly and DOM supported browsers, Electron, etc
* Dependent reactive system for both view and state management
  + Conceptually unify view component and state management
  + No hooks, redux, recoil, mobx ...
* Auto and efficient DOM tree updating
  + Based on virtual DOM diff and patch

## Prerequisites

* Go compiler 1.16 or newer
* If you are not familiar with compiling and running WebAssembly program, 
please read [the official wiki](https://github.com/golang/go/wiki/WebAssembly)

## Table of contents

* [Tutorial](#tutorial)
  + [Minimal runnable program](#minimal)
  + [The dependent system](#dependent)
  + [The reactive system](#reactive)
  + [DOM element specifications](#dom)
  + [Parameterized element](#parameterized)
  + [Miscellaneous usages](#misc)
* [Additional topics](#additional)
  + [Differences between DomUI and reactjs](#reactjs)

<a name="tutorial" />

## Tutorial

<a name="minimal" />

### Minimal runnable program

```go
package main

import (
	"github.com/reusee/domui"
	"syscall/js"
	"time"
)

// ergonomic aliases
var (
	Div = domui.Tag("div")
	T   = domui.Text
)

// A type to hold all definitions
type Def struct{}

// define the root UI element
func (_ Def) RootElement() domui.RootElement {
	// same to <div>Hello, world!</div>
	return Div(T("Hello, world!"))
}

func main() {
	domui.NewApp(
		// render on <div id="app">
		js.Global().Get("document").Call("getElementById", "app"),
		// pass Def's methods as definitions
		domui.Methods(new(Def))...,
	)
	// prevent from exiting
	time.Sleep(time.Hour * 24 * 365 * 100)
}
```

<a name="dependent" />

### The dependent system

The definition of RootElement can be refactored to multiple dependent components.

```go
package main

import (
	"github.com/reusee/domui"
	"syscall/js"
	"time"
)

var (
	Div = domui.Tag("div")
	T   = domui.Text
)

type (
	Def  struct{}
	Spec = domui.Spec
)

// A string-typed state
type Greetings string

// define Greetings
func (_ Def) Greetings() Greetings {
	return "Hello, world!"
}

// A UI element
type GreetingsElement Spec

// define GreetingsElement
func (_ Def) GreetingsElement(
	// use Greetings
	greetings Greetings,
) GreetingsElement {
	return Div(T("%s", greetings))
}

// The root UI element
func (_ Def) RootElement(
	// use GreetingsElement
	greetingsElem GreetingsElement,
) domui.RootElement {
	return Div(
		greetingsElem,
	)
}

func main() {
	domui.NewApp(
		js.Global().Get("document").Call("getElementById", "app"),
		domui.Methods(new(Def))...,
	)
	time.Sleep(time.Hour * 24 * 365 * 100)
}
```

<a name="reactive" />

### The reactive system

Definitions can be updated. 
All affected definitions will be re-calculated recursively till the RootElement.

```go
package main

import (
	"github.com/reusee/domui"
	"syscall/js"
	"time"
)

var (
	Div     = domui.Tag("div")
	T       = domui.Text
	OnClick = domui.On("click")
)

type (
	Def    struct{}
	Spec   = domui.Spec
	Update = domui.Update
)

type Greetings string

func (_ Def) Greetings() Greetings {
	return "Hello, world!"
}

type GreetingsElement Spec

func (_ Def) GreetingsElement(
	greetings Greetings,
) GreetingsElement {
	return Div(T("%s", greetings))
}

func (_ Def) RootElement(
	greetingsElem GreetingsElement,
	// use the Update function
	update Update,
) domui.RootElement {
	return Div(
		greetingsElem,

		// when clicked, do update
		OnClick(func() {
			// provide a new definition for Greetings
			update(func() Greetings {
				return "Hello, DomUI!"
			})
		}),
	)
}

func main() {
	domui.NewApp(
		js.Global().Get("document").Call("getElementById", "app"),
		domui.Methods(new(Def))...,
	)
	time.Sleep(time.Hour * 24 * 365 * 100)
}
```

<a name="dom" />

### DOM element specifications

The above programs demonstrated tag and event usages.
Attributes, styles, classes can also be specified.

```go
package main

import (
	"github.com/reusee/domui"
	"syscall/js"
	"time"
)

var (
	Div        = domui.Tag("div")
	Link       = domui.Tag("a")
	Ahref      = domui.Attr("href")
	Sfont_size = domui.Style("font-size")
	T          = domui.Text
	ID         = domui.ID
	Class      = domui.Class
)

type Def struct{}

func (_ Def) RootElement() domui.RootElement {
	return Div(
		Link(
			T("Hello, world!"),
			// id
			ID("link"),
			// class
			Class("link1", "link2"),
			// href attribute
			Ahref("http://github.com"),
			// font-size style
			Sfont_size("1.6rem"),
		),
	)
}

func main() {
	domui.NewApp(
		js.Global().Get("document").Call("getElementById", "app"),
		domui.Methods(new(Def))...,
	)
	time.Sleep(time.Hour * 24 * 365 * 100)
}
```

<a name="parameterized" />

### Parameterized element

To make a reusable element, define it as a function.

```go
package main

import (
	"fmt"
	"github.com/reusee/domui"
	"reflect"
	"strings"
	"syscall/js"
	"time"
)

var (
	Div     = domui.Tag("div")
	T       = domui.Text
	OnClick = domui.On("click")
)

type (
	Def    struct{}
	any    = interface{}
	Spec   = domui.Spec
	Update = domui.Update
	Specs  = domui.Specs
)

// Greetings with name parameter
type Greetings func(name any) Spec

func (_ Def) Greetings(
	update Update,
) Greetings {
	return func(name any) Spec {
		return Specs{
			T("Hello, %s!", name),
			OnClick(func() {
				// when clicked, update the name argument to upper case
				// use reflect to support all string-typed arguments
				nameValue := reflect.New(reflect.TypeOf(name))
				nameValue.Elem().SetString(
					strings.ToUpper(fmt.Sprintf("%s", name)))
				update(nameValue.Interface())
			}),
		}
	}
}

// string-typed states
type (
	TheWorld string
	TheDomUI string
)

// define two types in one function
func (_ Def) Names() (TheWorld, TheDomUI) {
	return "world", "DomUI"
}

func (_ Def) RootElement(
	greetings Greetings,
	world TheWorld,
	domUI TheDomUI,
) domui.RootElement {
	return Div(
		// use Greetings
		Div(
			greetings(world),
		),
		Div(
			greetings(domUI),
		),
	)
}

func main() {
	domui.NewApp(
		js.Global().Get("document").Call("getElementById", "app"),
		domui.Methods(new(Def))...,
	)
	time.Sleep(time.Hour * 24 * 365 * 100)
}
```

<a name="misc" />

### Miscellaneous usages

When updating a definition that has no dependency,
instead of passing a function, a pointer can be used.

```go
newGreetings := Greetings("Hello!")
update(&newGreetings)
// is the same to
update(func() Greetings {
  return "Hello!"
})
```

To do something on App inits, define one or more OnAppInit

```go
func (_ Def) InitFoo() domui.OnAppInit {
  // do something
}

func (_ Def) InitBar() domui.OnAppInit {
  // do something
}
```

To access the DOM element in an event handler, use a js.Value parameter 

```go
type Foo Spec

func (_ Def) Foo() Foo {
  retrun Div(
    OnClick(func(elem js.Value) {
      _ = elem.Call("getAttribute", "href").String()
    }),
  )
}
```

Specs can be cached for reusing

```go
type Article func(title string, content string) Spec

func (_ Def) Article() Article {
  m := domui.NewSpecMap()
  return m(
    // key
    [2]any{title, content},
    // value
    func() Spec {
      return Div( /* ... */ )
    },
  )
}
```

Some conditional Spec constructors are provided

```go
var (
  A = domui.Tag("a")
  Ahref = domui.Attr("href")
  Sfont_weight = domui.Style("font-weight")
)

type Link func(href string, bold bool) Spec

func (_ Def) Link() Link {
  return func(href string, bold bool) Spec {
    return A(
      // If
      domui.If(href != "", Ahref(href)),
      // Alt
      domui.Alt(bold,
        Sfont_weight("bold"),
        Sfont_weight("normal"),
      ),
    )
  }
}
```

And loop constructors

```go
type List func(elems []string) Spec

func (_ Def) List() List {
  return func(elems []string) Spec {
    return Div(
      // For
      domui.For(elems, func(s string) Spec {
        return Div(T("%s", s))
      }),
      // Range
      domui.Range(elems, func(i int, s string) Spec {
        return Div(T("%d: %s", i, s))
      }),
    )
  }
}
```

<a name="additional" />

## Additional topics

<a name="reactjs" />

### Differences between DomUI and reactjs

In react, UI components are declared as classes or functions, and there may be multiple instances of the same class.
All component instances form the component tree.
In addition to the component tree, a state tree or other state management scheme may be used to hold non-UI data.

In DomUI, all UI and non-UI components are defined as type/value pairs.
UI components are just Spec types or function types that returning Spec.
This unification greatly reduces the complexity of nearly everything compared to other MVVM frameworks.

