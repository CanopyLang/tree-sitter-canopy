[![Build Status](https://github.com/CanopyLang/tree-sitter-canopy/actions/workflows/test.yml/badge.svg)](https://github.com/CanopyLang/tree-sitter-canopy/actions/workflows/test.yml)

# Tree-sitter for Canopy

A tree-sitter parser for the [Canopy programming language](https://github.com/quinten/canopy) - a modern functional language based on Elm with enhanced tooling and compilation features.

## About Canopy

Canopy is a delightful functional programming language that compiles to JavaScript. It's based on Elm but includes several enhancements:

- **Smaller JavaScript bundles** - Optimized compilation produces more compact output
- **Enhanced FFI system** - Type-safe Foreign Function Interface with capability-based Web API access
- **Improved tooling support** - Better language server integration and development experience
- **Full Elm compatibility** - Existing Elm code works seamlessly in Canopy

The main Canopy compiler can be found at [~/fh/canopy](~/fh/canopy) and provides complete tooling including `canopy make`, `canopy repl`, `canopy install`, and more.

## Why Tree-sitter for Canopy?

This tree-sitter grammar enables advanced tooling for Canopy development:

- **Fast, incremental parsing** - Parse on every keystroke for real-time feedback
- **Syntax highlighting** - Rich, context-aware highlighting in editors
- **Error resilience** - Continues parsing even with syntax errors
- **Language server foundation** - Building block for comprehensive IDE support

## Language Features Supported

The grammar supports all Canopy language constructs:

- **Module system** - Module declarations, imports, and exposing lists
- **Type system** - Type annotations, custom types, type aliases, and records
- **Pattern matching** - Comprehensive case expressions and pattern destructuring
- **Functions** - Function declarations, lambdas, and partial application
- **Data structures** - Lists, tuples, records, and record updates
- **Comments** - Both line comments (`--`) and block comments (`{- -}`)
- **FFI declarations** - Foreign import statements for JavaScript interop

## Testing

This parser is thoroughly tested against:

- All examples in the Canopy compiler repository
- Complex real-world Canopy applications
- Edge cases and error conditions
- Compatibility with existing Elm codebases

## Installation & Usage

### For Editor Integration

Most editors with tree-sitter support can use this grammar. Check your editor's documentation for adding custom tree-sitter grammars.

### For Development

```bash
# Clone the repository
git clone https://github.com/CanopyLang/tree-sitter-canopy
cd tree-sitter-canopy

# Generate the parser
tree-sitter generate

# Test the parser
tree-sitter test

# Parse a file
tree-sitter parse examples/basic.can
```

## Contributing

We welcome contributions! Here's how you can help:

- **Test cases** - Find Canopy (`.can`) files that fail to parse correctly
- **Grammar improvements** - Enhance parsing accuracy and performance
- **Documentation** - Improve examples and usage documentation
- **Editor integration** - Help integrate with more editors and tools

## File Extensions

Canopy uses the `.can` file extension (though `.canopy` is also supported for compatibility).

## Acknowledgments

Special thanks to:
- [@klazuka](https://github.com/klazuka) and the [intellij-elm](https://github.com/klazuka/intellij-elm/) team for the original Elm parser structure
- The Elm community for creating the foundation that Canopy builds upon
- The tree-sitter community for the excellent parsing framework

## License

This project is licensed under the same terms as the Canopy compiler.
