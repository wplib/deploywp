ARG VERSION=0.9

# 1. First reference the gearbox-base image.
FROM gearboxworks/gearbox-base:alpine-3.11 as gearbox-base

# 2. Next reference the third party image.
FROM composer:1.10.0
ARG VERSION
ARG GEARBOX_ENTRYPOINT
ARG GEARBOX_ENTRYPOINT_ARGS
USER root

# 3. Copy the build directory over from gearbox-base image.
COPY --from=gearbox-base /etc/gearbox /etc/gearbox

# 4. Set up env variables.
MAINTAINER Gearbox Team <team@gearbox.works>
ENV GEARBOX_CONTAINER_NAME "docker-deploywp"
ENV GEARBOX_CONTAINER_VERSION ${VERSION}
ENV GEARBOX_ENTRYPOINT ${GEARBOX_ENTRYPOINT}
ENV GEARBOX_ENTRYPOINT_ARGS ${GEARBOX_ENTRYPOINT_ARGS}
LABEL gearbox.json='{"schema": "gear-1","meta": {"state": "production","organization": "gearboxworks","name": "deploywp","label": "Composer","maintainer": "Gearbox Team <team@gearbox.works>","class": "dev","refurl": "https://getcomposer.org/"},"build": {"ports": {},"run": "","args": "","env": {"COMPOSER_USER": "gearbox"},"network": "--network gearboxnet","volumes": "","restart": "--restart no"},"run": {"commands": {"default": "deploywp","deploywp": "deploywp"}},"project": {},"extensions": {},"versions": {"1.1.3": {"majorversion": "1.1","latest": false,"ref": "composer:1.1.3","base": "gearboxworks/gearbox-base:alpine-3.4"},"1.2.4": {"majorversion": "1.2","latest": false,"ref": "composer:1.2.4","base": "gearboxworks/gearbox-base:alpine-3.4"},"1.3.3": {"majorversion": "1.3","latest": false,"ref": "composer:1.3.3","base": "gearboxworks/gearbox-base:alpine-3.4"},"1.4.3": {"majorversion": "1.4","latest": false,"ref": "composer:1.4.3","base": "gearboxworks/gearbox-base:alpine-3.7"},"1.5.6": {"majorversion": "1.5","latest": false,"ref": "composer:1.5.6","base": "gearboxworks/gearbox-base:alpine-3.7"},"1.6.5": {"majorversion": "1.6","latest": false,"ref": "composer:1.6.5","base": "gearboxworks/gearbox-base:alpine-3.8"},"1.7.3": {"majorversion": "1.7","latest": false,"ref": "composer:1.7.3","base": "gearboxworks/gearbox-base:alpine-3.10"},"1.8.6": {"majorversion": "1.8","latest": false,"ref": "composer:1.8.6","base": "gearboxworks/gearbox-base:alpine-3.11"},"1.9.3": {"majorversion": "1.9","latest": false,"ref": "composer:1.9.3","base": "gearboxworks/gearbox-base:alpine-3.11"},"0.9": {"majorversion": "1.10","latest": true,"ref": "composer:1.10.0","base": "gearboxworks/gearbox-base:alpine-3.11"}}}'
LABEL gearbox.version='0.9'
ENV GEARBOX_VERSION "0.9"
LABEL container.class="dev"
ENV GEARBOX_CLASS "dev"
LABEL container.label="Composer"
ENV GEARBOX_LABEL "Composer"
LABEL container.maintainer="Gearbox Team <team@gearbox.works>"
ENV GEARBOX_MAINTAINER "Gearbox Team <team@gearbox.works>"
LABEL container.name="deploywp"
ENV GEARBOX_NAME "deploywp"
LABEL container.organization="gearboxworks"
ENV GEARBOX_ORGANIZATION "gearboxworks"
LABEL container.refurl="https://getcomposer.org/"
ENV GEARBOX_REFURL "https://getcomposer.org/"
LABEL container.state="production"
ENV GEARBOX_STATE "production"
LABEL container.args=""
ENV GEARBOX_ARGS ""
# SKIP env
LABEL container.network="--network gearboxnet"
ENV GEARBOX_NETWORK "--network gearboxnet"
LABEL container.ports="map[]"
ENV GEARBOX_PORTS "map[]"
LABEL container.restart="--restart no"
ENV GEARBOX_RESTART "--restart no"
LABEL container.run=""
ENV GEARBOX_RUN ""
LABEL container.volumes=""
ENV GEARBOX_VOLUMES ""
LABEL container.base="gearboxworks/gearbox-base:alpine-3.11"
ENV GEARBOX_BASE "gearboxworks/gearbox-base:alpine-3.11"
LABEL container.latest="true"
ENV GEARBOX_LATEST "true"
LABEL container.majorversion=""
ENV GEARBOX_MAJORVERSION ""
LABEL container.ref="composer:1.10.0"
ENV GEARBOX_REF "composer:1.10.0"
ENV COMPOSER_USER "gearbox"

# 5. Now copy the local files specific to this container.
COPY build /etc/gearbox

# 6. Run the base.sh script to set everything up.
RUN /bin/sh /etc/gearbox/build/base.sh

# 7. Run the composer.sh script to set everything up.
RUN /bin/sh /etc/gearbox/build/composer.sh

# 8. Expose ports.
EXPOSE 22 9970 
WORKDIR /home/gearbox/projects

CMD ["/init"]

# END
################################################################################
