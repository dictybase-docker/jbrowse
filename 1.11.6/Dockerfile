FROM golang:1.4.2-onbuild
MAINTAINER Siddhartha Basu <siddhartha-basu@northwestern.edu>

ENV JBROWSE_BASE /usr/src
ENV JBROWSE_FOLDER $JBROWSE_BASE/jbrowse
ENV JBROWSE_VERSION 1.11.6
RUN curl -L -o /tmp/jbrowse.tar.gz https://github.com/GMOD/jbrowse/archive/1.11.6-release.tar.gz  && \
    tar xvzf /tmp/jbrowse.tar.gz -C $JBROWSE_BASE && \
    mv $JBROWSE_BASE/jbrowse-$JBROWSE_VERSION-release jbrowse
EXPOSE 9595
ONBUILD COPY jbrowse.conf $JBROWSE_FOLDER
