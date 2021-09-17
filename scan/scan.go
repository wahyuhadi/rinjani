package scan

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"sync"

	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/rinjani/constanta"
	"github.com/wahyuhadi/rinjani/models"
)

var wg sync.WaitGroup

// RunnerScan
func RunnerScan(files *[]models.Files, rules *[]models.Rules) [][]string {
	var found [][]string
	for _, isfile := range *files {
		file, err := os.Open(isfile.Location)
		if err != nil {
			gologger.Info().Msgf("Error read file in %v", isfile.Location)
			continue
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var text []string
		for scanner.Scan() {
			text = append(text, scanner.Text())
		}

		file.Close()
		for line, each_ln := range text {
			n := runtime.GOMAXPROCS(constanta.MaxProcess)
			wg.Add(1)
			go func(n int) {
				// get rules
				for _, i := range *rules {
					// looping in rules index or detect with.
					for _, index := range i.File.Indexwith {
						// if rules detect with regex
						if index.Type == constanta.RegexType {
							for _, re := range index.Regex {
								// match with regex in file
								match, _ := regexp.MatchString(re, each_ln)
								// if match
								if match {
									isFind := regexp.MustCompile(re)
									word := isFind.Find([]byte(each_ln))
									found = append(found,
										[]string{
											i.ID,
											fmt.Sprintf("%v:%v", isfile.Location, line),
											string(word),
											i.Info.Severity,
											i.Info.Remediation,
										},
									)
								}
							}
						}
						// not regex type
						// To Do
					}
				}
				wg.Done()
			}(n)
			wg.Wait()
		}
	}

	return found
}
