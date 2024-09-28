# Section 1 - Introduction

Database -> Collection (Table) -> Documents (Data, schemaless, JSON [BSON])

## MongoDB
- Company behind the database is also called MongoDB
- MongoDB has Self Managed (OSS), Enterprise, Atlas (Cloud), and Mobile versions
- Has a GUI called Compass
- Stitch is a serverless backend solution to work with MongoDB

### Start MongoDB
mongod --dbpath /Users/trevor/mongo/data --logpath /Users/trevor/mongo/logs/mongo.log

You can then connect to the mongo server with ```mongosh```

### Drivers
Allow you to communicate with MongoDB server from a programming language. There are drivers for many of the main languages