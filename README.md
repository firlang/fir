<div align="center">
    <h1>The Fir Programming language</h1>
    <a href="https://firlang.vercel.app">firlang.vercel.app</a><br/>
    <br/>
    <a href="https://github.com/firlang/fir/blob/main/LICENSE"><img src="https://img.shields.io/github/license/firlang/fir"></a>
    &nbsp;<a href="https://github.com/firlang/fir/actions"><img src="https://img.shields.io/github/actions/workflow/status/firlang/fir/go.yml"></a>
    &nbsp;<a href="https://github.com/firlang/fir/blob/main/go.mod"><img src="https://img.shields.io/github/go-mod/go-version/firlang/fir"></a>
    &nbsp;<img src="https://img.shields.io/github/languages/code-size/firlang/fir">
</div>

> ### Fir is a programming language

(well, not yet. Fir is under heavy development and is right now just a lexer! Hopefully, it'll be in a usable state soon enough though!)

* [Todo](#todo)
* [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Building from source](#building-from-source)

## Todo
- [ ] Lexing
    - [x] Basic operators (+, -, *, /)
    - [x] Assignment (:=)
    - [ ] Numbers
        - [x] Integers
        - [x] Floating-points
    - [ ] Strings
    - [x] Idents
    - [ ] Keywords
- [ ] Parsing
- [ ] Eval

## Getting Started

### Prerequisites
* Go >1.20
* Make


### Building from source
1. Clone the latest version of the repo from Github
    - gh: `gh repo clone firlang/fir`

    - git: `git clone https://github.com/firlang/fir`
2. Build the project by running `make` (the executable will be in `debug`)