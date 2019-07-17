#!/bin/sh
if [ $# -eq 2 ]
  then
	eval $2
	smart-status --refresh=$1
fi
