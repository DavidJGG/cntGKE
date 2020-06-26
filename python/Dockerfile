FROM davidjgg/server-grpc-python:normal

MAINTAINER 201610648
RUN python3.6 -m pip install pymongo
RUN python3.6 -m pip install redis
RUN rm -R /servPython
RUN mkdir /servPython
COPY . /servPython
CMD ["python3.6","/servPython/servidor.py"]
