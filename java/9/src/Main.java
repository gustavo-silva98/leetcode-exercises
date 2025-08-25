public class Main {
    public static void main(String[] args) {
        Solution solution = new Solution();
        boolean result = solution.isPalindrome(233);
        System.out.println(result);

    }
}

class Solution {
    public boolean isPalindrome(int x) {
        String y = Integer.toString(x);
        String reversedString = "";
        for (int i = y.length() -1; i >= 0; i--) {
            reversedString = reversedString + y.charAt(i);
        }
        return reversedString.equals(y);
    }

}