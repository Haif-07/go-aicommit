package main

import (
	"encoding/json"
	"fmt"
)

type CommitType string

const (
	CommitTypeEmpty        CommitType = ""
	CommitTypeConventional CommitType = "conventional"
)

var commitTypeFormats = map[CommitType]string{
	CommitTypeEmpty:        "<commit message>",
	CommitTypeConventional: "<type>(<optional scope>): <commit message>",
}

func specifyCommitFormat(commitType CommitType) string {
	return fmt.Sprintf("The output response must be in format:\n%s", commitTypeFormats[commitType])
}

var commitTypes = map[CommitType]string{
	CommitTypeEmpty: "",
	CommitTypeConventional: fmt.Sprintf("Choose a type from the type-to-description JSON below that best describes the git diff:\n%s",
		serializeCommitDescriptions()),
}

func serializeCommitDescriptions() string {
	commitDescriptions := map[string]string{
		"docs":     "Documentation only changes",
		"style":    "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)",
		"refactor": "A code change that neither fixes a bug nor adds a feature",
		"perf":     "A code change that improves performance",
		"test":     "Adding missing tests or correcting existing tests",
		"build":    "Changes that affect the build system or external dependencies",
		"ci":       "Changes to our CI configuration files and scripts",
		"chore":    "Other changes that don't modify src or test files",
		"revert":   "Reverts a previous commit",
		"feat":     "A new feature",
		"fix":      "A bug fix",
	}

	// Convert to JSON string with indent
	data, err := json.MarshalIndent(commitDescriptions, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error serializing commit descriptions: %v", err)
	}
	return string(data)
}

func generatePrompt(locale string, maxLength int, commitType CommitType) string {
	prompt := []string{
		"Generate a concise git commit message written in present tense for the following code diff with the given specifications below:",
		fmt.Sprintf("Message language: %s", locale),
		fmt.Sprintf("Commit message must be a maximum of %d characters.", maxLength),
		"Exclude anything unnecessary such as translation. Your entire response will be passed directly into git commit.",
		commitTypes[commitType],
		specifyCommitFormat(commitType),
		",If multiple files are modified, each file uses this response format",
	}

	// Filter out empty elements from the slice
	var filteredPrompt []string
	for _, line := range prompt {
		if line != "" {
			filteredPrompt = append(filteredPrompt, line)
		}
	}

	return joinLines(filteredPrompt)
}

func joinLines(lines []string) string {
	return fmt.Sprintf("%s", lines)
}
