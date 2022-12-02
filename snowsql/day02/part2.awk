# 
# purpose:
# calculate the total score for the 'strategy guide' using the instructions in part 2.
#
# A: Opponent chooses Rock
# B: Opponent chooses Paper
# C: Opponent chooses Scissors
# X: I need to lose
# Y: I need to draw
# Z: I need to win
#

# I lose:
/A X/ { score += 3 + 0 }
/B X/ { score += 1 + 0 }
/C X/ { score += 2 + 0 }

# I draw:
/A Y/ { score += 1 + 3 }
/B Y/ { score += 2 + 3 }
/C Y/ { score += 3 + 3 }

# I win:
/A Z/ { score += 2 + 6 }
/B Z/ { score += 3 + 6 }
/C Z/ { score += 1 + 6 }

END	  { print score }
