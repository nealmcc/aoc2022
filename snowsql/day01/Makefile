default: query

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))

all: clean load query

clean:
	rm -f input.csv load

input.csv: input.txt group_input.awk
	awk -f group_input.awk input.txt > input.csv

load: input.csv
	snowsql -c aoc2022 -f load.snowql \
		-o variable_substitution=true \
		-D input_file=$(current_dir)input.csv
	touch load

query: load
	snowsql -c aoc2022 -f query.snowql
