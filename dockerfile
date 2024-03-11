FROM scratch

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/HuaTug/Work-3
COPY . $GOPATH/src/github.com/HuaTug/Work-3
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./4-ToDoList"]