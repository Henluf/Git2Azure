FROM ubuntu:latest
LABEL authors="Azureman"

ADD azureContainerTest ./azureContainerTest
RUN chmod a+rwx ./azureContainerTest
RUN ls -al

EXPOSE 31415/tcp

CMD "ls -al && ./azureContainerTest"
