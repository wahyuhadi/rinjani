package directory

import (
	"testing"

	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/rinjani/constanta"
	"github.com/wahyuhadi/rinjani/models"
)

func TestGetFileInDir(t *testing.T) {
	t.Run("Run with non existing folder", func(t *testing.T) {
		opts := &models.Options{Location: "no folder in your xxxxxx", Ext: ".go"}
		_, err := GetFileInDir(*opts)
		if err == nil {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
			t.Errorf("GetFileInDir() error = %v, wantErr %v", err, true)
			return
		}

		if err.Error() != constanta.PathNotFound {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
			t.Errorf("GetFileInDir() error message = %v, want error message %v", err, constanta.PathNotFound)
			return
		}

	})

	t.Run("Run with no file in folder", func(t *testing.T) {
		opts := models.Options{Location: "../directory/dir_unitest/empty_file", Ext: ".go"}
		_, err := GetFileInDir(opts)
		if err == nil {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
			t.Errorf("GetFileInDir() error = %v, wantErr %v", err, true)
			return
		}

		if err.Error() != constanta.FileNotFound {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
			t.Errorf("GetFileInDir() error message = %v, want error message %v", err.Error(), constanta.FileNotFound)
			return
		}

	})

	t.Run("Run with only single file", func(t *testing.T) {
		opts := models.Options{Location: "../directory/dir_unitest/single_file", Ext: ".go"}
		got, _ := GetFileInDir(opts)
		if got == nil {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg("expect return no nil")
			t.Errorf("GetFileInDir() error = %v, wantErr %v", nil, false)
			return
		}

		for _, data := range *got {
			if data.Location != "../directory/dir_unitest/single_file/1.go" {
				gologger.Info().Str("state", "errored").Str("status", "error").Msg(data.Location)
				t.Errorf("GetFileInDir() error = %v, wantErr %v", nil, false)
				return
			}
		}

	})

	t.Run("Filter with .go extention", func(t *testing.T) {
		opts := models.Options{Location: "../directory/dir_unitest/", Ext: ".go"}
		got, _ := GetFileInDir(opts)
		if got == nil {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg("expect return no nil")
			t.Errorf("GetFileInDir() error = %v, wantErr %v", nil, false)
			return
		}

		for _, data := range *got {
			if data.Ext != ".go" {
				gologger.Info().Str("state", "errored").Str("status", "error").Msg("expect get .go")
				t.Errorf("GetFileInDir() error = %v, wantErr %v", "only get .go", false)
				return
			}
		}

	})

}

func TestGetFileInRuleDir(t *testing.T) {
	t.Run("Run with non existing folder", func(t *testing.T) {
		opts := &models.Options{Rules: "no folder in your xxxxxx"}
		_, err := GetFileInRuleDir(*opts)
		if err == nil {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
			t.Errorf("GetFileInDir() error = %v, wantErr %v", err, true)
			return
		}

		if err.Error() != constanta.RulePathNotFOund {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
			t.Errorf("GetFileInDir() error message = %v, want error message %v", err, constanta.PathNotFound)
			return
		}

	})

	t.Run("Run with no file in folder", func(t *testing.T) {
		opts := models.Options{Rules: "../directory/dir_unitest/empty_file"}
		_, err := GetFileInRuleDir(opts)
		if err == nil {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
			t.Errorf("GetFileInDir() error = %v, wantErr %v", err, true)
			return
		}

		if err.Error() != constanta.RuleFileNotFound {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
			t.Errorf("GetFileInDir() error message = %v, want error message %v", err.Error(), constanta.RuleFileNotFound)
			return
		}

	})

	t.Run("Run with only single file", func(t *testing.T) {
		opts := models.Options{Rules: "../directory/dir_unitest/dir_rule"}
		got, _ := GetFileInRuleDir(opts)
		if got == nil {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg("expect return no nil")
			t.Errorf("GetFileInDir() error = %v, wantErr %v", nil, false)
			return
		}

		for _, data := range *got {
			if data.Location != "../directory/dir_unitest/dir_rule/1.yaml" {
				gologger.Info().Str("state", "errored").Str("status", "error").Msg(data.Location)
				t.Errorf("GetFileInDir() error = %v, wantErr %v", nil, false)
				return
			}
		}

	})

	t.Run("Filter with .go extention", func(t *testing.T) {
		opts := models.Options{Rules: "../directory/dir_unitest/"}
		got, _ := GetFileInRuleDir(opts)
		if got == nil {
			gologger.Info().Str("state", "errored").Str("status", "error").Msg("expect return no nil")
			t.Errorf("GetFileInDir() error = %v, wantErr %v", nil, false)
			return
		}

		for _, data := range *got {
			if data.Ext != ".yaml" {
				gologger.Info().Str("state", "errored").Str("status", "error").Msg("expect get .go")
				t.Errorf("GetFileInDir() error = %v, wantErr %v", "only get .go", false)
				return
			}
		}

	})
}
