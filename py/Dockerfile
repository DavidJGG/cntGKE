FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install -y python3
RUN apt-get install -y python3-pip
RUN python3.6 -m pip install grpcio

MAINTAINER 201610648
RUN mkdir -p /servPython
COPY . /servPython
CMD ["python3.6","/servPython/servidor.py"]
