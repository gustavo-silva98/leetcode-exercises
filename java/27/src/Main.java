import java.util.ArrayList;
import java.util.Arrays;

public class Main {
    public static void main(String[] args){
        int[] nums = {3,2,2,3};
        int val = 3;
        Solution solution = new Solution();
        int result = solution.removeElement(nums,val);
        System.out.println(result);
        System.out.println(Arrays.toString(nums));
    }
}

class Solution {
    public int removeElement(int[] nums, int val) {
        int intialLenght = nums.length;
        int count = 0;
        ArrayList<Integer> newArray = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == val) {
                nums[i] = 0;
                count++;
                continue;
            }
            newArray.add(nums[i]);
        }
        System.arraycopy(newArray.stream().mapToInt(i -> i).toArray(),0,nums,0,newArray.size());
        return intialLenght- count;
    }
}