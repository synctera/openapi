#!/usr/bin/env bash

if [ "${#}" != 2 ]; then
  echo "Usage: ${0} <internal/external> <generator-name>"
  exit 1
fi

if [ "${1}" != internal ] && [ "${1}" != external ]; then
  echo "Usage: ${0} <internal/external> <generator-name>"
  exit 1
fi

config_file=client/"${2}".config.json
if [ ! -f "${config_file}" ]; then
    echo "${config_file}" not found, creating with empty config
    echo {} > "${config_file}"
fi

log_file=$(mktemp --suffix .openapi-generator-cli.log)

openapi-generator-cli generate --strict-spec true --generator-name "$2" \
  --input-spec spec/"$1"-api-merged-bundled.yml --output client/"$1"/"$2"/ \
	--package-name synctera --config "${config_file}" 2>&1 | tee "${log_file}"

bad_warnings=$(grep WARN "${log_file}" | grep --invert-match "error (reserved word) cannot be used as model name. Renamed to model_error$")
if [ "${bad_warnings}" != "" ]; then
  echo
  echo openapi-generator-cli produced the following warnings:
  echo
  echo "${bad_warnings}"
  exit 1
fi
