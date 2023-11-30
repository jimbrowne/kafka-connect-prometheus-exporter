FROM golang:alpine as BUILDER

WORKDIR /build

COPY . .
RUN go mod download && go mod tidy


RUN CGO_ENABLED=0 go build

FROM alpine
WORKDIR /app
COPY --from=BUILDER /build/kafka-connect-exporter .

ARG USER=app
ARG UID=12345
ARG GID=12345

RUN addgroup -g "$GID" "$USER" && \
    adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

EXPOSE 8080

USER $USER
CMD ["sh", "-c", "/app/kafka-connect-exporter -telemetry-path \"${TELEMETRY_PATH}\" -scrape-uri \"${SCRAPE_URI}\" -user \"${USERNAME}\" -pass \"${PASSWORD}\""]