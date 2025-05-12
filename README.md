# AITUmoment

## Description  
AITUmoment is a social network targeting students of AITU that resembles Reddit. 
Functionality will include:
- Forming individual feed page for users.
- Publishing posts that will form threads.
- Voting functionality.
- Profile functionality.

## How To Run

Required software:
- Golang
- PostgreSQL server
- Docker

1. Clone repository through git / download and unarchive the repository.
2. Make sure you have local postrgeSQL server running.
3. create and fill following env files:
    ```
    # ./api_gateway/.env
    
    HTTP_PORT=8080
    GRPC_CORE_SERVICE_URL=0.0.0.0:5001
    # DEV for development(logs into STDOUT with logs level Debug)
    # PROD for production(logs to file "./logs/aitu_mom.log" relative to the working directory last used in dockerfile)
    APP_MODE=DEV 
    # change to something reliable
    JWT_SECRET=superduper 

    # ./core_service/.env

    PGHOST=
    PGUSER=
    PGPASSWORD=
    PGDATABASE=
    PGSSLMODE= 

    GRPC_PORT=5001
    # see line 25, 26
    APP_MODE=DEV

    NATS_HOSTS="nats:4222"
    # Change NKEY because this one was used in a tutorial
    NATS_NKEY=SUACSSL3UAHUDXKFSNVUZRF5UHPMWZ6BFDTJ7M6USDXIEDNPPQYYYCU3VY
    NATS_EMAIL_VERIFICATION_EVENT_SUBJECT=aitu_moment.email_service.job.send_email_verification

    # change to something reliable
    VERIFICATION_JWT_SECRET=superduper
    # origin should be reachable from the network users are from or email verification links will not work
    VERIFICATION_ORIGIN=http://localhost:8080 

    # ./email_service/.env

    # Currently is not used in other microservices
    GRPC_PORT=5002
    # see line 25, 26
    APP_MODE=DEV

    # Same as lines 42-45
    NATS_HOSTS="nats:4222"
    NATS_NKEY=SUACSSL3UAHUDXKFSNVUZRF5UHPMWZ6BFDTJ7M6USDXIEDNPPQYYYCU3VY
    NATS_EMAIL_VERIFICATION_EVENT_SUBJECT=aitu_moment.email_service.job.send_email_verification

    # Google email and app password(need to turn on two-step verification)
    EMAIL=<email>@gmail.com
    EMAIL_PASSWORD="abcd abcd abcd abcd"
    # Standard for google smtp server
    EMAIL_HOST="smtp.gmail.com"
    EMAIL_HOST_PORT=587
    ```
4. From the project root folder run docker compose: `docker compose up`
5. Open your browser and open the `http://localhost:8080` url address.

## Tools and resources

- Go https://go.dev/doc
- Golang-migrate https://github.com/golang-migrate/migrate/v4
- Gin Web framework https://gin-gonic.com/
- Logrus https://github.com/sirupsen/logrus
- JWT with https://github.com/dgrijalva/jwt-go
- sqlx library https://github.com/jmoiron/sqlx
- Htmx https://htmx.org/
- Tailwindcss https://tailwindcss.com/
- Docker https://docs.docker.com/get-started/
- Go mail https://pkg.go.dev/gopkg.in/mail.v2@v2.3.1
- Protobuf https://protobuf.dev/programming-guides/proto3/
- GRPC https://grpc.io/docs/languages/go/basics/
- NATS https://docs.nats.io/
- Postgresql https://www.postgresql.org/
