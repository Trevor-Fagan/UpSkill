# Section 8 - Unions and Intersections with Sets
## Union
Basic set theory, what you would expect.
```
(SELECT *
FROM products 
ORDER BY price DESC
LIMIT 4)
UNION
(SELECT *
FROM products
ORDER BY (price / weight) DESC
LIMIT 4);
```
Union by default will remove any duplicates. You can use
`UNION ALL`
to not remove any duplicates.

You are only able to use the union keyword on two queries whose resulting columns are the same
in name and type.

## Set Keywords
- UNION: join together the results of two queries and remove duplicate rows
- UNION ALL: join together the results of two queries and keep duplicate rows
- INTERSECT: find rows common in the results of two queries, remove duplicates
- INTERSECT ALL: find common rows, don't remove duplicates
- EXCEPT: find the rows that are present in first query but not in second query, remove duplicates
- EXCEPT ALL: find the rows that are present in first query but not in second query, remove duplicates