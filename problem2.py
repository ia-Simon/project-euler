def memoize(fn):
	cache = {}

	def inner (n):
		if n in cache:
			return cache[n]

		out = fn(n)
		cache[n] = out
		return out

	return inner 

@memoize
def fib(n):
	if (n == 0):
		return 0

	if (n == 1):
		return 1

	return fib(n-1) + fib(n-2)


def tailFib(n):
	last = 0
	accum = 1
	for _ in range(n-1):
		[last, accum] = [accum, last+accum]

	return accum

print(fib(42))
# print(tailFib(42))
