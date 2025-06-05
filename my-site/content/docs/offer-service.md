---
date: '2025-06-05T21:59:11+02:00'
draft: true
title: 'Angebot Service'
---

# Offer Service - Dokumentation

Der Offer Service bietet REST-Endpoints zum Erstellen, Abfragen einzelner Angebote sowie zur Abfrage von Angeboten nach Filterkriterien.

---

## Basis-URL

`/offers`

---

## Endpunkte

### 1. `POST /offers/`

* **Beschreibung:** Erstellt ein neues Angebot.
* **Request Header:**

    * `UserId` (UUID): ID des Benutzers, der das Angebot erstellt.
* **Request Body (JSON):**

  ```json
  {
    "title": "string",
    "description": "string",
    "price": "float",
    // weitere Angebotsfelder entsprechend repoangebot.Offer
  }
  ```
* **Response Body (JSON):**

  ```json
  {
    "id": "string",          // ID des erstellten Angebots (UUID)
    "image_url": "string"    // URL für compound images (vom Media-Service)
  }
  ```
* **Status Codes:**

    * 200 OK bei Erfolg
    * 401 Unauthorized bei ungültiger oder fehlender UserId im Header
    * 500 Internal Server Error bei Fehlern beim Erstellen

---

### 2. `GET /offers/{id}`

* **Beschreibung:** Holt ein Angebot anhand seiner ID.
* **URL-Parameter:**

    * `id` (UUID): ID des Angebots.
* **Response Body (JSON):** Objekt des Angebots (entsprechend `repoangebot.Offer`).
* **Status Codes:**

    * 200 OK bei Erfolg
    * 401 Unauthorized bei ungültiger ID
    * 500 Internal Server Error bei Fehlern

---

### 3. `GET /offers/`

* **Beschreibung:** Holt Angebote, die einem Filter entsprechen.
* **Request Body (JSON):**

  ```json
  {
    // Filterkriterien entsprechend repoangebot.Filter,
    // z.B. Preisbereich, Suchbegriffe etc.
  }
  ```
* **Response Body (JSON):** Liste von Angeboten.
* **Status Codes:**

    * 200 OK bei Erfolg
    * 400 Bad Request bei fehlerhaftem Filter-JSON
    * 500 Internal Server Error bei Fehlern

---

## Header

| Header | Bedeutung             | Erforderlich? |
| ------ | --------------------- | ------------- |
| UserId | ID des Nutzers (UUID) | Ja bei POST   |

---

## Beispiel Curl Create Offer

```bash
curl -X POST http://localhost:8083/offers/ \
  -H "Content-Type: application/json" \
  -H "UserId: 123e4567-e89b-12d3-a456-426614174000" \
  -d '{"title": "Neues Angebot", "description": "Beschreibung", "price": 19.99}'
```

---

## Fehlerhandling

* 401 Unauthorized: Ungültige oder fehlende UserId bei geschützten Endpunkten.
* 400 Bad Request: Fehlerhafte Eingaben, z.B. Filter-JSON.
* 500 Internal Server Error: Interne Fehler beim Service.

---
