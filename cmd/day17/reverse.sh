#!/bin/sh

if [ "$#" -eq 0 ]; then
  echo "Usage: $0 <pattern>" >&2
  exit 1
fi

for filename in "$@"; do
	awk '{a[i++]=$0} END {for (j=i-1; j>=0;) print a[j--] }' "$filename" > "$filename".reverse
done

for filename in *.reverse; do
	mv "$f" "$(echo "$f" | sed s/\.reverse//)"
done
