#!/bin/sh
cat settings.go | grep '"version"' | sed 's/^.*("version", "\(.*\)")$/\1/'
