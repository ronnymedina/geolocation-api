FROM golang:1.18.1

ENV workdir /usr/src/app
ENV buildName geoapp

WORKDIR ${workdir}

COPY go.mod go.sum ${workdir}/
RUN go mod download
RUN go mod verify

COPY ./run.sh ${workdir}/
COPY ./src ${workdir}/src

RUN cd src && go build -v -o ${workdir}/${buildName}

CMD ["bash", "run.sh"]
