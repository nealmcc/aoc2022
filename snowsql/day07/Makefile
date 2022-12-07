default: test

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))

all: clean load query

test: clean test.csv
	cat test.csv

clean:
	rm -f load test.csv input.csv

test.csv: gen_csv.awk test.txt
	awk -f gen_csv.awk test.txt > test.csv

input.csv: gen_csv.awk input.txt
	awk -f gen_csv.awk input.txt > input.csv

load: input.csv load.snowql
	snowsql -c aoc2022 -f load.snowql \
		-o variable_substitution=true \
		-D input_file=$(current_dir)input.csv
	touch load

query: load query.snowql
	snowsql -c aoc2022 -f query.snowql
