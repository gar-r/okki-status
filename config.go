package main

import "bitbucket.org/dargzero/smart-status/providers"

// Config contains the required settings to start the application
type Config struct {
	Entries []Entry
}

// Entry is a single configuration entry
type Entry struct {
	provider providers.Provider
	format   string
}

// NewConfig creates a new Config instance
func NewConfig() *Config {
	return &Config{
		Entries: []Entry{
			{provider: &providers.Clock{Layout: "2006-01-02 15:04"}},
		},
	}
}
