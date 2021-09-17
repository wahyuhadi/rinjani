package models

type Options struct {
	// Location is path folder to scan
	// input with flag -l defaul value is current folder
	Location string
	// Ext is extention file would scan
	// input with flag -e defaulr value is ""
	Ext string
	// Rules is folder use to detection
	// yaml file
	Rules string
}
