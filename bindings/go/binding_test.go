package tree_sitter_lsl_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_lsl "github.com/tokeli/tree-sitter-lsl2/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_lsl.Language())
	if language == nil {
		t.Errorf("Error loading LSL grammar")
	}
}
