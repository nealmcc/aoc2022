default: part1 part2

test: test1 test2

% : %.awk input.txt
	awk -F',|-' -f $< input.txt

test1:
	awk -F',|-' -f part1.awk test.txt	

test2:
	awk -F',|-' -f part2.awk test.txt	
