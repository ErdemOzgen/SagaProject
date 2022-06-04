FROM golang

WORKDIR $GOPATH/src/SagaProject

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o .

EXPOSE 8000

CMD ["sagaAlienInvasion"]