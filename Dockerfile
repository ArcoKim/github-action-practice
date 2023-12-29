FROM public.ecr.aws/docker/library/golang:alpine AS builder

WORKDIR /build

COPY go.mod main.go ./

RUN go build -o main .

FROM public.ecr.aws/docker/library/alpine

WORKDIR /app

COPY --from=builder /build/main .

RUN apk --no-cache add curl

CMD ["./main"]