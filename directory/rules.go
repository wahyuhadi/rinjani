package directory

import (
	"io/ioutil"
	"runtime"
	"sync"

	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/rinjani/constanta"
	"github.com/wahyuhadi/rinjani/models"
	"gopkg.in/yaml.v2"
)

var rules []models.Rules

var wg sync.WaitGroup

func RulesToStruct(rule *[]models.Files) *[]models.Rules {
	n := runtime.GOMAXPROCS(constanta.MaxProcess)
	wg.Add(1)
	go func(n int) {
		for _, items := range *rule {
			yamlFile, err := ioutil.ReadFile(items.Location)
			if err != nil {
				gologger.Warning().Msgf("Error parsing YAML file in %v", items.Location)
				continue
			}
			var yamlConfig models.Rules
			err = yaml.Unmarshal(yamlFile, &yamlConfig)
			if err != nil {
				gologger.Info().Msgf("Error parsing YAML file in %v", items.Location)
				continue
			}
			rules = append(rules, yamlConfig)
		}
		wg.Done()
	}(n)
	wg.Wait()
	return &rules
}
