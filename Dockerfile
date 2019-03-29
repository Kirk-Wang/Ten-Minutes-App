FROM frolvlad/alpine-glibc:glibc-2.29

WORKDIR /bin
ADD release/linux/amd64/ten-minutes-app-api /bin/
ADD config.yml /bin/

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "/bin/ten-minutes-app-api", "-ping" ]

EXPOSE 6868
ENTRYPOINT ["/bin/ten-minutes-app-api"]
