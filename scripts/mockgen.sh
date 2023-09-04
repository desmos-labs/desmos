#!/usr/bin/env bash

mockgen_cmd="mockgen"

# Define an array of source and destination paths
declare -a paths=(
    "types/expected_keepers.go:testutil/expected_keepers_mocks.go"
    "types/hooks.go:testutil/hooks_mocks.go"
)

# Loop through the paths array and generate mocks for each source file
for path_pair in "${paths[@]}"; do
    source_path=$(echo "$path_pair" | cut -d ':' -f 1)
    destination_path=$(echo "$path_pair" | cut -d ':' -f 2)

    find ./x -path "*$source_path" | while read -r source_file; do
        destination_file="${source_file/$source_path/$destination_path}"
        destination_dir=$(dirname "$destination_file")
        package=$(basename "$destination_dir")

        mkdir -p "$destination_dir"
        $mockgen_cmd -source "$source_file" -package "$package" -destination "$destination_file"
    done
done