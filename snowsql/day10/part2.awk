# 
# purpose:
# solve part 2
#

function draw(col) {
	if (col < x-1 || col > x+1) {
		printf "."
	} else {
		printf "#"
	}
	if (col == 39) { print ""}
}

function tick() {
	cycle++
	draw(col)
	col = (col+1) % 40
}

BEGIN  { cycle=1; x=1; col=0; }
/noop/ {  tick() }
/addx/ {  tick() ; tick() ; x += $2 }

