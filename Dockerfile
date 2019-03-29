FROM frolvlad/alpine-glibc:glibc-2.29

WORKDIR /bin
ADD release/linux/amd64/api-ten-minutes /bin/
ADD config.yml /bin/

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "/bin/api-ten-minutes", "-ping" ]

EXPOSE 6868
ENTRYPOINT ["/bin/api-ten-minutes"]
