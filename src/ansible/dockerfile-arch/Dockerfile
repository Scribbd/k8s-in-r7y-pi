# Dockerfile for an ansible environment for my windows (or any other OS with docker)
# Taken from https://github.com/willhallonline/docker-ansible
# with the following modifications:
# - ARGS for alpine and ansible version specification
# - Have libyaml support (https://www.jeffgeerling.com/blog/2021/ansible-might-be-running-slow-if-libyaml-not-available)
# - Setup enviromnent variables for config
# - Create an ansible user/group and set folders

ARG ALPINE_VERSION=3

FROM alpine:${ALPINE_VERSION}

ARG ANSIBLE_VERSION=4.5.0

ENV ANSIBLE_PIPELINING=True \
    ANSIBLE_GATHERING=smart \
    ANSIBLE_RETRY_FILES_ENABLED=False \
    DEFAULT_ROLES_PATH=\ansible\roles

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
        yaml-dev \
        docker \
        docker-compose && \
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
    pip3 install mitogen molecule molecule-docker ansible-lint yamllint jmespath docker docker-compose openshift && \
    pip3 install --no-cache-dir --no-binary pyyaml ansible==${ANSIBLE_VERSION} && \
    pip3 install --upgrade pywinrm && \
    apk del build-dependencies && \
    rm -rf /var/cache/apk/* && \
    rm -rf /root/.cache/pip && \
    rm -rf /root/.cargo

# Install minikube
WORKDIR /tmp
RUN apk --no-cache add libc6-compat && \
    curl -LO https://dl.k8s.io/release/v1.22.0/bin/linux/amd64/kubectl && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# Setup ansible user
RUN addgroup -g 2000 ansible && \
    adduser -u 2001 -D -H -G ansible ansible && \
    addgroup ansible docker && \
    addgroup ansible wheel

# Setup ansible environment
RUN mkdir /ansible && \
    chown ansible:ansible /ansible && \
    chmod -R 760 /ansible && \
    mkdir -p /etc/ansible

# Prepare for entry
WORKDIR /ansible