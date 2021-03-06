package models

type Rules struct {
	ID   string `yaml:"id"`
	Info Info   `yaml:"info"`
	File struct {
		Indexwith Indexwith `yaml:"indexwith"`
	} `yaml:"file"`
}

type Info struct {
	Name        string `yaml:"name"`
	Author      string `yaml:"author"`
	Severity    string `yaml:"severity"`
	Details     string `yaml:"details"`
	Impact      string `yaml:"impact"`
	Remediation string `yaml:"remediation"`
}

type Indexwith []struct {
	Type  string   `yaml:"type"`
	Regex []string `yaml:"regex"`
}
