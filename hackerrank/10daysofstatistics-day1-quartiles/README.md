Objective 
In this challenge, we practice calculating quartiles. Check out the Tutorial tab for learning materials and an instructional video!

Task 
Given an array, X, of n integers, calculate the respective first quartile (Q1), second quartile (Q2), and third quartile (Q3). It is guaranteed that Q1, Q2, and Q3 are integers.

Input Format

The first line contains an integer, n, denoting the number of elements in the array. 
The second line contains n space-separated integers describing the array's elements.

Constraints

5 <= n <= 50
0 < xi <= 100, where xi is the ith element of the array. 
Output Format

Print 3 lines of output in the following order:

The first line should be the value of Q1.
The second line should be the value of Q2.
The third line should be the value of Q3.
Sample Input

9
3 7 8 5 12 14 21 13 18
Sample Output

6
12
16
Explanation

X = {3, 7, 8, 5, 12, 14, 21, 13, 18}. When we sort the elements in non-decreasing order, we get X={3,5,7,8,12,13,14,18,21}. It's easy to see that median(X) = 12.
As there are an odd number of data points, we do not include the median (the central value in the ordered list) in either half:

Lower half (L): 3, 5, 7, 8

Upper half (U): 13, 14, 18, 21

Now, we find the quartiles:

Q1 is the median(L). So, Q1 = 5 + 7/2 = 6.
Q2 is the median(X). So, Q2 = 12.
Q3 is the median(U). So, Q3 = 14+18/2=16.