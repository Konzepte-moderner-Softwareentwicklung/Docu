---
date: '2025-06-05T21:59:11+02:00'
draft: true
title: 'Media Service'
---

# Media Service - Dokumentation

Der Media Service stellt Endpunkte zum Upload und Download von Bildern bereit und verwaltet einzelne sowie mehrere (komplexe) Bild-Links.

---

## Basis-URL

`/media`

---

## Endpunkte

### 1. `GET /media/image`

* **Beschreibung:** Test- oder Index-Endpunkt, gibt "Hello World" zurück.
* **Request Body:** Kein
* **Response:** Plain Text `"Hello World"`
* **Status Codes:** 200 OK

---

### 2. `POST /media/image`

* **Beschreibung:** Upload eines einzelnen Bildes.
* **Request Header:**

    * `Content-Type`: Muss den Medientyp des Bildes angeben (z.B. `image/jpeg`).
    * `UserId`: ID des hochladenden Benutzers (wird als Header erwartet).
* **Request Body:** Binärdaten des Bildes.
* **Response Body (JSON):**

  ```json
  {
    "name": "string",
    "success": true
  }
  ```
* **Status Codes:**

    * 200 OK bei Erfolg
    * 400 Bad Request, wenn Content-Type oder UserId fehlt
    * 500 Internal Server Error bei Upload-Fehlern

---

### 3. `GET /media/image/{id}`

* **Beschreibung:** Download eines Bildes nach Bild-ID (Name).
* **URL-Parameter:**

    * `id` (string): Bildname/ID.
* **Response:**

    * Bilddaten mit Header `Content-Type: image/jpeg`
* **Status Codes:**

    * 200 OK bei Erfolg
    * 400 Bad Request, wenn `id` fehlt
    * 500 Internal Server Error bei Fehlern

---

### 4. `GET /media/multi/{id}`

* **Beschreibung:** Gibt eine Liste von Bild-URLs (compound links) zurück, die zu einer zusammengesetzten Entität gehören.
* **URL-Parameter:**

    * `id` (UUID): ID der zusammengesetzten Entität.
* **Response Body (JSON):** Array von URLs, z.B.

  ```json
  [
    "/media/image/abc123",
    "/media/image/def456"
  ]
  ```
* **Status Codes:**

    * 200 OK bei Erfolg
    * 400 Bad Request, wenn `id` fehlt oder ungültig ist
    * 500 Internal Server Error bei Fehlern

---

### 5. `POST /media/multi/{id}`

* **Beschreibung:** Upload eines Bildes zu einer zusammengesetzten Entität.
* **URL-Parameter:**

    * `id` (string): ID der zusammengesetzten Entität.
* **Request Header:**

    * `Content-Type`: Medientyp des Bildes
    * `UserId`: ID des Benutzers (muss gesetzt sein)
* **Request Body:** Binärdaten des Bildes
* **Response:** Kein Body, nur Status-Code
* **Status Codes:**

    * 200 OK bei Erfolg
    * 400 Bad Request bei fehlenden Headern oder Parametern
    * 500 Internal Server Error bei Upload-Fehlern

---

## Header

| Header   | Bedeutung                    | Erforderlich?  |
| -------- | ---------------------------- | -------------- |
| `UserId` | ID des Nutzers, der hochlädt | Ja bei Uploads |

---

## Fehlerhandling

* 400 Bad Request bei ungültigen oder fehlenden Parametern oder Headern.
* 500 Internal Server Error bei internen Fehlern.

---

## Beispiel Curl Upload (einzelnes Bild)

```bash
curl -X POST https://localhost/api/media/image \
  -H "Content-Type: image/jpeg" \
  -H "UserId: 1234" \
  --data-binary "@pfad/zum/bild.jpg"
```

---

## Beispiel Curl Download (Bild)

```bash
curl https://localhost/media/image/abc123 > downloaded.jpg
```