# 
# purpose:
# transform the input.txt into a csv file like this:
#
#  parent   |   name    |  size
# -----------------------------------
#  /        |  a        | 
#  /        |  b        | 14848514
#  /        |  c        | 8504156
#  /        |  d        | 
#  /a       |  e        | 
#  /a       |  f        | 29116
#  /a       |  g        | 
#  /a       |  h.lst    | 62596
#  /a/e     |  i        | 584
#

function path_push(val) {
	path[path_len++] = val;
}

function path_pop() {
	if (path_len <= 0) {
		return ""
	}
	return path[--path_len]
}

function path_string() {
	if (path_len == 0) {
		return "/"
	}
	res = ""
	for ( i = 0 ; i < path_len ; i++ ) {
		res = res "/" path[i]
	}
	return res
}


BEGIN           { print "parent,name,size" }
/\$ cd [a-z]+/  { path_push($3) }
/\$ cd \.\./    { path_pop()   }
/^dir/          { print path_string() "," $2 "," }
/^[0-9]+/       { print path_string() "," $2 "," $1 }
