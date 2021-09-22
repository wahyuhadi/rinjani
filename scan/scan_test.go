package scan

import (
	"testing"

	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/rinjani/directory"
	"github.com/wahyuhadi/rinjani/models"
)

func TestRunnerScan(t *testing.T) {
	t.Run("Scan File .yaml", func(t *testing.T) {
		files := []models.Files{}
		files = append(files, models.Files{Location: "../directory/dir_unitest/", Ext: ".yaml"})
		options := models.Options{Rules: "../directory/dir_unitest/rules_test/"}
		rules, _ := directory.GetFileInRuleDir(options)
		rulesData := directory.RulesToStruct(rules)
		found := RunnerScan(&files, rulesData)
		if len(found) != 0 {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg("expect get 0 length")
			t.Errorf("GetFileInDir() error = %v, wantErr %v", "error file found", false)
			return
		}
	})
}
