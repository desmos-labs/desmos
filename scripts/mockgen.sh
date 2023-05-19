#!/usr/bin/env bash

mockgen_cmd="mockgen"
search_path="/types/expected_keepers.go"

find ./x -path "*$search_path" | while read -r source_file; do
    destination_file="${source_file/types\/expected_keepers.go/testutil\/expected_keepers_mocks.go}"
    destination_dir=$(dirname "$destination_file")
    package=$(basename "$destination_dir")

    mkdir -p "$destination_dir"
    $mockgen_cmd -source "$source_file" -package "$package" -destination "$destination_file"
done