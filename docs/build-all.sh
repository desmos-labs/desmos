#!/usr/bin/env bash

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
# Current commit
COMMIT=$(git rev-parse HEAD)

function is_legacy_version() {
    # Versions with the old documentation schema
    LEGACY_VERSIONS=("4.2.0" "2.3")
    version=$1
    for legacy_version in "${LEGACY_VERSIONS[@]}"; do
        if [ "$legacy_version" == "$version" ]; then
            echo "true"
            return 0
        fi
    done

    echo "false"
    return 0
}

# Clear old cached docs
rm -Rf ~/versioned_docs ~/versioned_sidebars
mkdir -p ~/versioned_docs ~/versioned_sidebars

echo "getting legacy documentation"
git clone https://github.com/desmos-labs/docs.git "$SCRIPT_DIR/legacy-docs"
# Cache the legacy documentation
mv "$SCRIPT_DIR/legacy-docs/versioned_docs/"* ~/versioned_docs/
mv "$SCRIPT_DIR/legacy-docs/versioned_sidebars/"* ~/versioned_sidebars/
# Remove the legacy docs files
rm -Rf "$SCRIPT_DIR/legacy-docs"

for version in $(jq -r .[] "$SCRIPT_DIR/versions.json"); do
    is_legacy=$(is_legacy_version $version)
    if ! $is_legacy ; then
        echo "building docusaurus $version docs"
        git clean -fdx && git reset --hard && git checkout "v$version"
        sh ./pre.sh
        # Remove the versions.json file to ensure that the documentation will be generated properly.
        # This is needed otherwise the `docusaurus docs:version` command will fail if
        # you are trying to generate the documentation for a version that is already present in the version.json file.
        rm "$SCRIPT_DIR/versions.json"
        yarn install --frozen-lockfile && yarn run docusaurus docs:version $version
        mv ./versioned_docs/* ~/versioned_docs/
        mv ./versioned_sidebars/* ~/versioned_sidebars/
    fi
done
echo "building docusaurus main docs"
(git clean -fdx && git reset --hard && git checkout $COMMIT)
mv ~/versioned_docs ~/versioned_sidebars .
./pre.sh
yarn install --frozen-lockfile && NODE_OPTIONS=--max-old-space-size=8192 yarn build

mv build ~/output
echo "setup domain"
echo $DOCS_DOMAIN > ~/output/CNAME
