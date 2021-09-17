package directory

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/wahyuhadi/rinjani/constanta"
	"github.com/wahyuhadi/rinjani/models"
)

func GetFileInDir(options models.Options) (*[]models.Files, error) {
	var files []models.Files
	// checking is path exist
	if _, err := os.Stat(options.Location); os.IsNotExist(err) {
		return nil, errors.New(constanta.PathNotFound)
	}
	// parsing the file inside the nested path
	err := filepath.Walk(options.Location, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// if no ext
			if options.Ext == "" && filepath.Ext(path) != ".yaml" {
				files = append(files, models.Files{
					Location: path,
					Size:     info.Size(),
					Ext:      filepath.Ext(path),
					Time:     info.ModTime(),
					Name:     info.Name(),
				})
			} else {
				// filter with ext
				if filepath.Ext(path) == options.Ext {
					files = append(files, models.Files{
						Location: path,
						Size:     info.Size(),
						Ext:      filepath.Ext(path),
						Time:     info.ModTime(),
						Name:     info.Name(),
					})
				}
			}
		}
		return err
	})

	if len(files) == 0 {
		return nil, errors.New(constanta.FileNotFound)
	}

	return &files, err

}

func GetFileInRuleDir(options models.Options) (*[]models.Files, error) {
	var files []models.Files
	// checking is path exist
	if _, err := os.Stat(options.Rules); os.IsNotExist(err) {
		return nil, errors.New(constanta.RulePathNotFOund)
	}

	// parsing the file inside the nested path
	err := filepath.Walk(options.Rules, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// if no ext
			if filepath.Ext(path) == constanta.ExtRules {
				files = append(files, models.Files{
					Location: path,
					Size:     info.Size(),
					Ext:      filepath.Ext(path),
					Time:     info.ModTime(),
					Name:     info.Name(),
				})
			}

		}
		return err
	})

	if len(files) == 0 {
		return nil, errors.New(constanta.RuleFileNotFound)
	}
	return &files, err

}
