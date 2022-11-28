#!/usr/bin/env bash

# shellcheck disable=SC2164
rm -rf docs/02-developers/modules && mkdir -p docs/02-developers/02-modules

for D in ../x/*; do
  if [ -d "${D}" ]; then
    module=$(echo "$D" | awk -F/ '{print $NF}')
    module_dir="./docs/02-developers/02-modules/$module"
    rm -r $module_dir
    if [ -d "${D}/spec" ]; then
      mkdir -p $module_dir
      cp -r "$D/spec/." $module_dir
    fi
  fi
done

## Add architecture documentation
cp -r ./architecture ./docs
