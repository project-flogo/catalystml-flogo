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
    Aggregate - map[string]string : Map key is groupKey, map value is aggregate function. Currently support Sum, Count, Mean, Min, Max
    Level - int : Group level in above <0 -> group by both Column1 & Column2, 0 -> group by Column1, 1 -> group by Column2.

OutputType - map[string][]interface{} : grouped dataframe

## Compliance to Spec

### Rough level of compliance  
100%