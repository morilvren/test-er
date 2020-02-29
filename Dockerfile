FROM registry.icp.com:5000/service/devops/runtime/golang:4.5.0
WORKDIR /usr/bin
RUN mkdir -p /etc/config/
ADD ./hello /usr/bin
RUN chmod -R a+rwx /usr/bin/hello
EXPOSE 80
EXPOSE 90
EXPOSE 8090
ENTRYPOINT /usr/bin/hello
