package handler_test

import (
	"fmt"
	"strings"
	"testing"

	golightrag "github.com/MegaGrindStone/go-light-rag"
	"github.com/MegaGrindStone/go-light-rag/handler"
	"github.com/MegaGrindStone/go-light-rag/internal"
)

//nolint:gocognit,cyclop,gocyclo // Test cases are complex
func TestJsonl_ChunksDocument(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
		verify  func(t *testing.T, chunks []golightrag.Source)
	}{
		{
			name: "Package and imports",
			content: `{"id": "pubmed23n0975_26276", "title": "OsDCL3b"}
{"id": "pubmed23n0975_262", "title": "OsDCL"}`,
			verify: func(t *testing.T, chunks []golightrag.Source) {
				if len(chunks) != 2 {
					t.Fatalf("Expected 1 chunk, got %d", len(chunks))
				}

				expectedTokens, _ := internal.CountTokens(chunks[0].Content)
				if chunks[0].TokenSize != expectedTokens {
					t.Errorf("TokenSize mismatch: got %d, want %d", chunks[0].TokenSize, expectedTokens)
				}
				fmt.Println(chunks[0].Content)
				if !strings.Contains(chunks[0].Content, `{"id": "pubmed23n0975_26276", "title": "OsDCL3b"}`) {
					t.Errorf("Chunk should contain " + `{"id": "pubmed23n0975_26276", "title": "OsDCL3b"}`)
				}
			},
		},
	}
	runJsonlChunksDocumentTests(t, tests)
}

func runJsonlChunksDocumentTests(t *testing.T, tests []struct {
	name    string
	content string
	wantErr bool
	verify  func(t *testing.T, chunks []golightrag.Source)
},
) {
	jsonlHandler := handler.Jsonl{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonlHandler.ChunksDocument(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jsonl.ChunksDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if tt.verify != nil {
				tt.verify(t, got)
			}
		})
	}
}
