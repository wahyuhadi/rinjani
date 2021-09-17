package directory

import (
	"testing"

	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/rinjani/models"
)

func TestRulesToStruct(t *testing.T) {
	t.Run("Filter only valid format rules with yaml lint", func(t *testing.T) {
		opts := models.Options{Rules: "../directory/dir_unitest/rules_test"}
		rules, _ := GetFileInRuleDir(opts)
		got := RulesToStruct(rules)
		for _, items := range *got {
			if items.Info.Name != "success" {
				gologger.Info().Str("state", "errored").Str("status", "error").Msg("expect get success")
				t.Errorf("GetFileInDir() error = %v, wantErr %v", "error file found", false)
				return
			}
		}
	})

}
