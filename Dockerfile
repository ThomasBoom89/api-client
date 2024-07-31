FROM golang:1.22 AS Test-Backend

RUN useradd -ms /bin/sh user
USER user

WORKDIR /app

COPY --chown=user . .
RUN make install-backend-dependencies

CMD ["make", "test-backend"]
