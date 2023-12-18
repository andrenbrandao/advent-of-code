# The Problem

Now we receive a list of numbers and have to predict the previous one

-3 <-- 0 3 6 9 12 15
        3 3 3 3  3

## Algorithm

If we follow the same approach as the other one but try to insert an element in the beginning of the array, the time complexity of each operation will be O(N) because of having to shift the whole array to the right.

- Iterate over the original array
- Create an array of differences from the current array
- Repeat this until the resulting subarray is filled with zeroes
- After creating the array, return the current first element - the return value of the function called with the subarray

Base case:

- If every element in the array is zero, return 0

prev(arr) = {
  prevElem = arr[0] - prev(subarray)
  return prevElem
}
