#!/usr/bin/env bash

function is_legacy_version() {
  version=$1
  # Versions wich the documentation is already present inside the "versioned_*" folders
  legacy_versions=("2.3" "4.2.0")

  for legacy_version in "${legacy_versions[@]}"; do
    if [ "$legacy_version" == "$version" ] ; then
      echo "true"
      return 0
    fi
  done

  echo "false"
  return 0
}

# This builds the docs.desmos.network docs using docusaurus.
# Old documentation, which have not been migrated to docusaurus are generated with vuepress.
COMMIT=$(git rev-parse HEAD)

# Create directories where will be stored the old versions
mkdir -p ~/versioned_docs  ~/versioned_sidebars

for version in $(jq -r .[] versions.json); do
  is_legacy=$(is_legacy_version "$version")
  if [ ! $is_legacy ] ; then
    echo "not legacy version $version"
    echo "building docusaurus $version docs"
    git clean -fdx && git reset --hard && git checkout release/$version.x
    sh ./pre.sh
    yarn install && yarn run docusaurus docs:version $version
    # Archive the built versions
    mv ./versioned_docs/* ~/versioned_docs/
    mv ./versioned_sidebars/* ~/versioned_sidebars/
  fi
done

echo "building docusaurus main docs"
# git clean -fdx && git reset --hard && git checkout $COMMIT
mv ~/versioned_docs/* ./versioned_docs/
mv ~/versioned_sidebars/* ./versioned_sidebars/
./pre.sh
yarn install && yarn build
# mv build ~/output
