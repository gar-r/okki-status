package main

import "bitbucket.org/dargzero/smart-status/modules"

var separator = " "

var activeModules = []modules.Module{
	&modules.Clock{Layout: "2006-01-02 15:04"},
}
