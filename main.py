def missingNumber(nums):
    n = len(nums)
    total = n*(n+1)//2
    return total - sum(nums)


arr1 = [3,0,1]
print("missing number: ", missingNumber(arr1))

def findDisappearedNumbers(nums):
    result = []
    for num in nums:
        index = abs(num) - 1
        if nums[index] >0 :
            nums[index] = -nums[index]
    
    for i in range (len(nums)):
        if nums[i] > 0:
            result.append(i+1)

    return result

def two_sum(nums , target):
    num_map = {}
    for i, num in enumerate(nums):
        complement = target - num
        if complement in num_map:
            return [num_map[complement],i]
        
        num_map[num] = i

    return []

def minTimeToVisitAllPoints(points):
    total_time = 0
    for i in range(1, len(points)):
        x1, y1 = points[i-1]
        x2, y2 = points[i]

        dx = abs(x2 - x1)
        dy = abs(y2 - y1)
        total_time += max(dx, dy)
    return total_time


input_array = [4, 3, 2, 7, 8, 2, 3, 1]
output = findDisappearedNumbers(input_array)
print("missing numbers:", output)

nums = [2, 7, 11, 15]
target = 9
result = two_sum(nums, target)
print("two_sum answer indexs:",result)

points = [[1, 1], [3, 4], [-1, 0]]
points2 = [(3, 2), (-2, 2)]
print("minTimeToVisitAllPoints",minTimeToVisitAllPoints(points))
print("minTimeToVisitAllPoints",minTimeToVisitAllPoints(points2))  