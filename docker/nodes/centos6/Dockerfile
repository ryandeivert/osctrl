FROM centos:6
LABEL maintainer="javuto"

RUN curl -# "https://osquery-packages.s3.amazonaws.com/rpm/osquery-3.3.2-1.linux.x86_64.rpm" -o "/tmp/osquery.rpm"
RUN rpm -ivh "/tmp/osquery.rpm"

COPY docker/nodes/centos6/wait.sh .

CMD [ "/bin/sh", "wait.sh" ]
