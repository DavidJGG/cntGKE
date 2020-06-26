import datos_pb2
import pymongo
from pymongo import MongoClient
import math
import redis

def insertarenBD(nombre, departamento, edad, forma_contagio, estado):
    res = datos_pb2.Respuesta()

    #---------------------- MONGO -------------------------------
    con=MongoClient('34.70.196.45',27017)
    res.Enviado=True
    aux=""
    try:
        db=con.proyecto2
        db.casos.insert_one({"name" : nombre, "depto" : departamento, "age" : int(edad), "form" : forma_contagio, "state" : estado})        
    except Exception as e:
        print(e)
        aux=", NO SE PUDO CONECTAR CON MONGO "+str(e)
    finally:
        con.close()
        
#---------------------- REDIS -------------------------------

    r = redis.Redis(host='34.70.196.45', port=6379)
    try:
        r.rpush("proyecto2",'{"name" : "'+nombre+'", "depto" : "'+departamento+'", "age" : '+edad+', "form" : "'+forma_contagio+'", "state" : "'+estado+'"}')
    except Exception as e:
        print(e)        
        aux=", NO SE PUDO CONECTAR CON REDIS "+str(e)
    finally:
        r.close()

    res.msg=""+nombre+", "+departamento+", "+edad+", "+forma_contagio+", "+estado+" "+aux
    return res
