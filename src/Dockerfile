FROM    golang:1.19 as builder
ENV     root /var/go_build
WORKDIR ${root}
ENV     binFile Dlog 
COPY    . ${root}
RUN     go env -w GO111MODULE=on\
        && go env -w GOPROXY=https://goproxy.cn,direct\
        && export CGO_ENABLED=0\
        && export GOOS=linux\
        && export GOARCH=amd64\
        && go build -o ${binFile} main.go


FROM    alpine
ENV     TZ Asia/ShangHai
WORKDIR /data
ENV     DOCKERIZE_VERSION v0.6.1
RUN     wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
        && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
        && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
COPY    --from=builder /var/go_build/Dlog /data/
COPY    conf.yaml /data/
COPY    cert/ /data/cert/
COPY    utils/ubuntu.ttf /data/utils/
EXPOSE  9090
CMD     [ "./Dlog" ]