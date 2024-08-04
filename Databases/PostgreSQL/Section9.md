# Section 9 - Subqueries
Taking one or more queries and merging them together.

'''
SELECT * FROM products WHERE price > (SELECT SUM(price) FROM products WHERE department = 'Toys')
'''

## Using subqueries in FROM clause
- Subqueries when using with FROM must us an alias
- You can use any subquery so long as the outer selects/wheres/etc are compatible