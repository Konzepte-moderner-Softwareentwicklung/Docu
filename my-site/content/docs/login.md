---
title: 'Login'
---

# Dokumentation zum Login im `UserController`

## Übersicht

Der `UserController` bietet mehrere Endpunkte zur Authentifizierung von Benutzern, sowohl über klassische E-Mail/Passwort-Logins als auch über WebAuthn (Passkey)-basierte Authentifizierung. Das Ziel ist es, den Benutzer sicher zu authentifizieren und bei Erfolg ein JWT-Token zurückzugeben, das für 24 Stunden gültig ist.

---

## Endpunkte und Funktionen

### 1. Klassischer Login: `/users/login` (POST)

**Funktion:** `GetLoginToken`

* **Beschreibung:**
  Authentifiziert einen Benutzer anhand von E-Mail und Passwort.
  Bei erfolgreicher Authentifizierung wird ein JWT-Token generiert, das 24 Stunden gültig ist.

* **Request Body (JSON):**

  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```

* **Ablauf:**

    1. E-Mail und Passwort werden aus dem JSON-Request ausgelesen.
    2. Benutzer wird anhand der E-Mail gesucht.
    3. Passwort wird mit dem gespeicherten Hash verglichen (über `hasher.VerifyPassword`).
    4. Bei Erfolg wird ein JWT-Token erzeugt (`EncodeUUID`) und als JSON zurückgegeben.
    5. Bei Fehlern wird ein entsprechender HTTP-Statuscode mit Fehlermeldung zurückgegeben.

* **Antwort (bei Erfolg):**

  ```json
  {
    "token": "JWT-Token"
  }
  ```

* **Fehler:**

    * 400: Fehler beim Lesen der Anfrage
    * 401: Falsches Passwort
    * 500: Nutzer nicht gefunden oder Fehler bei Token-Generierung

---

### 2. WebAuthn Login

WebAuthn-Login erfolgt in zwei Schritten: Optionen abrufen und Login abschließen.

#### a) Begin WebAuthn Login: `/users/webauthn/login/options` (GET)

**Funktion:** `beginLogin`

* **Beschreibung:**
  Startet den WebAuthn-Login-Prozess, indem Login-Optionen generiert werden.

* **Query Parameter:**

    * `email`: E-Mail-Adresse des Benutzers

* **Ablauf:**

    1. Validiert die E-Mail-Adresse.
    2. Benutzer wird anhand der E-Mail gesucht.
    3. Login-Optionen und Session-Daten werden erzeugt (`service.webauth.BeginLogin`).
    4. Session-Daten werden im Benutzerobjekt gespeichert.
    5. Login-Optionen werden als JSON zurückgegeben.

* **Antwort (bei Erfolg):** JSON mit WebAuthn-CredentialAssertion.

* **Fehler:**

    * 400: Ungültige E-Mail-Adresse
    * 404: Benutzer nicht gefunden
    * 500: Interner Serverfehler

---

#### b) Finish WebAuthn Login: `/users/webauthn/login` (POST)

**Funktion:** `finishLogin`

* **Beschreibung:**
  Verifiziert die WebAuthn-Anmeldung und gibt bei Erfolg ein JWT-Token zurück.

* **Query Parameter:**

    * `email`: E-Mail-Adresse des Benutzers

* **Ablauf:**

    1. Validiert die E-Mail-Adresse.
    2. Benutzer wird anhand der E-Mail gesucht.
    3. WebAuthn-Login wird mit Session-Daten überprüft (`service.webauth.FinishLogin`).
    4. Bei Erfolg wird ein JWT-Token für 24 Stunden generiert und zurückgegeben.

* **Antwort (bei Erfolg):**

  ```json
  {
    "token": "JWT-Token"
  }
  ```

* **Fehler:**

    * 400: Ungültige E-Mail-Adresse
    * 401: Authentifizierung fehlgeschlagen
    * 404: Benutzer nicht gefunden
    * 500: Interner Serverfehler

---

## Wichtige Details

* **JWT-Token**:
  Ein JWT wird mit der Benutzer-ID als Payload erzeugt, gültig für 24 Stunden.

* **Session-Daten bei WebAuthn:**

    * `BeginLogin` und `FinishLogin` verwenden sessionData, die temporär im User-Objekt gespeichert werden.
    * Diese Session-Daten sind notwendig zur Validierung des WebAuthn-Ablaufs.

* **Fehlerbehandlung:**
  Alle Fehler führen zu einer aussagekräftigen HTTP-Antwort mit Statuscode und Fehlermeldung.


