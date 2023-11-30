FROM BUILD_IMG as builder

WORKDIR /go/src/PROJECT_NAME

COPY . .
RUN make build


FROM BASE_IMG

WORKDIR /app

COPY --from=builder /go/src/PROJECT_NAME/bin/NAME /app/NAME

CMD ["/app/NAME"]
