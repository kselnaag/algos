#!/usr/bin/env bash

echo -e "\n>>_Mirroring _Ord_ files_<<"
ORDFILES=`find . -type f -name "*Ord.go"`
for FILE in ${ORDFILES[*]}
do
	OUT="${FILE%Ord.go}Comp.go"
	sed -e "s/I\.Ord/I\.Comp/g;s/\[\]T/\[\]\*T/g;s/Ord/Comp/g" $FILE > $OUT
	echo "$FILE -> $OUT"
done