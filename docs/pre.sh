#!/usr/bin/env bash

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
MODULES_DIR="$SCRIPT_DIR/docs/02-developers/02-modules"

# remove modules directories
find "$MODULES_DIR" -mindepth 1 -maxdepth 1 -type d -exec rm -rf '{}' \;

for D in "$SCRIPT_DIR"/../x/*; do
    if [ -d "${D}" ]; then
        module=$(echo "$D" | awk -F/ '{print $NF}')
        module_dir="$MODULES_DIR/$module"
        # Remove the module dir
        rm -r $module_dir
        if [ -d "${D}/spec" ]; then
            # Recreate the directory
            mkdir -p $module_dir
            # Copy the specs
            cp -r "$D/spec/." $module_dir
            # Generate the section with the first letter capitalized
            echo "{\"label\":\"${module^}\",\"collapsed\":true}" > "$module_dir/_category_.json"
        fi
    fi
done

## Add architecture documentation
cp -r ./architecture ./docs
