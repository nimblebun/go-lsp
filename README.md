# Language Server Protocol types for Go

This is a module that contains type definitions and documentation for the
[Language Server Protocol Specification][lsp-spec]. It aims to provide
up-to-date and clear definitions to use in a language server.

**LSP spec support:** 3.15, 3.16.

## Getting Started

```
go get pkg.nimblebun.works/go-lsp
```

Import it in your project:

```go
import (
  "pkg.nimblebun.works/go-lsp"
)
```

You can now use the types and constants defined in `lsp`. See the
[documentation][docs-link] for more information.

## Disclaimer

Our goal was to create an organized, easy to use, well-documented, and
up-to-date LSP module for the Go programming language. To achieve this goal, we
decided not to rely on automated tools to generate the Go types, but to write
everything ourselves by hand. As a result, even though we tried to be as precise
as possible, there could be typos in the definitions or the JSON names. If you
notice a typo or bug, please report it in the [Issue Tracker][issues-link]!

## Future Plans

Future plans for this module include designing a language server SDK,
effectively making the module a swiss-army knife for creating language servers
in Go easily.

## License

MIT. Type declarations and documentations are taken from the
[official specification document][lsp-spec].

[docs-link]: https://pkg.go.dev/pkg.nimblebun.works/go-lsp
[issues-link]: https://github.com/nimblebun/go-lsp/issues
[lsp-spec]: https://microsoft.github.io/language-server-protocol/specifications/specification-current
