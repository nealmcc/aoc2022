package quickperm

// Permutations returns a channel which will produce permutations of the given input.
// Uses the quick perm algorithm described here: https://www.quickperm.org/
func Permutations[T any](data []T) <-chan []T {
	c := make(chan []T)
	go func(c chan []T) {
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}

func permutate[T any](c chan []T, inputs []T) {
	send := func() {
		output := make([]T, len(inputs))
		copy(output, inputs)
		c <- output
	}
	send()

	// begin the Coundown quickperm algorithm:
	size := len(inputs)

	// p controls the iteration
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}

	for i := 1; i < size; {
		p[i]--
		// if i is odd, then let j = p[i] otherwise let j = 0
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		// swap the values at i and j
		inputs[i], inputs[j] = inputs[j], inputs[i]

		send()

		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}
