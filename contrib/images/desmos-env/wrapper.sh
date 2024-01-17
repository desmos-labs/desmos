#!/usr/bin/env sh

BINARY=/desmos/${BINARY:-desmos}
ID=${ID:-0}
LOG=${LOG:-desmos.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'desmos'"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export DESMOSDHOME="/desmos/node${ID}/.desmos"

if [ -d "$(dirname "${DESMOSDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${DESMOSDHOME}" "$@" | tee "${DESMOSDHOME}/${LOG}"
else
  "${BINARY}" --home "${DESMOSDHOME}" "$@"
fi
