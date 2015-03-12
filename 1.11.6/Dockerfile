FROM golang:1.4.2-onbuild
MAINTAINER Siddhartha Basu <siddhartha-basu@northwestern.edu>

WORKDIR /usr/src
ENV JBROWSE_FOLDER jbrowse
ENV JBROWSE_VERSION 1.11.6
RUN apt-get update && apt-get -y install unzip \
    && rm -rf /var/lib/apt/lists/*
RUN curl -L -o jbrowse.zip http://jbrowse.org/wordpress/wp-content/plugins/download-monitor/download.php?id=99 \
    && unzip jbrowse.zip \
    && mv JBrowse-$JBROWSE_VERSION jbrowse \
    && rm jbrowse.zip
EXPOSE 9595
ONBUILD COPY jbrowse.conf $JBROWSE_FOLDER/
ONBUILD COPY config $JBROWSE_FOLDER/config/
