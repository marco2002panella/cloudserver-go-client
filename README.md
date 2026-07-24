## Seeweb Cloud Server Go Client

Un client Go robusto e modulare per comunicare direttamente con le API di backend del cloud server di Seeweb. Questo pacchetto è progettato per essere importato e integrato facilmente in qualsiasi applicazione scritta in Go che necessiti di interagire programmaticamente con l'infrastruttura Seeweb.

### Caratteristiche Principali

* **Architettura Modulare per Risorsa:** Ogni file sorgente gestisce i metodi dedicati a una specifica risorsa cloud, garantendo una codebase pulita e facilmente manutenibile.
* **Logging Tracciabile:** Il client genera automaticamente un registro dettagliato (log) di tutte le richieste HTTP inviate e delle relative risposte ricevute, utile per il monitoraggio e il debugging.
* **Integrazione Nativa:** Progettato come libreria riutilizzabile, importabile in qualsiasi progetto Go esterno.
* **Supporto alle API di Backend:** Comunica direttamente con gli endpoint di backend Seeweb.

---

## Installazione

Importa il client direttamente nel tuo modulo Go:

```bash
go get [github.com/marco2002panella/cloudserver-go-client](https://github.com/marco2002panella/cloudserver-go-client)
```
## Utilizzo Base

Ecco un esempio di come inizializzare il client e utilizzarlo all'interno del tuo codice Go:
```Go

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
)

func main() {
	fmt.Print("\033[H\033[2J")
	const (
		cyanBright = "\033[1;36m"
		green      = "\033[1;32m"
		yellow     = "\033[1;33m"
		reset      = "\033[0m"
	)
	banner := `
  ███████╗███████╗███████╗██╗    ██╗███████╗██████╗ 
  ██╔════╝██╔════╝██╔════╝██║    ██║██╔════╝██╔══██╗
  ███████╗█████╗  █████╗  ██║ █╗ ██║█████╗  ██████╔╝
  ╚════██║██╔══╝  ██╔══╝  ██║███╗██║██╔══╝  ██╔══██╗
  ███████║███████╗███████╗╚███╔███╔╝███████╗██████╔╝
  ╚══════╝╚══════╝╚══════╝ ╚══╝╚══╝ ╚══════╝╚══════╝ By Marco Panella (Not the original logo)`

	//read the XAPITOKEN from a file or from where you want to read it
	fmt.Println(cyanBright + banner + reset)
	tokenBytes, err := os.ReadFile("../token.txt")
	if err != nil {
		panic(fmt.Sprintf("Impossibile leggere il file token.txt: %v", err))
	}
	token := strings.TrimSpace(string(tokenBytes))

	client, err := seeweb.NewClient(&seeweb.Config{
		XAPITOKEN: token,
		DEBUG:     true,
		Timeout:   20, //timeout http request in seconds
		//JWTtoken: si puo usare tramite envToken ,
	})
	if err != nil {
		panic(err)
	}

	//list all the plan availables and write it in log responses if DEBUG==true or u can collect the response and do a print
	client.Plan.ListAvailables()

	resp, httpresp, err := client.Plan.ListAvailables()

}

## Struttura del Progetto

Il client adotta una suddivisione dei file basata sulle risorse dell'infrastruttura cloud, dove ciascun file contiene i metodi specifici per la gestione della risorsa corrispondente:
```Plaintext

.
├── client.go          # Configurazione principale, gestione HTTP client e log
├── server.go          # Metodi per la gestione delle istanze di calcolo (Cloud Server)
├── network.go         # Metodi per la configurazione delle reti e subnet
├── storage.go         # Metodi per la gestione dei volumi di storage
└── ...                # Altri file di risorsa
```
## Sistema di Logging

Il client include un meccanismo integrato di tracciamento che cattura payload, intestazioni (headers), codici di stato HTTP e tempi di risposta. Se abilitato, facilita l'ispezione del traffico di rete ed è particolarmente utile per identificare discrepanze dovute a endpoint non documentati o comportamenti anomali del backend.
Contribuire e Stato della Documentazione

    [!NOTE]
    Poiché la documentazione ufficiale di alcuni endpoint di backend Seeweb risulta incompleta o assente, il comportamento di alcune funzioni è stato dedotto tramite analisi empirica e reverse-engineering delle chiamate API.

Se riscontri comportamenti inattesi o desideri aggiungere supporto a nuovi endpoint:

    Apri una Issue descrivendo il problema riscontrato con la specifica risorsa.

    Invia una Pull Request con le correzioni o i miglioramenti ai metodi di comunicazione.

## Licenza

Questo progetto è distribuito sotto licenza MIT.
