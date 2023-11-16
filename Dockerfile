FROM golang:latest
RUN mkdir /stocks
COPY ./stocks/build /stocks
WORKDIR /stocks

EXPOSE 8080

CMD ["/stocks"]