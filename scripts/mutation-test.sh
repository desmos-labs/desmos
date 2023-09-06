#!/usr/bin/env bash

set -eo pipefail
oIFS="$IFS"; IFS=, ; set -- $1 ; IFS="$oIFS"
DISABLED_MUTATORS='branch/*'

# Only consider the following:
# * go files in types, keeper directories
# * ignore test and Protobuf files
# * ignore simulation test files
# * ignore legacy files
# * ignore legacy wasm files
# * ignore testutil files
# * ignore codec.go
go_file_exclusions="-type f -path */keeper/* -or -path */types/* -name *.go -and -not -path */legacy/* -and -not -name *_test.go -and -not -name *pb* -and -not -name codec.go"
MUTATION_SOURCES=$(find ./x $go_file_exclusions )

# Filter on a module-by-module basis as provided by input
arg_len=$#

for i in "$@"; do
  if [ $arg_len -gt 1 ]; then
    MODULE_FORMAT+="./x/$i\|"
    MODULE_NAMES+="${i} "
    let "arg_len--"
  else
    MODULE_FORMAT+="./x/$i"
    MODULE_NAMES+="${i}"
  fi
done

MUTATION_SOURCES=$(echo "$MUTATION_SOURCES" | grep "$MODULE_FORMAT")

# Collect multiple lines into a single line to be fed into go-mutesting
MUTATION_SOURCES=$(echo "$MUTATION_SOURCES" | tr '\n' ' ')

echo "################################################################################"
echo "### WARNING! This test will take hours to complete!"
echo "################################################################################"

echo "running mutation tests for the following module(s): $MODULE_NAMES"
OUTPUT=$(go run github.com/osmosis-labs/go-mutesting/cmd/go-mutesting --disable=$DISABLED_MUTATORS $MUTATION_SOURCES)

# Fetch the final result output and the overall mutation testing score
RESULT=$(echo "$OUTPUT" | grep 'The mutation score')
SCORE=$(echo "$RESULT" | grep -Eo '[[:digit:]]\.[[:digit:]]+')

echo "writing mutation test result to mutation_test_result.txt"
echo "$OUTPUT" > mutation_test_result.txt

# Print the mutation score breakdown
echo $RESULT

# Return a non-zero exit code if the score is below 75%
if (( $(echo "$SCORE < 0.75" |bc -l) )); then
  echo "Mutation testing score below desired level ($SCORE < 0.75)"
  exit 1
fi