FROM    golang:latest as builder
ENV     root /var/go_build
WORKDIR ${root}
ENV     binFile Dlog 
COPY    . ${root}
RUN     go env -w GO111MODULE=on\
        && go env -w GOPROXY=https://goproxy.cn,direct\
        && export CGO_ENABLED=0\
        && go build -o ${binFile} main.go


FROM    alpine
ENV     TZ Asia/ShangHai
WORKDIR /data
COPY    --from=builder /var/go_build/Dlog /data/
COPY    conf.yaml /data/
COPY    cert/ /data/cert/
EXPOSE  9090
CMD     [ "./Dlog" ]