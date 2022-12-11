default: part1 part2

% : %.awk input.txt
	awk -f $< input.txt

test% : part%.awk test.txt
	awk -f $< test.txt
