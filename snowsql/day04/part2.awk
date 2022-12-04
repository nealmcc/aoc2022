# 
# purpose:
# count the number of lines where one pair of numbers overlaps at all with the other
#

{ if (($1 <= $3 && $3 <= $2) || ($3 <= $1 && $1 <= $4)) { count += 1 } }
END	  { print count }
