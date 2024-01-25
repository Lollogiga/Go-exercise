# Go-exercise

Questo progetto Go include il codice per un'applicazione client-server con un load balancer. La logica di base è la seguente:

## Logica dell'applicazione

Un client invoca un servizio di richiesta dell'ora presso un server. Ci sono n repliche di tale server, ma ciò viene reso trasparente grazie all'utilizzo di un load balancer. Questo load balancer, situato tra il client e i server, intercetta le richieste del client e inoltra il servizio a uno degli n server seguendo una politica Round-Robin.

Il Load Balancer utilizza un file di configurazione (`configuration.txt`) per ottenere informazioni sugli indirizzi IP e i numeri di porta di ciascun server. Il client conosce solo la porta del Load Balancer.

## Struttura del Progetto

Progetto-Go/
|-- Go-exercise/
| |-- Client/
| | |-- main.go
| |-- Load_Balancer/
| | |-- main.go
| | |-- configuration.txt
| |-- Server/
| | |-- main.go


## Come Eseguire

Nella cartella di lavoro (`Progetto-Go/Go-exercise`), segui i seguenti passaggi:

1. **Lanciare n Server:**
   ```bash
   go run Server/main.go "numeroPorta"
2. **Lanciare il Load Balancer:**
   Aggiornare il file: Load_Balancer\configuration.txt affinchè abbia lo stesso numero di porta specificato nl passo precedente
   ```bash
   go run Load_Balancer/main.go
4. **Lanciare il client:**
   ```bash
   go run Client/main.go
