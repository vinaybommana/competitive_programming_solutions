// twoSum problem

/* using normal arrays
* O(n ^ 2) complexity
*/
public int[] twoSum(int[] nums, int target) {
    for(int i = 0; i < nums.length - 1; i++) {
        for(int j = i + 1; j < nums.length; j++) {
            if(array[i] + array[j] == target) {
                return int [] {i, j};
            }
        }
    }
    throw new IllegalArgumentException("two Sum not found")
}

/* using one-pass Hash Table
* O(n) complexity
*/

public int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> map = new HashMap<>();
    for(int i = 0; i < nums.length; i++) {
        int complement = target - nums[i];
        if (map.containsKey(complement)) {
            return new int[] {map.get(complement), i};
        }
        map.put(nums[i], i);
    }
    throw new IllegalArgumentException("two Sum not found");
}
