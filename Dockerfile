FROM alpine:3.14

COPY moonlogs /moonlogs/moonlogs

WORKDIR "/moonlogs"
ENTRYPOINT ["./moonlogs"]
CMD []
