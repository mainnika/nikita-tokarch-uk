# syntax = docker/dockerfile:1.2
FROM registry.access.redhat.com/ubi8/ubi as go-builder

RUN dnf makecache

WORKDIR /usr/src/nikita-tokarch-uk-frontend

RUN dnf install -yq golang

COPY go.mod .
COPY go.sum .
COPY templates/go.mod templates/go.mod
COPY templates/go.sum templates/go.sum

RUN --mount=type=cache,id=gopath,target=${GOPATH} \
    go mod \
      download

ARG APP_VERSION=containerized

COPY cmd cmd
COPY pkg pkg
COPY templates templates

RUN --mount=type=cache,id=gopath,target=${GOPATH} \
    go build \
      -o nikita-tokarch-uk-frontend \
      -ldflags "-X main.Version=${APP_VERSION}" \
      code.tokarch.uk/mainnika/nikita-tokarch-uk/cmd/renderer

FROM registry.access.redhat.com/ubi8/ubi as js-builder

RUN dnf makecache

WORKDIR /usr/src/nikita-tokarch-uk-frontend

RUN dnf install -yq nodejs npm

COPY package-lock.json .
COPY package.json .

RUN npm ci

ARG NODE_ENV=production

COPY webpack.config.js .
COPY web web

RUN npm run build

FROM registry.access.redhat.com/ubi8/ubi as final

RUN dnf makecache
RUN dnf install -yq nginx

RUN mv /usr/share/nginx/html/index.html /usr/share/nginx/html/version.html

COPY nginx /etc/nginx

COPY --from=go-builder \
    /usr/src/nikita-tokarch-uk-frontend/nikita-tokarch-uk-frontend \
    /usr/local/bin/nikita-tokarch-uk-frontend

COPY --from=js-builder \
    /usr/src/nikita-tokarch-uk-frontend/out \
    /usr/share/nginx/html

CMD ["/bin/false"]