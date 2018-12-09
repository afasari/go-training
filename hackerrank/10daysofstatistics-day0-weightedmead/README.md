Objective 
In the previous challenge, we calculated a mean. In this challenge, we practice calculating a weighted mean. Check out the Tutorial tab for learning materials and an instructional video!

Task 
Given an array, X, of N integers and an array, W, representing the respective weights of X's elements, calculate and print the weighted mean of X's elements. Your answer should be rounded to a scale of 1 decimal place (i.e., 12.3 format).

Input Format

The first line contains an integer, N, denoting the number of elements in arrays X and W. 
The second line contains N space-separated integers describing the respective elements of array X. 
The third line contains N space-separated integers describing the respective elements of array W.

Constraints
5 <= N <= 50
0 < xi <= 100, where xi is the ith element of array X.
0 < wi, where wi is the ith element of array W.
Output Format

Print the weighted mean on a new line. Your answer should be rounded to a scale of 1 decimal place (i.e., 12.3 format).

Sample Input

5
10 40 30 50 20
1 2 3 4 5
Sample Output

32.0
Explanation

We use the following formula to calculate the weighted mean:

mw = 10 x 1 + 40 x 2 + 30 x 3 + 50 x 4 + 20 x 5/ 1 + 2 + 3 + 4 + 5 = 480/15 = 32.0

And then print our result to a scale of 1 decimal place (32.0) on a new line.