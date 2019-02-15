#!/bin/bash

cp * ./script

for D in *; do
	if [ -d "${D}" ]; then
		if [ "${D}" != "Servers" ]; then		
		  cd "${D}"
		  echo "---===<<<*>>>===--- ${D} ----===<<<*>>>===---"
                  expect ../sync
		  cd ..
		fi
	fi
done
