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
