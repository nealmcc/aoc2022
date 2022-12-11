# 
# purpose:
# solve part 1
#

function trace() {
	if ((cycle-20)%40 == 0) {
		ss = cycle*x
		print "during c:" cycle " x:" x " ss:" ss
		sum += ss
	}
}

function tick() {
    trace()
    cycle++
}

BEGIN  { cycle=1; x=1; }
/noop/ { tick() }
/addx/ { tick() ; tick() ; x += $2 }
END    { print sum }
