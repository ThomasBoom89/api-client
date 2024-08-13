FROM golang:1.22 AS Test-Backend

RUN useradd -ms /bin/sh user
USER user

WORKDIR /app

COPY --chown=user . .
RUN make install-backend-dependencies

CMD ["make", "test-backend"]


FROM node:20.16.0 AS Test-Unit-Frontend

RUN useradd -ms /bin/sh user
USER user

WORKDIR /app

COPY --chown=user . .
RUN make install-frontend-dependencies

CMD ["make", "test-frontend-unit"]
