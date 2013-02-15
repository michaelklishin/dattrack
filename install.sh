#!/bin/sh

go build .
echo "Installing to ~/bin/dattrack"
mv dattrack.go.git ~/bin/dattrack
echo "Done"
