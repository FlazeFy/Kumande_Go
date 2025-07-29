package seeders

import (
	"kumande/configs"
	"kumande/factories"
	"kumande/modules/dictionary"
	"log"
)

func SeedDictionaries(repo dictionary.DictionaryRepository) {
	// Empty Table
	repo.DeleteAll()

	var seedData = []struct {
		DictionaryType  string
		DictionaryNames []string
	}{
		{"currency", configs.Currencies},
	}

	// Fill Table
	var success = 0
	for _, dt := range seedData {
		for _, dictionaryName := range dt.DictionaryNames {
			dct := factories.DictionaryFactory(dictionaryName, dt.DictionaryType)
			err := repo.CreateDictionary(&dct)
			if err != nil {
				log.Printf("failed to seed dictionary %s/%s: %v\n", dt.DictionaryType, dictionaryName, err)
			}
			success++
		}
	}
	log.Printf("Seeder : Success to seed %d Dictionary", success)
}
