# Section 1 - Simple SQL Statements
## Create a table
```
CREATE TABLE cities (
    name VARCHAR(50),
    country VARCHAR(50),
    population INTEGER,
    area INTEGER
);
```

## Data Types
- VARCHAR: Variable length of characters
- INTEGER: Number without a decimal (~ -2 billion to 2 billion)

## Insert into a table
```
INSERT INTO cities (name, country, population, area) VALUES ('DeLand', 'United States', 25000, 500);
```

Can exclude the first set of parentheses if values align exactly with columns.
You are able to make multiple insertions with a single insert to statement such as the following.

```
INSERT INTO cities VALUES
    ('first', 'second', 500, 200)
    ('third', 'fourth', 850, 100)
```

## Retrieve elements from a table
`SELECT * FROM cities`

## Calculated Columns
`SELECT name, population / area AS population_density FROM cities;`

### Allowed Operators
- +
- -
- *
- /
- ^ => Exponent
- |/ => Square root
- @ => Absolute value
- % => Remainder

## String Operators and Functions
- || => join two strings
  ex: `SELECT name || ', ' || country AS name_country FROM cities;`
- CONCAT() => join two strings
  ex: `SELECT CONCAT(name, ', ', country) AS name_country FROM cities;`
- LOWER() => lower case of string
- LENGTH() => number of characters in a string
  ex: `SELECT LENGTH(country) AS country_length FROM cities;`
- UPPER() => upper case of a string