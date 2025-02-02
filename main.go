package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	utils "github.com/MdSadiqMd/search-service/utils"
)

func main() {
	log.SetFlags(log.Ltime)
	dumpPath := "enwiki-latest-abstract1.xml"

	start := time.Now()
	log.Printf("🚀 Starting search service...")

	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatalf("❌ Failed to load documents: %v", err)
	}
	utils.LogWithEmoji("📚", fmt.Sprintf("Loaded %d documents", len(docs)), start)

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	utils.LogWithEmoji("📇", fmt.Sprintf("Indexed %d documents", len(docs)), start)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n🔍 Enter search query (or 'exit' to quit): ")
		if !scanner.Scan() {
			break
		}

		query := strings.TrimSpace(scanner.Text())
		if query == "exit" {
			fmt.Println("👋 Goodbye!")
			break
		}
		if query == "" {
			continue
		}

		start = time.Now()
		matches := idx.Search(query)
		utils.LogWithEmoji("🔍", fmt.Sprintf("Found %d matches", len(matches)), start)
		if len(matches) == 0 {
			log.Printf("😕 No results found for %q", query)
			continue
		}
		log.Printf("📋 Top results for %q:", query)
		for i, id := range matches {
			if i >= 10 {
				log.Printf("   ... and %d more matches", len(matches)-10)
				break
			}
			log.Printf("   %d. %s\n      URL: %s", i+1, docs[id].Title, docs[id].URL)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("❌ Error reading input: %v", err)
	}
}
