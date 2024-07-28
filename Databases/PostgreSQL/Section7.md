# Section 7 - Sorting Records
You can utilize the `ORDER BY` keywords in order to sort a resulting dataset. Along with this you can
use ASC or DESC. The default option is ASC.

```
SELECT * FROM products ORDER BY price;
```

You can also sort on string values and it will sort by alphabetical order.
You can also have multiple order conditions, which will first match by the first
and then go to the second to break ties.

```
SELECT * FROM products ORDER BY price, weight;
```

## Offset
Offset skips the first x number of records in a resulting dataset

## Limit
Only gives the first two rows of the result set.

```
SELECT * FROM users OFFSET 45 LIMIT 2;
```