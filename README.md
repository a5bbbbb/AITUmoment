# AITUmoment

## Description  
AITUmoment is a social network targeting students of AITU that resembles Reddit. 
Functionality will include:
- Forming individual feed page for users.
- Publishing posts that will form threads.
- Voting functionality.

## Team
- Serafim Bronnikov-Belogurov
- Viktor Suprunov

## How To Run

Required software:
- Golang
- PostgreSQL server

1. Clone repository through git / download and unarchive the repository.
2. Make sure you have local postrgeSQL server running.
3. create .env file in root folder, paste following variables, and fill them with corresponding settings of your postgreSQL server:
    ```
    PGHOST=
    PGUSER=
    PGPASSWORD=
    PGDATABASE=
    PGSSLMODE= 
    ```
    or if your postgreSQL server has only default settings, then you can skip this step.
4. From the root folder run following command: `go run .`
5. Open your browser and open the `http://localhost:8080` url address.

## Tools and resources

- https://go.dev/doc
- https://github.com/golang-migrate/migrate/v4
- Gin Web framework 
- sqlx library https://github.com/jmoiron/sqlx

