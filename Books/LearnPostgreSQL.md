# Learn PostgreSQL
## Use, manage, and build secure and scalable databses with PostgreSQL 16
## Luca Ferrari and Enrico Pirozzi

## Useful Commands
- `ps -ef | grep postgres`
- `sudo chown -R ubuntu:ubuntu /var/run/postgresql`
- `sudo -u ubuntu pg_ctl reload -D test`              => reload after updating pga_hba.conf

## 1 - Introduction
- Open source
- ACID compliant
    A- Atomicity: complex operations processed as a single instruction even when composed of different operations
    C- Consistency: Data is always consistent- it is not corrupted by partial operations
    I- Isolation: Handles concurrency properly; No corruption from interleaved changes
    D- Durability: Protects data it contains as well as possible
- PostgreSQL runs as a service (daemon) on a server and is referred to as a cluster since it can contain multiple databases
- Databases can have “namespaces” called schemas, which is just a way of organizing database objects (such as tables) in a structured collection. They cannot be nested
- Database objects: tables, functions, triggers, and data types
- Every object belongs to only one schema and if not specified, is named as the user that creates the object
- Catalogs: special tables and views that present information in a SQL-interactive way. Tracks tables’ structures, indexes, functions, dedicated storage, etc.
- PostgreSQL has normal users and superusers
- PostgreSQL stores all of its content (user data and internal status) in a single filesystem directory called PGDATA
- PGDATA contains Write-Ahead Logs (WALs) and the data storage
- WALs- Before applying changes to a chunk of data, an intent log will be made persistent so if a failure occurs, can reference the log to recover
- When initialized, a PostgreSQL cluster starts a single process called the postmaster, and every user connection is then controlled by the postmaster that created a background process to handle requests

## 2 - Getting to Know Your Cluster
- Pg_ctl - command line utility that allows you to perform different actions on a cluster (initialize, start, restart, stop, etc). To use, need to add to path to .profile. This command interacts with the postmaster and sends requests to background processes
    - This command takes either -d path, or uses environment variable $PGDATA in order to know where the PGDATA is stored on disk
    - The user that you are going to have access to the data cannot be root and must be given access
- Pg_ctl commands
    - Init  => initialize pg data
    - Start => start the postgresql server
    - Stop  => stop the postgresql server
        - Different options for this as it can be dangerous to stop a server. Use -m and one of the following
        - Smart - waits for connected clients to disconnect and then shuts down
        - Fast - immediately disconnect every client and will shut down the server
        - Immediate - aborts every PostgreSQL process, including client connections and shut down in dirty way
- PostgreSQL Processes
    - Postmaster is root of all postgresql processes
    - Use the following to view all processes tied to postgresql:
        - Pstree -p postgres
        - Ps -C postgres -af
    - The processes:
        - Checkpointer: executes checkpoints, which are points in time where the DB ensures that all the data is actually stored persistently on disk
        - Background writer: pushes data out of memory to permanent storage
        - Walwriter: writing the Write Ahead Logs
        - Logical replication launcher: process responsible for handling logical replication
        - Autovacuum launcher: launches the autovacuum process
- When running the init command for a postgres database, there are two databases created by default called template0 and template1. They are identical and used as a default template when creating new DBs
- Psql command:
    - Psql -l => List of databases
    - All SQL commands must be semicolon terminated
    - Once connected with psql, you can use \e to write SQL with whatever editor you want
    - You can also use \i file_name.ql to run SQL from a particular file
    - Flags:
        - -d: the database name
        - -U: the username
        - -h: the host
- PGDATA
    - PGDATA directory acts as the disk container that stores all the data of the cluster. Stored here: /usr/lib/postgresql/16/data
    - Note: ls -1 will list all files/directories vertically
    - Files:
        - Postgresql.conf - main configuration file
        - Pg_hba.conf - HBA (host bus adapter) file that provides configuration for DB connections
        - PG_VERSION is text version of the postgres version
        - Postmaster.pid is the PID of the postmaster process
    - Directories:
        - Base - contains all users’ data, including databases, tables, and other objects
        - Global - cluster-wide objects
        - Pg_wal - contains WAL files
        - Pg_stat and pg_stat_tmp are for storage of permanent and temporary statistical information about status and health of cluster
    - Objects in PGDATA are not stored with the actual names of objects, but rather number codes that are translated in Postgres. We can see the mapping of numbers to the actual names using the oid2name utility
        - You can use -d to specify a database to look in
        - -f to specify the filenode/object number
    - Since PostgreSQL only allows up to 1GB of storage for each file, additional files will be attached and named 123.1 for the second file and 123.2, etc.
- Tablespaces
    - A tablespace is a directory that can be outside the PGDATA directory and can also belong to different storage
    - They are mapped to the PGDATA directory through symbolic links
    - View tablespaces with oid2name -s

## 3 - Managing Users and Connections
- Allows for users and user groups to specify permissions
- Role- a collection of database permissions and connection properties
- Roles are defined at the cluster level, permissions are defined at the database level
- SQL for Roles
    - CREATE ROLE, ALTER ROLE, and DROP ROLE
    - CREATE ROLE name [ WITH option ]
    - CREATE ROLE enrico WITH LOGIN PASSWORD ‘xxx’ IN ROLE book_authors;
    - Create new role enrico that is part of the book_authors group
    - DROP ROLE role_name
    - Can do DROP ROLE IF EXIST role_name to prevent getting error messages if it does not exist
    - SELECT current_role;
    - See the current user’s role
    - psql -U ubuntu -d test
    - When connected to psql, you can us “\du” for describe users to list all of the available roles within the system
    - You can also use “\drg” to list all the groups a role is member of
    - Query the pg_roles catalog for user info
    - SELECT rolname, rolcanlogin, roleconnlimit, rolpassword FROM pg_roles WHERE rolname = ’role’;
- Pg_hba.conf File
    - Every line within the file has the following structure
    - <connection-type> <database> <role> <remote-machine> <auth-method>
    - HBA - host based access
    - The first role in the file is used, all others are skipped


