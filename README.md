# Problem

You need to calculate the distance in whole days between two dates, counting only
the days in between those dates, i.e. 01/01/2001 to 03/01/2001 yields “1”. The valid
date range is between 01/01/1900 and 31/12/2999, all other dates should be
rejected.
When testing your solution, use the following sample data to demonstrate your code
works:
a) 2/6/1983 to 22/6/1983 19 days
b) 4/7/1984 to 25/12/1984 173 days
c) 3/1/1989 to 3/8/1983 2036 days

## Solution
To calculate the number of days between two different dates.
First: We calculate the number of days between different years and all 12 months.
Second:  then we calculate the months that are added more than the first and last date e.g (3/1/1989 to 3/8/1983)
We calculate days from Month 8 and the last month (12) then we
Subtract 8 - 12 and calculate days between these months and
subtract from a total of the days we already calculated from years
then we subtract of extra days from and to date
and then subtract from the total of the days
at the end, we subtract all days with 1