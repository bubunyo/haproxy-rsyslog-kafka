FROM centos:8

RUN cd /etc/yum.repos.d/
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*

RUN yum -y update \
    && yum install -y rsyslog rsyslog-kafka \
    && yum clean all

# Create directory for extra confguration files and logs.
RUN mkdir -p /etc/rsyslog.d /srv/rsyslog

# Copy configuration files.
COPY ./rsyslog.conf /etc/
COPY ./kafka.conf /etc/rsyslog.d/

# Run rsyslogd in foreground with custom PID file path, since it's running under unprivileged user.
ENTRYPOINT [ "sh", "-c", "/usr/sbin/rsyslogd -n -i /srv/rsyslog/rsyslog.pid" ]

EXPOSE 514 514/udp