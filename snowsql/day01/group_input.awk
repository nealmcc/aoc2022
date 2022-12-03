# 
# purpose:
# insert each elf's ID and convert to CSV
# 

BEGIN { print "elf,food" ; elf=1 }
NF    { print elf "," $0 }
!NF   { elf+=1 }
