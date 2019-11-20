#!/usr/bin/env sh

##
## Input parameters
##
BINARY=/desmosd/${BINARY:-desmosd}
ID=${ID:-0}
LOG=${LOG:-desmosd.log}

##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'desmosd' E.g.: -e BINARY=desmosd_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

##
## Run binary with all parameters
##
export DESMOSDHOME="/desmosd/node${ID}/desmosd"

if [ -d "$(dirname "${DESMOSDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${DESMOSDHOME}" "$@" | tee "${DESMOSDHOME}/${LOG}"
else
  "${BINARY}" --home "${DESMOSDHOME}" "$@"
fi