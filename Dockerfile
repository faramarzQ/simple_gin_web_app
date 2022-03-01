FROM golang:1.17-alpine as base

FROM base as dev

RUN mkdir /app
WORKDIR /api
COPY . .

CMD ["go", "run", "."]