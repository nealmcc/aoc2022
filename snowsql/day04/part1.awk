# 
# purpose:
# count the number of lines where one pair of numbers encloses the other.
#

{ if (($1 <= $3 && $4 <= $2) || ($3 <= $1 && $2 <= $4)) { count += 1 } }
END	  { print count }
