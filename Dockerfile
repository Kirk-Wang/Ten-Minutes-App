FROM plugins/base:multiarch

LABEL maintainer="lotteryjs <lotter8js@gmail.com>" \
  org.label-schema.name="Ten Minutes App API" \
  org.label-schema.vendor="LOTTERYJS" \
  org.label-schema.schema-version="1.0"

ADD release/linux/amd64/ten-minutes-app-api /bin/
ADD config.yml /bin/

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "/bin/ten-minutes-app-api", "-ping" ]

ENTRYPOINT ["/bin/ten-minutes-app-api"]
