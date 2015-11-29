FROM daocloud.io/library/golang
MAINTAINER xdlinear

ENV PORT 8088

COPY src /src
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8088

ENTRYPOINT ["/entrypoint.sh"]


