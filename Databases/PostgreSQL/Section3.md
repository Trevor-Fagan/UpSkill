# Section 3 - Working with Tables

## Types of Relationships
- One-to-many: A user has many different photos
- Many-to-one: From perspective of photo, many photos belong to one user
- One-to-one: One boat has one captain
- Many-to-many: Students have many classes, and classes have many students

## Primary Keys and Foreign Keys
**Primary Key:** Uniquely identifies record in table
- most of the time just called ID
- will never change
- either integer or UUID

By specifying "Serial", PostgreSQL will automatically fill in this column
for us and auto increment.

```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50)
);
```

**Foreign Key:** Identifies a record in another table that the row is associated with
- represents "belonging" to another record
- many rows can have the same foreign key
- equal to the primary key of the referenced row

```
CREATE TABLE photos (
    id SERIAL PRIMARY KEY,
    url VARCHAR(255),
    user_id INTEGER REFERENCES users(id)
);
```

## Foreign Key Constraints
When trying to use a foreign key constraint and inserting into a table, it must be
valid for the foreign key. For example if you have a photos table and a users table,
where each photo references an id in the users table, the id must be present in the
users table to be valid or you may get a similar error to the following:

`insert or update on table "photos" violates foreign key constraint "photos_user_id_fkey"`

If we want to insert into a table that uses a foreign key column, but we don't want to
follow the foreign key, we can use `NULL`.

```
INSERT INTO photos (url, user_id)
VALUES
	('photo-site.org/242423432', NULL);
```

### Deletion options for foreign keys
- ON DELETE RESTRICT => Throw an error. This is the default option.
- ON DELETE NO ACTION => Throw an error
- ON DELETE CASCADE => Delete the item referenced by foreign key as well
- ON DELETE SET NULL => Set the foreign key to NULL
- ON DELETE SET DEFAULT => Set the foreign key to the default value

Example:
```
CREATE TABLE photos (
id SERIAL PRIMARY KEY,
url VARCHAR(200),
user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
);

DELETE FROM users WHERE id=1;
```
*This will delete the user with id 1 as well as any photos with 1 as the foreign key*

### Set Default Values
```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) DEFAULT NULL
);
```