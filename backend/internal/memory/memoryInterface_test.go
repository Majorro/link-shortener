package memory_test

import (
	"testing"
	"time"

	"snaLinkShortener/internal/memory"
	"snaLinkShortener/internal/memory/inMemory"
)

func TestGenerator(t *testing.T) {
	entry1 := memory.MemoryEntry{
		LongLink:  "https://google.com",
		ShortId:   "",
		Author:    "test",
		CreatedAt: time.Now(),
	}

	{
		var db = memory.Memory{
			Storage: inMemory.GetMemoryInstance(),
		}
		entry1.GenerateUniqueShortLink(db)

	}
}
