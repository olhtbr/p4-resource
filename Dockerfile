FROM quay.io/justcontainers/base:v0.12.0

ADD bin/in /opt/resource/in
ADD bin/out /opt/resource/out
ADD bin/check /opt/resource/check
ADD http://ftp.perforce.com/pub/perforce/r15.2/bin.linux26x86_64/p4 /usr/bin/p4

RUN chmod +x /usr/bin/p4 && \
    chmod +x /opt/resource/in && \
    chmod +x /opt/resource/out && \
    chmod +x /opt/resource/check
