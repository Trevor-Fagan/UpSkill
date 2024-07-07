# Section 2 - Filtering Records
## Where statement
`SELECT name, area FROM cities WHERE area > 4000;`
Postgres executes this query in the following order:
1. FROM cities
2. WHERE area > 4000
3. SELECT name, area

## Comparison Math Operators
- =
- \>
- <
- \>=
- IN => check if value present in a list
- <=
- <> => not equal
- != => not equal
- BETWEEN => `SELECT * FROM users WHERE age BETWEEN 10 AND 21;`
- NOT IN => `SELECT * FROM users WHERE age NOT IN (19, 20, 21);`

## Updating Rows
`UPDATE users SET name='ceira' WHERE name='trevor';`

## Delete Rows
`DELETE FROM users WHERE name='ceira'`