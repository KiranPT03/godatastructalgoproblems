# Slices

## Operations
1. Iterate over slice
2. Append the item at end
3. Pop the item
4. Insert Item at specific position
5. Delete the item from specific position

### Points to note
1. If we are referencing slice to other function the changes will affect the original slice.
To over come above need to copy the slice first

2. Slices can not be comapred with "==" operator.
But array can be equated.

3. Treat string problems always as arrays.

## Question to practice
1. Reverse the string
2. Merge two sorted arrays
3. Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\
You may assume that each input would have exactly one solution, and you may not use the same element twice.\
You can return the answer in any order.

        Example 1:

        Input: nums = [2,7,11,15], target = 9
        Output: [0,1]

        Output: Because nums[0] + nums[1] == 9, we return [0, 1].

        Example 2:

        Input: nums = [3,2,4], target = 6
        Output: [1,2]

        Example 3:

        Input: nums = [3,3], target = 6
        Output: [0,1]

        **Constraints**:

        * 2 <= nums.length <= 103
        * -109 <= nums[i] <= 109
        * -109 <= target <= 109
        * Only one valid answer exists.

4. Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

        Example 1:

        Input: nums = [2,7,11,15], target = 9
        Output: [0,1]
        Output: Because nums[0] + nums[1] == 9, we return [0, 1].

        Example 2:

        Input: nums = [3,2,4], target = 6
        Output: [1,2]

        Example 3:

        Input: nums = [3,3], target = 6
        Output: [0,1]
        

        **Constraints:**

        * 2 <= nums.length <= 103
        * -109 <= nums[i] <= 109
        * -109 <= target <= 109
        **Only one valid answer exists.**


5. Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

        Example:

        Input: [0,1,0,3,12]
        Output: [1,3,12,0,0]
        
        Note:
        * You must do this in-place without making a copy of the array.
        * Minimize the total number of operations.

6. Given an array of integers, find if the array contains any duplicates.\
Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

        Example 1:

        Input: [1,2,3,1]
        Output: true

        Example 2:

        Input: [1,2,3,4]
        Output: false

        Example 3:

        Input: [1,1,1,3,3,4,3,2,4,2]
        Output: true

7. Given an array, rotate the array to the right by k steps, where k is non-negative.\
Follow up:\
Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.\
Could you do it in-place with O(1) extra space?
 
        Example 1:

        Input: nums = [1,2,3,4,5,6,7], k = 3
        Output: [5,6,7,1,2,3,4]
        Explanation:
        rotate 1 steps to the right: [7,1,2,3,4,5,6]
        rotate 2 steps to the right: [6,7,1,2,3,4,5]
        rotate 3 steps to the right: [5,6,7,1,2,3,4]
        Example 2:

        Input: nums = [-1,-100,3,99], k = 2
        Output: [3,99,-1,-100]
        Explanation: 
        rotate 1 steps to the right: [99,-1,-100,3]
        rotate 2 steps to the right: [3,99,-1,-100]
        

        Constraints:

        1 <= nums.length <= 2 * 104
        -231 <= nums[i] <= 231 - 1
        0 <= k <= 105

8. Have the function LongestWord(sen) take the sen parameter being passed and return the largest word in the string. If there are two or more words that are the same length, return the first word from the string with that length. Ignore punctuation and assume sen will not be empty.

        Examples

        Input: "fun&!! time"
        Output: time

        Input: "I love dogs"
        Output: love