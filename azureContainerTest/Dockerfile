FROM ubuntu:latest
LABEL authors="Azureman"

ADD azTest /azTest
RUN chmod a+rwx azTest/azTest

EXPOSE 31415/tcp

CMD ["ls","-al"]
