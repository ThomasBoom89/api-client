FROM golang:1.22 AS test-backend

RUN useradd -ms /bin/sh user
USER user

WORKDIR /app

COPY --chown=user . .
RUN make install-backend-dependencies

CMD ["make", "test-backend"]


FROM node:20.16.0 AS test-unit-frontend

RUN useradd -ms /bin/sh user
USER user

WORKDIR /app

COPY --chown=user . .
RUN make install-frontend-dependencies

CMD ["make", "test-frontend-unit"]


FROM node:20.16.0 AS test-e2e-frontend
RUN apt update
RUN apt install -y xvfb libgtk-3-dev libwebkit2gtk-4.0-dev

RUN wget https://go.dev/dl/go1.22.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.22.6.linux-amd64.tar.gz

RUN useradd -ms /bin/sh user
USER user

RUN mkdir /home/user/go
ENV PATH=$PATH:/usr/local/go/bin:/home/user/go/bin
ENV DISPLAY=:1
ENV GOPATH=/home/user/go

WORKDIR /app

COPY --chown=user . .

RUN make install
RUN make frontend-build
RUN make install-frontend-playwright

CMD ["make", "test-frontend-e2e-container"]
