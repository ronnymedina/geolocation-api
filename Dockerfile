FROM golang:1.18.1

ENV workdir /usr/src/app

WORKDIR ${workdir}

COPY go.mod go.sum ./run.sh ${workdir}/
RUN go mod download && go mod verify

COPY ./src ${workdir}/src

# RUN go build -v -o /usr/local/bin/app

CMD ["bash", "run.sh"]
