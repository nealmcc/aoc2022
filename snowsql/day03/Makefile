default: query

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))

all: clean load query

clean:
	rm -f load

load: input.txt items.csv
	snowsql -c aoc2022 -f load.snowql \
		-o variable_substitution=true \
		-D input_items=$(current_dir)items.csv \
		-D input_bags=$(current_dir)input.txt
	touch load

query: load
	snowsql -c aoc2022 -f query.snowql
