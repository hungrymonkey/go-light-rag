package handler

import (
	"bufio"
	"fmt"
	"strings"

	golightrag "github.com/MegaGrindStone/go-light-rag"
	"github.com/MegaGrindStone/go-light-rag/internal"
)

// Jsonl implements specialized document handling for Jsonl files.
// It extends the Default handler with Jsonl-specific functionality for parsing
// and processing Jsonl source files during RAG operations.
type Jsonl struct {
	Default
}

func (j Jsonl) ChunksDocument(content string) ([]golightrag.Source, error) {
	if content == "" {
		return []golightrag.Source{}, nil
	}

	var chunks []golightrag.Source
	var orderIndex int = 0
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		var token = scanner.Text()
		tokenSize, err := internal.CountTokens(token)
		if err != nil {
			return nil, fmt.Errorf("failed to count tokens on declaration: %w", err)
		}
		chunks = append(chunks, golightrag.Source{
			Content:    token,
			TokenSize:  tokenSize,
			OrderIndex: orderIndex,
		})
		orderIndex++
	}
	return chunks, nil
}

// EntityExtractionPromptData returns the data needed to generate prompts for extracting
// entities and relationships from Go source code content.
// It provides Go-specific entity extraction configurations, including custom goals,
// entity types, and examples tailored for Go language parsing.
func (j Jsonl) EntityExtractionPromptData() golightrag.EntityExtractionPromptData {
	language := j.Language
	if language == "" {
		language = defaultLanguage
	}
	return golightrag.EntityExtractionPromptData{
		Goal:        defaultEntityExtractionGoal,
		EntityTypes: defaultEntityTypes,
		Language:    language,
		Examples:    defaultEntityExtractionExamples,
	}
}

const jsonlKeywordExtractionGoal string = `Given a text document that is potentially relevant to this activity and a list of entity types, identify all entities of those types from the text and all relationships among the identified entities.`

// KeywordExtractionPromptData returns the data needed to generate prompts for extracting
// keywords from Go source code and related queries.
// It provides Go-specific keyword extraction configurations with custom goals
// and examples optimized for Go language patterns.
func (j Jsonl) KeywordExtractionPromptData() golightrag.KeywordExtractionPromptData {
	return golightrag.KeywordExtractionPromptData{
		Goal:     jsonlKeywordExtractionGoal,
		Examples: defaultKeywordExtractionExamples,
	}
}
