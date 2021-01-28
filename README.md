**IDENTIFICADOR DE MUTANTES**

Magneto quiere reclutar la mayor cantidad de mutantes para poder luchar
contra los X-Men.
Sabrás si un humano es mutante, si encuentras ​ más de una secuencia de cuatro letras
iguales​ , de forma oblicua, horizontal o vertical.
Ejemplo (Caso mutante):
String[] dna = {"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"};
En este caso el llamado a la función isMutant(dna) devuelve “true”.

**CLONAR EL PROYECTO**

El proyecto descansa en un repository GitHub el cual puedes clonar desde la siguiente ruta:

**CONFIGURACION DE LA APLICACIÓN**

La aplicación ha sido desarrollada en Golang y corre contra una base datos mongo.
Tanto la aplicación como la base de datos no relacional mongo estan desplegados en una instancia o EC2
Linux/Ubuntu en Amazon Web Service (AWS) bajo una licencia personal Free tier a nombre
de Leonardino Lima.

PASO 1:
Conectar por ssh a la instancia desde una consola linux: 

- Copiamos el archivo "MyKey.pem" en algun lugar de nuestro computador
- Abrimos una consola linux.
- Ejecutamos el siguiente comando para conectarnos por SSH, ssh -i ~/.ssh/certs/MyKey.pem ubuntu@3.138.55.67 
Asumiendo que he guardado
previamiante mi archivo de claves "MyKey.pem" en la ruta .ssh/certs/MyKey.pem de mi equipo local.
  
PASO 2: 
Ejecutar el programa hecho en golang
- Para realizar esto primero debemos aseguramos que estemos ubicados en la ruta ubuntu@ip-172-31-34-69 de nuestra EC2 de AWS
y ejecutamos la siguiente instrucción ./main 
  (No te preocupes, no pasará nada a simple vista pero el servicio ya está listo para ser usado)
  
PASO 3: 
Probando los servicios
- Para probar el servicio /stats (Ya que es un get) simplemente ingresamos a nuestro 
navegador e ingresamos a la ruta http://3.138.55.67:8080/stats y deberíamos 
  obtener un resultado parecido a esto:
  {"CountMutantDna":1,"CountHumanDna":0,"Ratio":0}
  
- Para probar el servicio /mutant (Ya que es un post)
necesitaremos alguna herramienta como POSTMAN para hacerlo.
  
  - URL: http://3.138.55.67:8080/mutant
  - ENCABEZADOS : Content-Type:application/json
  - BODY HUMANO 
    {"dna" : "{ZTGCZA,CAGTGC,TTATGT,AGAAGG,ZCCCTA,TCACTA}"
    }  
  - BODY MUTANTE
    {"dna" : "{ATGCGA,CAGTGC,TTATGT,AGAAGG,CCCCTA,TCACTG}"
    }

**POR ULTIMO**

Si presentas problemas o alguna inquietud no dudes en contactarme al 3177296823 inmediatamente ya que podrías necesitar algun soporte desde la consola AWS

MUCHAS GRACIAS!!


  


  
  
