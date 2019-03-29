FROM frolvlad/alpine-glibc:glibc-2.29

WORKDIR /bin
ADD release/linux/amd64/api-ten-minutes /bin/
ADD config.yml /bin/

EXPOSE 6868
ENTRYPOINT ["/bin/api-ten-minutes"]
