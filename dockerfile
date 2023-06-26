FROM golang:latest


WORKDIR /revx

COPY . .

RUN go install github.com/revx-official/revxbuildtool@latest

RUN go mod download
RUN revxbuildtool --release

EXPOSE 80
ENTRYPOINT [ "./bin/revxdaemon" ]