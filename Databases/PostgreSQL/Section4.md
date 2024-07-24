# Section 4 - Relating Records with Joins
## Joins + Aggregation
- Joins: Produces values by merging together rows from different related tables
- Aggregation: Looks at many rows and calculates a single value ('most', 'least', 'average')

## Join Syntax
`SELECT contents AS contents, username AS username FROM comments JOIN users ON comments.user_id = users.id;`

The idea behind this is that the comments table is temporarily duplicated and PostgreSQL iterates
over the users and comments table and matches on user_id.

## Join notes:
- Table order between 'FROM' and 'JOIN' makes a difference.
- We must give context if column names collide (SELECT contents.id, users.id)
- Tables can be renamed using the 'AS' keyword (FROM photos AS p) or you can even shorten
    this to be (FROM photos p)
- There are multiple types of joins

## Four Kinds of Joins
### Inner Join 
The default type of join when you use the join keyword. You can also specify that you want to do the inner join
by doing `INNER JOIN`. When items do not line up in both tables, they are dropped from the resulting set. So, in this:
```
SELECT url, username FROM photos JOIN users ON users.id = photos.user_id;
```
You would only get rows that match exactly between the tables (think AND)

### Left Outer Join
You can specify by using the `LEFT JOIN` keywords. So as in the above example, anything from the photos table that
does not match up to the users table, we do not drop the item in the photos table. Instead, the data that would be
matching up with the users table would be filled with NULL.

### Right Outer Join
Same as left outer join, just for the rightmost table.
`RIGHT JOIN`

### Full Join
Match everything that you can, and anything else will be added with the values of just being NULL.