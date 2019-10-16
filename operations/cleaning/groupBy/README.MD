# groupBy

## Overview
group by a given column in an axis and aggregate the value of another column (like SQL) 

### Description
group by works same as following SQL command

	SELECT Column1, Column2, mean(Column3)
	FROM DataTable
	GROUP BY Column1, Column2

### Implementation details

Input
    Data - map[string][]interface{} : map key is column name (in above SQL Column1, Column2 and Column3), map value are column values. Optional=False
Params
    Index - []string : Group key of data (in above SQL they are Column1 and Column2).
    Target - string : aggreagte column (Column3 in above SQL)
    Function - string : aggreagte function (mean in above SQL). Currently sum, count, mean, max and min are suported.
    Level - int : Group level in above 0 -> group by Column1, 1 -> group by Column2.

OutputType - map[string][]interface{} : grouped dataframe

## Compliance to Spec

### Rough level of compliance  
70%