// file_handler.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func saveToFiles(adventure *Adventure, outputDir string) error {
	for i, episode := range adventure.Episodes {
		// Create episode directory
		episodeDir := filepath.Join(outputDir, fmt.Sprintf("Episode-%02d", i+1))
		if err := os.MkdirAll(episodeDir, 0755); err != nil {
			return fmt.Errorf("creating episode directory: %w", err)
		}

		// Save episode content
		episodePath := filepath.Join(episodeDir, "Episode.md")
		if err := ioutil.WriteFile(episodePath, []byte(episode.FullAdventure), 0644); err != nil {
			return fmt.Errorf("saving episode: %w", err)
		}

		// Save illustration prompts
		for j, illus := range episode.Illustrations {
			illusPath := filepath.Join(episodeDir, fmt.Sprintf("Illustration-%02d.md", j+1))
			content := fmt.Sprintf("Description: %s\nStyle: %s\nIs Map: %v",
				illus.Description, illus.Style, illus.IsMap)
			if err := ioutil.WriteFile(illusPath, []byte(content), 0644); err != nil {
				return fmt.Errorf("saving illustration prompt: %w", err)
			}
		}
	}
	return nil
}
