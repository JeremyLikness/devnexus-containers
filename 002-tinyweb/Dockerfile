FROM busybox:latest
RUN mkdir /www
COPY index.html /www
EXPOSE 80
CMD ["httpd", "-f", "-p", "80", "-h", "/www"]