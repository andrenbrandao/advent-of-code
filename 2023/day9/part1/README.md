# The Problem

We receive a list of numbers and have to predict the next one.

0 3 6 9 12 15 -> 18
 3 3 3 3  3

Every number has a difference of 3. By adding 3 to the last number we have 18.

1   3   6  10  15  21  28
  2   3   4   5   6   7
    1   1   1   1   1
      0   0   0   0

Numbers with more complex variations require us to recursively create sequences
of intermediaries that are used to calculate the next one.

## Algorithm

- Iterate over the original array
- Create an array of differences from the current array
- Repeat this until the resulting subarray is filled with zeroes
- After creating the array, return the current last element + the last element of the new array

This can be done recursively or with a stack. Let's do it first recursively.

Base case:

- If every element in the array is zero, return 0

f(arr) = {
  nextElem = arr[-1] + f(subarray)
  return nextElem
}
