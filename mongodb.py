import pymongo
from pymongo import MongoClient
import redis


print("\n---------------- MONGO ------------------\n")
con=MongoClient('35.238.79.144',27017)
try:
    db=con.proyecto2

    #db.casos.delete_many({"name":"LUCHO"})
    print("Elementos: "+str(db.casos.count()))
    for x in db.casos.find():
        print(x)
except Exception as e:
    print(e)
finally:
    con.close()



#db.tabla1.insert({"name3":"tutorials point3"})
#db.tabla1.insert({"name4":"tutorials point4"})
#db.tabla1.insert({"name5":"tutorials point5"})
#db.tabla1.insert({"name6":"tutorials point6"})
#db.tabla1.insert({"name7":"tutorials point7"})
#db.tabla1.insert({"name8":"tutorials point8"})



#r = redis.StrictRedis(host='34.70.196.45', port=6379, db=0)
#r.hset("llave", "campo1", "valor 1")

print("---------------- REDIS ------------------\n")

r = redis.Redis(host='35.238.79.144', port=6379)
try:
 
    #r.delete("s")
    #r.rpush("proyecto2",'{"Nombre" : "Prueba22", "Departamento" : "San marcos", "Edad" : 37, "Forma de contagio" : "Comunitario", "Estado" : "Muerto"}')
    print("Elementos: "+str(r.llen("proyecto2")))
    rango=r.lrange('proyecto2', 0, -1)
    for x in rango:
        print(x)
except Exception as e:
    print(e)
finally:
    r.close()
