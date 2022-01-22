FROM  golang:1.17.5-bullseye

RUN go install github.com/cosmtrek/air@v1.27.8 && \
    go install github.com/rubenv/sql-migrate/sql-migrate@latest
