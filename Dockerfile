FROM  golang:1.17.5-bullseye

RUN go install github.com/cosmtrek/air@v1.27.8
