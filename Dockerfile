FROM golang:1.23.6

WORKDIR /usr/src/gomall

ENV GOPOXY=https://goproxy.io,direct

COPY app/frontend/go.mod app/frontend/go.sum ./app/frontend/
COPY rpc_gen rpc_gen
COPY common common

RUN cd app/frontend/ && go mod tidy

COPY app/frontend app/frontend
 
RUN cd app/frontend/ && go build -v -o /opt/gomall/frontend/server

COPY app/frontend/conf /opt/gomall/frontend/conf
COPY app/frontend/static /opt/gomall/frontend/static
COPY app/frontend/template /opt/gomall/frontend/template

WORKDIR /opt/gomall/frontend

EXPOSE 8090

CMD ["./server"]

# docker run -v ./app/frontend/conf:/opt/gomall/frontend/conf --network gomall-study_default --env-file ./app/frontend/.env -p 8090:8090 6fa