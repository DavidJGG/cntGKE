import grpc
from concurrent import futures
import time

import datos_pb2
import datos_pb2_grpc
import adminDatos

class InsertarService(datos_pb2_grpc.ServicioDatosServicer):
    def obtenerDatos(self, request, context):
        #response = datos_pb2.Respuesta()
        response = adminDatos.insertarenBD(request.nombre, request.departamento, request.edad, request.forma_contagio, request.estado)
        return response

servidor = grpc.server(futures.ThreadPoolExecutor(max_workers=20))

datos_pb2_grpc.add_ServicioDatosServicer_to_server(
    InsertarService(), servidor
)

print('Servidor gRPC en puerto: 50051')
servidor.add_insecure_port('[::]:50051')
servidor.start()

try:
    while True:
        time.sleep(86400)
except KeyboardInterrupt:
    servidor.stop(0)