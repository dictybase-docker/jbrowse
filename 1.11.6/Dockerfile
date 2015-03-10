FROM golang:1.4.2-onbuild
MAINTAINER Siddhartha Basu <siddhartha-basu@northwestern.edu>

WORKDIR /usr/src
ENV JBROWSE_FOLDER jbrowse
ENV JBROWSE_VERSION 1.11.6
RUN curl -L -o jbrowse.tar.gz https://github.com/GMOD/jbrowse/archive/1.11.6-release.tar.gz  && \
    tar xvzf jbrowse.tar.gz && \
    mv jbrowse-$JBROWSE_VERSION-release jbrowse && \
    rm jbrowse.tar.gz
EXPOSE 9595
ONBUILD COPY jbrowse.conf $JBROWSE_FOLDER/
