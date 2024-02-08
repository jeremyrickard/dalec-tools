FROM mcr.microsoft.com/oss/go/microsoft/golang:1.21-cbl-mariner2.0 as BUILDER
COPY . /dalec-tools
WORKDIR /dalec-tools
RUN go build -o bin/dalec-tools .

FROM mcr.microsoft.com/cbl-mariner/distroless/base:2.0

COPY --from=BUILDER /dalec-tools/bin/dalec-tools .
ENTRYPOINT [ "/dalec-tools" ]