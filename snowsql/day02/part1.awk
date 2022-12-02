# 
# purpose:
# calculate the total score for the 'strategy guide' using the instructions in part 1.
#
# A: Opponent chooses Rock
# B: Opponent chooses Paper
# C: Opponent chooses Scissors
# X: I choose Rock
# Y: I choose Paper
# Z: I choose Scissors
#

# I choose Rock:
/A X/ { score += 1 + 3 }
/B X/ { score += 1 + 0 }
/C X/ { score += 1 + 6 }

# I choose Paper:
/A Y/ { score += 2 + 6 }
/B Y/ { score += 2 + 3 }
/C Y/ { score += 2 + 0 }

# I choose Scissors:
/A Z/ { score += 3 + 0 }
/B Z/ { score += 3 + 6 }
/C Z/ { score += 3 + 3 }

END	  { print score }
