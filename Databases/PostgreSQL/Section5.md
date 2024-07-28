## Section 5 - Aggregation of Records
Used to take a large set of rows and condense into a smaller set of values.

### Grouping
- reduces many rows down to fewer rows
- done by using 'GROUP BY'
- visualizing the result is key to use

Using the keywords `GROUP BY` allows you to return results with a unique column name.

```
SELECT id FROM comments GROUP BY id;
```

The above query will return all of the ids of comments with a unique id. When using SELECT with the
`GROUP BY` keywords, you can only SELECT the elements you specified in the GROUP BY clause.

### Aggregates
- reduces many values down to one
- done by using aggregate functions

Examples of aggregate functions: COUNT, SUM, AVG, MIN, MAX

When using COUNT(), if the column name's value is NULL, they are not counted.