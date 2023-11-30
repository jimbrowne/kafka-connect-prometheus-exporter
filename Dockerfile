FROM --platform=$BUILDPLATFORM golang:alpine as BUILDER
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download && go mod tidy

ADD cmd cmd
ADD pkg pkg

RUN CGO_ENABLED=0 go build

FROM alpine
WORKDIR /app
COPY --from=BUILDER /app/kafka-connect-exporter .
RUN groupadd -r -g 1001 app \
    && useradd -r -u 1001 -g app app\
    && chown -R app:app /app

EXPOSE 8080

USER app
CMD ["/app/kafka-connect-exporter", \
     "-telemetry-path", "${TELEMETRY_PATH}", \
     "-scrape-uri", "${SCRAPE_URI}", \
     "-user", "${USERNAME}", \
     "-pass", "${PASSWORD}" \
]
