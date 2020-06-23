import datos_pb2

def insertarenBD(nombre, departamento, edad, forma_contagio, estado):
    res = datos_pb2.Respuesta()
    print(nombre+", "+departamento+", "+edad+", "+forma_contagio+", "+estado)
    print("\n---------------------------\n")
    res.Enviado=True
    res.msg=""+nombre+", "+departamento+", "+edad+", "+forma_contagio+", "+estado
    return res
