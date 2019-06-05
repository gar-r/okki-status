package main

import "bitbucket.org/dargzero/smart-status/providers"

type Config struct {
	Entries []Entry
}

type Entry struct {
	provider providers.Provider
	format   string
}

func NewConfig() *Config {
	return &Config{
		Entries: []Entry{
			{provider: &providers.Clock{Layout: "2006-01-02 15:04"}},
		},
	}
}
