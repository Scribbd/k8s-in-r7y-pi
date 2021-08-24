# Dockerfile for an ansible environment for my windows (or any other OS with docker)
# Taken from https://github.com/willhallonline/docker-ansible
# with the following modifications:
# - ARGS for alpine and ansible version specification
# - Have libyaml support (https://www.jeffgeerling.com/blog/2021/ansible-might-be-running-slow-if-libyaml-not-available)
# - Setup enviromnent variables for config
# - Create an ansible user/group and set folders
# - Copy local directory inside container

ARG ALPINE_VERSION=3
ARG ANSIBLE_VERSION=4.4.0

FROM alpine:${ALPINE_VERSION}

ENV ANSIBLE_PIPELINING=True \
    ANSIBLE_GATHERING=smart \
    ANSIBLE_RETRY_FILES_ENABLED=False

# Install all
RUN apk --no-cache add \
        curl \
        sudo \
        python3\
        py3-pip \
        openssl \
        ca-certificates \
        sshpass \
        openssh-client \
        rsync \
        git \
        yaml-dev && \
# yaml-dev is not just a build dependency. It is actively required to have libyaml support enabled
    apk --no-cache add --virtual build-dependencies \
        python3-dev \
        libffi-dev \
        musl-dev \
        gcc \
        cargo \
        openssl-dev \
        libressl-dev \
        build-base && \
    pip3 install --upgrade pip wheel && \
    pip3 install --upgrade cryptography cffi && \
    pip3 install mitogen ansible-lint jmespath && \
    pip3 install --no-cache-dir --no-binary pyyaml ansible==${ANSIBLE_VERSION} && \
    pip3 install --upgrade pywinrm && \
    apk del build-dependencies && \
    rm -rf /var/cache/apk/* && \
    rm -rf /root/.cache/pip && \
    rm -rf /root/.cargo

# Setup ansible user
RUN addgroup -g 2000 ansible && \
    adduser -u 2001 -D -H -G ansible ansible

# Setup ansible environment
RUN mkdir /ansible && \
    chown ansible:ansible /ansible && \
    mkdir -p /etc/ansible 

# Copy in files inside container
COPY --chown=ansible:ansible . /ansible

# Prepare for entry
USER ansible
WORKDIR /ansible