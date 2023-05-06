ints = list(map(int,input().split()))
nums = dict()

for i in ints:
	if i not in nums:
		nums[i]=None

print(len(nums))