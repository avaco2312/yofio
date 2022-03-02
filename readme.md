### YoFio - Backend Golang - Prueba técnica

Dividimos los retos de la prueba en:

- [Algoritmo de solución](#algoritmo-de-solución)
- [Nivel básico](#nivel-básico)
- [Nivel intermedio](#nivel-intermedio)
- [Nivel avanzado](#nivel-avanzado)
- [Testing](#testing)

#### Algoritmo de solución

El problema propuesto implica la solución de la ecuación diofántica en tres variables

```
300*x + 500*y + 700*z = investment
```

Esta ecuación tiene 0 o infinitas soluciones sobre los números enteros. El criterio de solubilidad es que **investment** sea divisible por el máximo común divisor de 300, 500 y 700 que es 100.

Luego, si **investment** no es divisible por 100 no hay solución. Aunque en el planteamiento nos garantizan que es múltiplo de 100, en el código lo chequeamos. También que investment sea positivo.

Para simplificar la resolución trabajamos con la ecuación simplificada

```
3*x + 5*y + 7*z = f donde f = investment/100
```

Para resolver la ecuación el método es descomponerla en dos ecuaciones diofánticas de dos variables, de las cuales hay que determinar una solución particular. La solución no la detallamos aquí. Nos ofrece las fórmulas para generar los valores enteros de x, y, z para las infinitas soluciones enteras.

Las fórmulas son:

```
x = 2*f - 14*v -5*u
y = -f + 7*v + 3*u
z = v

donde u, v son variales enteras libres
```
Dada la naturaleza del problema x, y, z tienen que ser no negativas. Aplicando esta restricción, podemos acotar las v que producen este resultado:

```
0 <= v <= f/7

Nota: f/7 se trunca al entero menor
````

Y para cada valor de v, el límite permisible de u:

```
(f-7*v)/3 <= u <= (2*f-14*v)/5

Nota: (f-7*v)/3 se aproxima al entero mayor y (2*f-14*v)/5 se trunca al entero menor
```

Aunque no lo exige la prueba técnica, hacemos un programa que dado un conjunto de prueba de **investment**, calcule TODAS las posibles distribuciones, si existen. El código en el directorio **diofanto/main.go**. Para correrlo, ubicado en ese directorio **go run .**

Los datos de prueba:

```
{-12000, 300, 400, 500, 510, 1000, 1050, 3000, 4000, 7000, 12000}
```
Los resultados:

```
-12000  error: no es positivo o no es múltiplo de 100
300       1 * 300 +   0 * 500 +   0 * 700 =    300   
400     no tiene distribución válida
500       0 * 300 +   1 * 500 +   0 * 700 =    500   
510     error: no es positivo o no es múltiplo de 100
1000      0 * 300 +   2 * 500 +   0 * 700 =   1000   
          1 * 300 +   0 * 500 +   1 * 700 =   1000   
1050    error: no es positivo o no es múltiplo de 100
3000     10 * 300 +   0 * 500 +   0 * 700 =   3000   
          5 * 300 +   3 * 500 +   0 * 700 =   3000   
          0 * 300 +   6 * 500 +   0 * 700 =   3000   
          6 * 300 +   1 * 500 +   1 * 700 =   3000
          1 * 300 +   4 * 500 +   1 * 700 =   3000
          2 * 300 +   2 * 500 +   2 * 700 =   3000
          3 * 300 +   0 * 500 +   3 * 700 =   3000
4000     10 * 300 +   2 * 500 +   0 * 700 =   4000
          5 * 300 +   5 * 500 +   0 * 700 =   4000
          0 * 300 +   8 * 500 +   0 * 700 =   4000
         11 * 300 +   0 * 500 +   1 * 700 =   4000
          6 * 300 +   3 * 500 +   1 * 700 =   4000
          1 * 300 +   6 * 500 +   1 * 700 =   4000
          7 * 300 +   1 * 500 +   2 * 700 =   4000
          2 * 300 +   4 * 500 +   2 * 700 =   4000
          3 * 300 +   2 * 500 +   3 * 700 =   4000
          4 * 300 +   0 * 500 +   4 * 700 =   4000
          0 * 300 +   1 * 500 +   5 * 700 =   4000
7000     20 * 300 +   2 * 500 +   0 * 700 =   7000
         15 * 300 +   5 * 500 +   0 * 700 =   7000
         10 * 300 +   8 * 500 +   0 * 700 =   7000
          5 * 300 +  11 * 500 +   0 * 700 =   7000
          0 * 300 +  14 * 500 +   0 * 700 =   7000
         21 * 300 +   0 * 500 +   1 * 700 =   7000
         16 * 300 +   3 * 500 +   1 * 700 =   7000
         11 * 300 +   6 * 500 +   1 * 700 =   7000
          6 * 300 +   9 * 500 +   1 * 700 =   7000
          1 * 300 +  12 * 500 +   1 * 700 =   7000
         17 * 300 +   1 * 500 +   2 * 700 =   7000
         12 * 300 +   4 * 500 +   2 * 700 =   7000
          7 * 300 +   7 * 500 +   2 * 700 =   7000
          2 * 300 +  10 * 500 +   2 * 700 =   7000
         13 * 300 +   2 * 500 +   3 * 700 =   7000
          8 * 300 +   5 * 500 +   3 * 700 =   7000
          3 * 300 +   8 * 500 +   3 * 700 =   7000
         14 * 300 +   0 * 500 +   4 * 700 =   7000
          9 * 300 +   3 * 500 +   4 * 700 =   7000
          4 * 300 +   6 * 500 +   4 * 700 =   7000
         10 * 300 +   1 * 500 +   5 * 700 =   7000
          5 * 300 +   4 * 500 +   5 * 700 =   7000
          0 * 300 +   7 * 500 +   5 * 700 =   7000
          6 * 300 +   2 * 500 +   6 * 700 =   7000
          1 * 300 +   5 * 500 +   6 * 700 =   7000
          7 * 300 +   0 * 500 +   7 * 700 =   7000
          2 * 300 +   3 * 500 +   7 * 700 =   7000
          3 * 300 +   1 * 500 +   8 * 700 =   7000
          0 * 300 +   0 * 500 +  10 * 700 =   7000
12000    40 * 300 +   0 * 500 +   0 * 700 =  12000
         35 * 300 +   3 * 500 +   0 * 700 =  12000
         30 * 300 +   6 * 500 +   0 * 700 =  12000
         25 * 300 +   9 * 500 +   0 * 700 =  12000
         20 * 300 +  12 * 500 +   0 * 700 =  12000
         15 * 300 +  15 * 500 +   0 * 700 =  12000
         10 * 300 +  18 * 500 +   0 * 700 =  12000
          5 * 300 +  21 * 500 +   0 * 700 =  12000
          0 * 300 +  24 * 500 +   0 * 700 =  12000
         36 * 300 +   1 * 500 +   1 * 700 =  12000
         31 * 300 +   4 * 500 +   1 * 700 =  12000
         26 * 300 +   7 * 500 +   1 * 700 =  12000
         21 * 300 +  10 * 500 +   1 * 700 =  12000
         16 * 300 +  13 * 500 +   1 * 700 =  12000
         11 * 300 +  16 * 500 +   1 * 700 =  12000
          6 * 300 +  19 * 500 +   1 * 700 =  12000
          1 * 300 +  22 * 500 +   1 * 700 =  12000
         32 * 300 +   2 * 500 +   2 * 700 =  12000
         27 * 300 +   5 * 500 +   2 * 700 =  12000
         22 * 300 +   8 * 500 +   2 * 700 =  12000
         17 * 300 +  11 * 500 +   2 * 700 =  12000
         12 * 300 +  14 * 500 +   2 * 700 =  12000
          7 * 300 +  17 * 500 +   2 * 700 =  12000
          2 * 300 +  20 * 500 +   2 * 700 =  12000
         33 * 300 +   0 * 500 +   3 * 700 =  12000
         28 * 300 +   3 * 500 +   3 * 700 =  12000
         23 * 300 +   6 * 500 +   3 * 700 =  12000
         18 * 300 +   9 * 500 +   3 * 700 =  12000
         13 * 300 +  12 * 500 +   3 * 700 =  12000
          8 * 300 +  15 * 500 +   3 * 700 =  12000
          3 * 300 +  18 * 500 +   3 * 700 =  12000
         29 * 300 +   1 * 500 +   4 * 700 =  12000
         24 * 300 +   4 * 500 +   4 * 700 =  12000
         19 * 300 +   7 * 500 +   4 * 700 =  12000
         14 * 300 +  10 * 500 +   4 * 700 =  12000
          9 * 300 +  13 * 500 +   4 * 700 =  12000
          4 * 300 +  16 * 500 +   4 * 700 =  12000
         25 * 300 +   2 * 500 +   5 * 700 =  12000
         20 * 300 +   5 * 500 +   5 * 700 =  12000
         15 * 300 +   8 * 500 +   5 * 700 =  12000
         10 * 300 +  11 * 500 +   5 * 700 =  12000
          5 * 300 +  14 * 500 +   5 * 700 =  12000
          0 * 300 +  17 * 500 +   5 * 700 =  12000
         26 * 300 +   0 * 500 +   6 * 700 =  12000
         21 * 300 +   3 * 500 +   6 * 700 =  12000
         16 * 300 +   6 * 500 +   6 * 700 =  12000
         11 * 300 +   9 * 500 +   6 * 700 =  12000
          6 * 300 +  12 * 500 +   6 * 700 =  12000
          1 * 300 +  15 * 500 +   6 * 700 =  12000
         22 * 300 +   1 * 500 +   7 * 700 =  12000
         17 * 300 +   4 * 500 +   7 * 700 =  12000
         12 * 300 +   7 * 500 +   7 * 700 =  12000
          7 * 300 +  10 * 500 +   7 * 700 =  12000
          2 * 300 +  13 * 500 +   7 * 700 =  12000
         18 * 300 +   2 * 500 +   8 * 700 =  12000
         13 * 300 +   5 * 500 +   8 * 700 =  12000
          8 * 300 +   8 * 500 +   8 * 700 =  12000
          3 * 300 +  11 * 500 +   8 * 700 =  12000
         19 * 300 +   0 * 500 +   9 * 700 =  12000
         14 * 300 +   3 * 500 +   9 * 700 =  12000
          9 * 300 +   6 * 500 +   9 * 700 =  12000
          4 * 300 +   9 * 500 +   9 * 700 =  12000
         15 * 300 +   1 * 500 +  10 * 700 =  12000
         10 * 300 +   4 * 500 +  10 * 700 =  12000
          5 * 300 +   7 * 500 +  10 * 700 =  12000
          0 * 300 +  10 * 500 +  10 * 700 =  12000
         11 * 300 +   2 * 500 +  11 * 700 =  12000
          6 * 300 +   5 * 500 +  11 * 700 =  12000
          1 * 300 +   8 * 500 +  11 * 700 =  12000
         12 * 300 +   0 * 500 +  12 * 700 =  12000
          7 * 300 +   3 * 500 +  12 * 700 =  12000
          2 * 300 +   6 * 500 +  12 * 700 =  12000
          8 * 300 +   1 * 500 +  13 * 700 =  12000
          3 * 300 +   4 * 500 +  13 * 700 =  12000
          4 * 300 +   2 * 500 +  14 * 700 =  12000
          5 * 300 +   0 * 500 +  15 * 700 =  12000
          0 * 300 +   3 * 500 +  15 * 700 =  12000
          1 * 300 +   1 * 500 +  16 * 700 =  12000
```

Pasemos ahora a los retos de la prueba técnica.

#### Nivel básico

Creamos un package **asigna** (código en **asigna/asigna.go**). En él creamos la interface **CreditAssigner** y además un tipo struct **AsignaCredito** que cumple esta interface. Su método **Assign** es el mismo algoritmo usado anteriormente con la diferencia de que retorna al encontrar la primera distribución posible. De no haberla, o si el valor de **investment** es no positivo o no divisible por 100, retorna error. Claro, esto favorece la distribución que primero aparece. En un caso real pudieran implementarse otros criterios de asignación.

Para hacer el programa más general, usamos la interfase  **CreditAssigner** de la que obtendremos una instancia mediante la función **NewCreditAssigner**. Por ahora lo que hace es crear una instancia del tipo "AsignaCredito" que responde al algoritmo estudiado. Pero en el futuro podemos crear otros tipos que respondan a otras formas de distribuir los créditos.

Para este caso básico creamos un programa principal en el directorio **basico/main.go**. Para correrlo, ubicados en ese directorio teclear **go run .**

Lo corremos con los mismos datos de prueba y el resultado es:

```
-12000: error: -12000 no es positivo o no es múltiplo de 100
  1 * 300 +   0 * 500 +   0 * 700 =    300
400: error: 400 no es distribuible
  0 * 300 +   1 * 500 +   0 * 700 =    500
510: error: 510 no es positivo o no es múltiplo de 100
  0 * 300 +   2 * 500 +   0 * 700 =   1000
1050: error: 1050 no es positivo o no es múltiplo de 100
 10 * 300 +   0 * 500 +   0 * 700 =   3000
 10 * 300 +   2 * 500 +   0 * 700 =   4000
 20 * 300 +   2 * 500 +   0 * 700 =   7000
 40 * 300 +   0 * 500 +   0 * 700 =  12000
```

#### Nivel intermedio

Creamos la API REST pedida. Esta puede ser ejecutada en un servidor local o en la nube (por ejemplo, sobre una máquina virtual EC2 en AWS). Utiliza el package **asigna** y su código está en **intermedio/main.go** y en **intermedio/handler/handler.go**. El programa principal utiliza el package **gorilla/mux** para delimitar los métodos permitidos. Hemos separado el handler http en un package, **handler**, para mayor precisión del coverage en el test que desarrollaremos más adelante.

Hagamos unas pruebas con el servidor local y llamadas curl. En el directorio **intermedio** teclear **go run .** para iniciar el servidor local:

```
curl -X POST http://localhost/credit-assignment
-d '{ "investment": 3000 }'

HTTP/1.1 200 OK
Date: Wed, 02 Mar 2022 04:24:40 GMT
Content-Length: 67
Content-Type: text/plain; charset=utf-8
Connection: close

{
  "credit_type_300": 10,
  "credit_type_500": 0,
  "credit_type_700": 0
}
-----
curl -X POST http://localhost/credit-assignment
-d '{ "investment": 7000 }'

HTTP/1.1 200 OK
Date: Wed, 02 Mar 2022 04:25:21 GMT
Content-Length: 67
Content-Type: text/plain; charset=utf-8
Connection: close

{
  "credit_type_300": 20,
  "credit_type_500": 2,
  "credit_type_700": 0
}
-----
curl -X POST http://localhost/credit-assignment
-d '{ "investment": 400 }'

HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 02 Mar 2022 04:25:55 GMT
Content-Length: 30
Connection: close

error: 400 no es distribuible
-----
curl -X POST http://localhost/credit-assignment
-d '{ "investment": 10050 }'

HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 02 Mar 2022 04:27:34 GMT
Content-Length: 53
Connection: close

error: 10050 no es positivo o no es múltiplo de 100
```

Para colocar el servidor en la nube AWS (una forma, después implementaremos otra) creamos una máquina virtual EC2. Copiamos el código compilado del servidor REST al directorio **home** de esa máquina y hacemos que se ejecute al iniciar la máquina virtual mediante **cron**. Damos permiso para que la máquina EC2 escuche en el puerto 80 desde cualquier origen. Las llamadas en este caso serían:

```
curl -X POST http://ec2-52-38-103-53.us-west-2.compute.amazonaws.com/credit-assignment
-d '{ "investment": 3000 }'

HTTP/1.1 200 OK
Date: Wed, 02 Mar 2022 04:28:19 GMT
Content-Length: 67
Content-Type: text/plain; charset=utf-8
Connection: close

{
  "credit_type_300": 10,
  "credit_type_500": 0,
  "credit_type_700": 0
}
-----
curl -X POST http://ec2-52-38-103-53.us-west-2.compute.amazonaws.com/credit-assignment
-d '{ "investment": 400 }'

HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 02 Mar 2022 04:30:23 GMT
Content-Length: 30
Connection: close

error: 400 no es distribuible
```

#### Nivel Avanzado

En el nivel avanzado, para tener más variedad, haremos un deployment diferente sobre AWS. Esta vez emplearemos sevicios **serverless**. Las peticiones de la API REST se gestionarán mediante API Gateway, la que que llamará a funciones Lambda. Para almacenar los pedidos y las estadísticas usaremos como base de datos DynamoDB.

Primero la estructura de la base de datos. Como es común en DynamoDB usaremos la técnica de "una sola tabla". La estructura de la tabla, que llamaremos **yofio**, es:

Id ("partition key"): string. Contiene la Id que se asigna a cada petición, la generamos usando la librería **ksuid**. Como caso especial, el registro con Id "*" contendrá las estadísticas. No usaremos "sort key".

Los registros son:

Id = "*": estadísticas
- asignaciones exitosas (cantidad, importe) (de inicio 0, 0)
- asignaciones no exitosas (cantidad, importe) (de inicio 0, 0)

Id = "xxx..." (generada por ksuid): datos de una asignación
- x, y, z si es una asignación exitosa
- si es una asignación no exitosa, registramos el investment en el campo importe de asignaciones no exitosas.

Por ejemplo, para las peticiones de prueba anteriores, el registro en la tabla quedaría:

| Id                          | AExCan | AExImp | ANoExCan | ANoExImp | x   | y   | z   |
| --------------------------- | ------ | ------ | -------- | -------- | --- | --- | --- |
| *                           | 2      | 10000  | 1        | 400      |     |     | -   |
| 24tfoCSE8Rp5Kj4OTQ6jsBWo3Ga |        |        |          |          | 10  | 0   | 0   |
| 24vygIpnaXhfaoOnuz6OnmwNDDD |        |        |          |          | 20  | 2   | 0   |
| 24LISRUZ5ifj744Oe28S2639WUD |        |        |          | 400      |     |     |     |

Para crear e inicializar la tabla creamos un programa ubicado en **avanzado/local/main.go**. Ubicados en ese directorio el comando **go run -c** crea la tabla y escribe el record de estadísticas. Este programa utiliza el package **initdb** ubicado en **avanzado/initdb**. También puede correrse **go run -d** para eliminar la tabla.

Las peticiones a la API REST las haremos a través de API Gateway. Creamos una API llamada **yofio**. La definición exportada de la API se puede ver en **avanzado/apigateway**. Las llamadas serán tres:

- POST /credit-assignment con body { "investment": xxx }.

- GET /credit-assignment/id: recupera la asignación con la id indicada.

- GET /statistics: recupera las estadísticas requeridas.

Cada entrada a la API llama una función lambda, respectivamente:

- postInvestment

- getInvestment

- getStatistics

El código en **avanzado/lambdas**. 

El enlace de API Gateway con las funciones lambda se define siguiendo la convención **lambda proxy** y se autoriza a cada entrada de la API a utilizar la correspondiente función lambda. También hay que autorizar cada función lambda a ejecutar las operaciones DynamoDB que requiera.

Sólo comentaremos del código en el caso de la lambda **postInvestment**. Determina, usando el algoritmo, si el monto es distribuible. De serlo, escribe un registro en la base de datos con la distribución. Si no lo es, registra el monto. En ambos casos se asigna una Id y se actualiza el registro de estadísticas.

La escritura del registro de la asignación nueva y la actualización del registro de estadísticas se realizan mediante la operación de DynamoDB **TransactWriteItems** que garantiza que ambas operaciones se realicen de forma consistente (ambas se hacen ó ambas fallan).

Una vez implementadas la API, nos da una URL para las llamadas, que es:

```
https://6xiekuxzq2.execute-api.us-west-2.amazonaws.com/produccion
```

Podemos probar las llamadas:

```
curl -X POST https://6xiekuxzq2.execute-api.us-west-2.amazonaws.com/produccion/credit-assignment
-d '{ "investment": 10000 }'

HTTP/1.1 200 OK
Date: Wed, 02 Mar 2022 04:54:47 GMT
Content-Type: application/json
Content-Length: 104
Connection: close
x-amzn-RequestId: 2eebe090-afcb-41ee-b2b0-e34a92857546
x-amz-apigw-id: OVvHmEupPHcFTog=
X-Amzn-Trace-Id: Root=1-621ef896-7feba57e04039a462f6cfbbf;Sampled=0

{
  "id": "25ocTaPAf7YENuQEc6wBmaTVxrO",
  "credit_type_300": 30,
  "credit_type_500": 2,
  "credit_type_700": 0
}
-----
curl -X GET https://6xiekuxzq2.execute-api.us-west-2.amazonaws.com/produccion/credit-assignment/25ocTaPAf7YENuQEc6wBmaTVxrO

HTTP/1.1 200 OK
Date: Wed, 02 Mar 2022 04:56:10 GMT
Content-Type: application/json
Content-Length: 104
Connection: close
x-amzn-RequestId: 43b777fa-bfd7-4a36-b961-743bd41d6963
x-amz-apigw-id: OVvUnG6QPHcF0AA=
X-Amzn-Trace-Id: Root=1-621ef8ea-611d7a9b66c85b9155b2f582;Sampled=0

{
  "id": "25ocTaPAf7YENuQEc6wBmaTVxrO",
  "credit_type_300": 30,
  "credit_type_500": 2,
  "credit_type_700": 0
}
-----
curl -X POST https://6xiekuxzq2.execute-api.us-west-2.amazonaws.com/produccion/credit-assignment
-d '{ "investment": 200 }'

HTTP/1.1 400 Bad Request
Date: Wed, 02 Mar 2022 04:56:57 GMT
Content-Type: application/json
Content-Length: 79
Connection: close
x-amzn-RequestId: fdf59167-bea8-4dbc-8752-eeff1770eab1
x-amz-apigw-id: OVvb_EQavHcFZtA=
X-Amzn-Trace-Id: Root=1-621ef919-44d05b62165ee7595b31bcef;Sampled=0

{
  "id": "25ocjqpoUESIuM7Ia8PCGYZwK8y",
  "error": "error: 200 no es distribuible"
}
-----
curl -X GET https://6xiekuxzq2.execute-api.us-west-2.amazonaws.com/produccion/statistics

HTTP/1.1 200 OK
Date: Wed, 02 Mar 2022 04:57:41 GMT
Content-Type: application/json
Content-Length: 172
Connection: close
x-amzn-RequestId: e08b794e-244a-48bc-a6ea-eb65af4f267a
x-amz-apigw-id: OVvi5EZlPHcF0zw=
X-Amzn-Trace-Id: Root=1-621ef945-100e6185277d647165951066;Sampled=0

{
  "asignaciones_realizadas": 2,
  "asignaciones_exitosas": 1,
  "asignaciones_no_exitosas": 1,
  "promedio_inversión_exitosa": 10000.00,
  "promedio_inversión_no_exitosa": 200.00
}
```

#### Testing

Creamos el package **test_data**, código en el directorio **test_data**. Este define una función que genera los datos de prueba. Devuelve una slice donde cada elemento es un juego de datos de prueba. 

Cada juego de datos es una estructura que contiene:

- ImpAsignar: monto a distribuir.
- ImpResultado: resultado de la distribución hecha por el algoritmo.
- Esperado: Lo que esperamos como resultado (true distribuible, false no).
- Recibido: Lo que determina el algoritmo (true distribuible, false no).
- Id: Este campo nos será necesario en el test del nivel avanzado.

La función genera juegos de datos aleatoriamente, que se agrupan, según porcientos prefijados, en los siguientes casos:

- Montos no divisibles por 100 y montos no positivos: Aunque el planteamiento de la tarea excluía estos casos, consideramos conveniente incluirlos tanto en los programas como en los test. En este caso esperamos, claro está, que no es distribuible.

- Montos divisibles por 100 no distribuibles: Aquí nos basamos en casos conocidos, ya que determinarlos implica usar el mismo algoritmo de Diofanto usado en los programas. Al parecer se reducen a los montos 100, 200 y 400, mediante prueba empírica hecha hasta 100000. El resultado debe ser no distribuible.

- Montos divisibles por 100 distribuibles: Los generamos aplicando la fórmula: 300x + 500y + 500z, asignando valores positivos aleatorios a x, y, z. El resultado debe ser distribuible y retornar el mismo monto, aunque posiblemente la distribución del algoritmo no sea la misma que usamos al generar.

Teniendo la generación de datos de prueba preparada, pasamos a los test.

**Test unitario 1**

Este test prueba el algoritmo utilizado para determinar si un monto es distribuible y, de serlo, determinar la distribución. Esto corresponde al package **asigna** en el directorio **asigna**. Creamos el test **testDiofanto** contenido en **asigna/asigna_test.go**.

El primer paso es usar la interface **NewCreditAssigner()** para evaluar, para cada juego de datos de prueba, si es distribuible o no. En caso de serlo guardamos ese resultado en Recibido, y calculamos y guardamos el monto de la distribución recibida, usando la fórmula y los x, y, z que obtenemos. En caso de no ser distribuible guardamos esa condición en Recibido.

En este primer paso en vez de hacer los cálculos secuencialmente lo implementamos de manera concurrente, usando un número ajustable de goroutines. Esto hace más rápido el cálculo.

Concluido el primer paso, el segundo va comparando cada juego de datos, lo esperado con lo recibido. Los casos de fallos son:

- Lo esperado difiere de lo recibido.

- El monto es distribuible pero el monto generado por nuestra distribución difiere del monto de la distribución hecha por el algoritmo.

Nos ubicamos en el directorio **asigna** para correr** el test:

```
D:\yofio\asigna>go test -cover -race
PASS
coverage: 100.0% of statements
ok      yofio/asigna    1.891s
```

**Test unitario 2**

Este test lo realizamos en el nivel intermedio. Su objetivo es comprobar, además del algoritmo, la implementación del handler HTTP del servicio API REST. La implementación del handler está en el directorio **intermedio/handler**. Allí creamos el test **TestIntermedio** contenido en **handler_test.go**.

El test es muy similar al anterior, lo que en lugar de usar directamente la interface en **asigna**, hacemos llamadas HTTP a nuestro handler. Para ello usamos los packages de Go **http** y **httptest**. Para cada juego de datos creamos un http request y un http response recorder. Llamamos al handler y el resultado lo obtenemos del response recorder, en los campos StatusCode y Body.

Si el StatusCode es 200 OK, es distribuible y hacemos el unmarhal del Body para obtener la distribución x, y, z. Si no, el monto no es distribuible.

Nos ubicamos en el directorio **intermedio/handler** y corremos el test:

```
D:\yofio\intermedio\handler>go test -cover -race
PASS
coverage: 100.0% of statements
ok      yofio/intermedio/handler        1.824s
```

**Test de integración**

Lo realizamos al nivel avanzado. Aquí probaremos el mecanismo de nuestra REST API en la nube AWS y además el registro de cada petición en la base de datos DynamoDB, incluyendo la correcta actualización de las estadísticas.

Implementamos la función **TestAvanzado** en **avanzado/avanzado_test.go**. En este caso, para cada juego de datos de prueba, hacemos las llamadas al punto de entrada que nos da API Gateway, mediante el paquete http. El resultado en el StatusCode y el Body de la respuesta.

El mecanismo es similar al test anterior, con las diferencias:

- Paso 1: Hemos implementado esta API en una cuenta gratis de AWS, lo que impone límites de recursos y concurrencia a la máquina EC2, las lambdas, las llamadas API Gateway y a la tabla de DynamoDB. El nivel de concurrencia usado en los test anteriores supera estos límites, por lo que aquí el paso 1 se realiza de forma secuencial.

- Paso 2: Además de chequear los resultados de cada juego de datos de prueba, va acumulando la cantidad y los montos de las distribuciones obtenidas. Esto permite, al terminar el paso, llamar a la API para obtener el registro de estadísticas y compararlo con estos resultados. Si no coinciden, el test falla.

Para realizar este test, dado que vamos a comprobar el registro de estadísticas, debemos partir de una tabla **yofio** limpia y preparada. Para ello, de existir, debemos eliminarla y luego inicializarla:

```
D:\yofio\avanzado/local>go run . -d
D:\yofio\avanzado/local>go run . -c
```

Para correr el test:

```
D:\yofio\avanzado>go test
ok      yofio/avanzado  80.452s
```

Es de notar que en este caso no tiene sentido ver concurrencia, no la hay, ni coverage, él codigo reside en la nube.