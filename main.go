package main

import (
	"github.com/NomadXD/go-native-plugin/pkg/plugin"
)

// SriLanka implements Country interface
type SriLanka string

// Name returns the name of the country
func (l SriLanka) Name() string {
	return "Sinhala"
}

// Continent returns the Continent of the country
func (l SriLanka) Continent() string {
	return "Asia"
}

// Language returns the Language of the country
func (l SriLanka) Language() string {
	return "Sinhala"
}

// Greeting returns the Greeting of the country
func (l SriLanka) Greeting() string {
	return "Ayubowan!"
}

// TimeZone returns the TimeZone of the country
func (l SriLanka) TimeZone() string {
	return "UTC+5.30"
}

// Currency returns the Currency of the country
func (l SriLanka) Currency() string {
	return "LKR"
}

// Init handles the plugin initialization logic
func (l SriLanka) Init() {
	plugin.RegisterPlugin("Sinhala")
}

// SriLankaImpl Exported Implementation
var SriLankaImpl SriLanka
