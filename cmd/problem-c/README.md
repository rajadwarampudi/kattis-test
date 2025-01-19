This sub directory contains the solution for pizza slicer problem.

Input
Input begins with an integer
 indicating the number of test cases that follow. The following
 lines each contain one test case. Each test case gives the pizza radius in centimeters,
, followed by the number of people sharing the pizza,
, followed by the rotation angle,
. The quantities
,
 and
 are all positive. The value
 is an integer no greater than 100, and
 is an integer no greater than
. The angle
 is given as an integer number of degrees, followed by an integer number of minutes and an integer number of seconds. Degrees are between 0 and 359 (inclusive), while minutes and seconds are between 0 and 59 (inclusive).

Output
For each test case, print the area in square centimeters of the largest resulting slice of pizza. You do not need to worry about the precise formatting of the answer (e.g. number of places past the decimal), but the absolute error of your output must be smaller than .0001

Example for sample input:
4
20 6 60 0 0
20 6 59 59 59
30 20 33 30 0
40 200 120 0 0

The output will be

209.439510
209.438541
392.699082
392.699082
