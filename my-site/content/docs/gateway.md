---
title: 'Gateway'
---

```mermaid
classDiagram
    direction TB

    class Service {
        +*server.Server
        +*auth.AuthMiddleware
        +*jwt.Decoder
        +NR *natsreciver.Receiver
        +WSHandler(w http.ResponseWriter, r *http.Request)
        +HandleChatWS(w http.ResponseWriter, r *http.Request)
        +HandleTracking(w http.ResponseWriter, r *http.Request)
        +HealthCheck(w http.ResponseWriter, r *http.Request)
        +LogNats()
        +Close() error
    }

    class server.Server {
        +GetLogger() zerolog.Logger
        +WithHandlerFunc(path string, handler http.HandlerFunc, methods...)
        +Error(w http.ResponseWriter, message string, code int)
    }

    class auth.AuthMiddleware {
        +EnsureJWT(handler http.HandlerFunc) http.HandlerFunc
    }

    class jwt.Decoder {
        +DecodeUUID(token string) (uuid.UUID, error)
    }

    class natsreciver.Receiver {
        +Subscribe(subject string, handler func(*nats.Msg)) (Subscription, error)
        +Publish(subject string, data []byte) error
        +Close() error
    }

    class zerolog.Logger {
        +Info()
        +Debug()
        +Err(error)
    }

    class http.Request
    class http.ResponseWriter
    class websocket.Conn

    class TrackingRequest {
        +Location repoangebot.Location
    }

    Service --> server.Server
    Service --> auth.AuthMiddleware
    Service --> jwt.Decoder
    Service --> natsreciver.Receiver

    Service --> websocket.Conn : Uses in WSHandler
    Service --> zerolog.Logger : Uses for logging
    Service --> TrackingRequest : Uses in HandleTracking
    Service --> http.ResponseWriter : Handles HTTP
    Service --> http.Request : Handles HTTP

```

# Gateway Service - Dokumentation

Der Gateway-Service ist ein zentraler HTTP-Proxy, der eingehende Anfragen basierend auf dem URL-Pfad an die jeweils zuständigen Microservices weiterleitet.

---

## Zweck

* Zentrale Anlaufstelle für Clients
* Weiterleitung von HTTP-Anfragen an Microservices
* Einfaches Routing basierend auf URL-Präfixen
* Einheitliche Schnittstelle für mehrere Backend-Services

---

## Routing-Regeln

| URL-Pfad-Präfix | Zielservice                                         | Beschreibung                   |
|-----------------|-----------------------------------------------------| ------------------------------ |
| `api/user/*`    | User-Service (z.B. `http://userservice:8080`)       | Benutzerverwaltung             |
| `api/angebot/*` | Angebot-Service (z.B. `http://angebotservice:8080`) | Angebotsverwaltung             |
| `api/media/*`   | Media-Service (z.B. `http://mediaservice:8080`)     | Medienverwaltung               |
| `/` (andere)    | Frontend oder 404-Fehler                            | Standardseite oder Fehlerseite |

---

## Nutzung als Client

Clients senden alle Anfragen an das Gateway (z.B. `https://localhost/api`). Das Gateway sorgt automatisch dafür, dass die Anfragen beim richtigen Service landen.

---

## Beispiel-Endpunkte (über Gateway)

| HTTP Methode | Pfad             | Zweck                        |
| ------------ |------------------| ---------------------------- |
| GET          | `api/user/`      | Liste aller Benutzer abrufen |
| POST         | `api/user/`      | Neuen Benutzer anlegen       |
| PUT          | `api/user/{id}`  | Benutzer aktualisieren       |
| DELETE       | `api/user/{id}`  | Benutzer löschen             |
| GET          | `api/angebot/`   | Angebote abrufen             |
| GET          | `api/media/{id}` | Mediendatei abrufen          |
| GET          | `/`              | Frontend oder Startseite     |

---

## Vorteile

* **Vereinfachte Client-Integration:** Ein einziger Endpunkt für alle Services
* **Flexibles Routing:** Einfaches Hinzufügen neuer Services und Pfade
* **Zentrale Sicherheitskontrolle:** Authentifizierung, Logging oder Rate-Limiting können hier zentral implementiert werden
* **Entkopplung der Clients von den Backend-Services**

---

## Beispiel Client-Code (JavaScript)

```js
const BASE_URL = "https://localhost";

// Benutzerliste abrufen
async function fetchUsers() {
  const response = await fetch(`${BASE_URL}/users/`);
  if (!response.ok) throw new Error("Fehler beim Abrufen der Benutzer");
  return response.json();
}

// Benutzer erstellen
async function createUser(userData) {
  const response = await fetch(`${BASE_URL}/users/`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(userData),
  });
  if (!response.ok) throw new Error("Fehler beim Erstellen des Benutzers");
  return response.json();
}

// Angebot abrufen
async function fetchAngebote() {
  const response = await fetch(`${BASE_URL}/angebote/`);
  if (!response.ok) throw new Error("Fehler beim Abrufen der Angebote");
  return response.json();
}
```
