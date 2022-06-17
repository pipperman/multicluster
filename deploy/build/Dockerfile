FROM golang:1.16 AS builder

ARG APP_RELATIVE_PATH

COPY . /src/${APP_RELATIVE_PATH}
WORKDIR /src/${APP_RELATIVE_PATH}

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

ARG APP_RELATIVE_PATH

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        procps \
        iputils \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/${APP_RELATIVE_PATH}/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./multicluster", "-conf", "/data/conf"]