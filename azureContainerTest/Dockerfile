FROM ubuntu:latest
LABEL authors="Azureman"

ADD azureContainerTest ./azureContainerTest
RUN chmod a+rwx ./azureContainerTest

EXPOSE 31415/tcp

CMD "./azureContainerTest"
