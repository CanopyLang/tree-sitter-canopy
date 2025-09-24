package tree_sitter_canopy_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_canopy "github.com/CanopyLang/tree-sitter-canopy/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_canopy.Language())
	if language == nil {
		t.Errorf("Error loading Canopy grammar")
	}
}
