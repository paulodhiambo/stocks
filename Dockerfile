FROM golang:latest
RUN mkdir /stocks
COPY ./stocks/build /stocks
WORKDIR /stocks
RUN /stocks/stocks