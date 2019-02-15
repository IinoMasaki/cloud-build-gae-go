#!/bin/bash

. /builder/prepare_workspace.inc
prepare_workspace || exit
echo "Running: dep $@"
dep "$@"
