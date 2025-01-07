FROM kong:3.9

ENV CUSTOM_PLUGINS_VERSION=0.0.1
ENV KONG_PLUGINS=bundled,auto-request-id

USER root

COPY auto-request-id/kong/plugins/auto-request-id /usr/local/share/lua/5.1/kong/plugins/auto-request-id

USER kong

EXPOSE 8000 8443 8001 8444 8002

CMD ["kong", "docker-start"]