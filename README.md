##Seeweb Cloud Server Go Client

Un client Go robusto e modulare per comunicare direttamente con le API di backend del cloud server di Seeweb. Questo pacchetto è progettato per essere importato e integrato facilmente in qualsiasi applicazione scritta in Go che necessiti di interagire programmaticamente con l'infrastruttura Seeweb.
Caratteristiche Principali

    Architettura Modulare per Risorsa: Ogni file sorgente gestisce i metodi dedicati a una specifica risorsa cloud, garantendo una codebase pulita e facilmente manutenibile.

    Logging Tracciabile: Il client genera automaticamente un registro dettagliato (log) di tutte le richieste HTTP inviate e delle relative risposte ricevute, utile per il monitoraggio e il debugging.

    Integrazione Nativa: Progettato come libreria riutilizzabile, importabile in qualsiasi progetto Go esterno.

    Supporto alle API di Backend: Comunica direttamente con gli endpoint di backend Seeweb.

##Installazione

Importa il client direttamente nel tuo modulo Go:
Bash

go get github.com/marco2002panella/cloudserver-go-client

##Utilizzo Base

Ecco un esempio di come inizializzare il client e utilizzarlo all'interno del tuo codice Go:
Go

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/marco2002panella/cloudserver-go-client/client"
)

func main() {
	// Inizializza il client configurando le credenziali e l'endpoint di backend
	cfg := client.Config{
		APIKey:  "il-tuo-api-token",
		BaseURL: "https://api.seeweb.it/v1", // Sostituisci con l'endpoint effettivo
		Debug:   true,                     // Abilita la generazione dei log di richieste/risposte
	}

	seewebClient := client.NewClient(cfg)

	// Esempio di utilizzo dei metodi specifici per risorsa (es. Server)
	ctx := context.TODO()
	servers, err := seewebClient.Servers.List(ctx)
	if err != nil {
		log.Fatalf("Errore durante il recupero dei server: %v", err)
	}

	for _, server := range servers {
		fmt.Printf("Server ID: %s - Nome: %s\n", server.ID, server.Name)
	}
}

##Struttura del Progetto

Il client adotta una suddivisione dei file basata sulle risorse dell'infrastruttura cloud, dove ciascun file contiene i metodi specifici per la gestione della risorsa corrispondente:
Plaintext

.
├── client.go          # Configurazione principale, gestione HTTP client e log
├── server.go          # Metodi per la gestione delle istanze di calcolo (Cloud Server)
├── network.go         # Metodi per la configurazione delle reti e subnet
├── storage.go         # Metodi per la gestione dei volumi di storage
└── ...                # Altri file di risorsa

##Sistema di Logging

Il client include un meccanismo integrato di tracciamento che cattura payload, intestazioni (headers), codici di stato HTTP e tempi di risposta. Se abilitato, facilita l'ispezione del traffico di rete ed è particolarmente utile per identificare discrepanze dovute a endpoint non documentati o comportamenti anomali del backend.
Contribuire e Stato della Documentazione

    [!NOTE]
    Poiché la documentazione ufficiale di alcuni endpoint di backend Seeweb risulta incompleta o assente, il comportamento di alcune funzioni è stato dedotto tramite analisi empirica e reverse-engineering delle chiamate API.

Se riscontri comportamenti inattesi o desideri aggiungere supporto a nuovi endpoint:

    Apri una Issue descrivendo il problema riscontrato con la specifica risorsa.

    Invia una Pull Request con le correzioni o i miglioramenti ai metodi di comunicazione.

##Licenza

Questo progetto è distribuito sotto licenza MIT.

