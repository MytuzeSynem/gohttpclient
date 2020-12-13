FROM golang:1.15
RUN mkdir /app && GO111MODULE=on go get github.com/cucumber/godog/cmd/godog@v0.10.0
ADD . /app
WORKDIR /app/usecase
CMD ["godog", "--random", "-c=1"]