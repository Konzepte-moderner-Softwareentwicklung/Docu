---
title: 'User Service'
---

Dieser Service ermöglicht das **Anlegen und Abrufen von Benutzern** über eine einfache API. Die Authentifizierung (Login, Token, WebAuthn) ist hierbei **nicht erforderlich**.

## 🚀 Funktionen

### 🔹 Benutzer anlegen

Verwende den bereitgestellten `UserBuilder`, um einen neuen Benutzer zu erstellen. Dies ist ohne vorherige Anmeldung möglich.

#### Beispiel:

```js
const client = new Client();

const user = new UserBuilder()
  .setFirstName("Max")
  .setLastName("Mustermann")
  .setEmail("max@example.com")
  .setPassword("geheim123")
  .setBirthDate("1990-01-01")
  .setPhoneNumber("0123456789")
  .setProfilePicture("https://example.com/profile.jpg")
  .build();

await client.createUser(user);
```

🔐 Hinweis: Das Passwort wird direkt im Klartext übergeben. Stelle sicher, dass du HTTPS verwendest.

---

### 🔹 Benutzer abrufen

#### Alle Benutzer abrufen

```js
const users = await client.getUsers();
```

#### Benutzer nach E-Mail

```js
const user = await client.getUserByEmail("max@example.com");
```

#### Benutzer nach ID

```js
const user = await client.getUserById("user_id_xyz");
```

---

## 📦 Abhängigkeiten

* Die Kommunikation erfolgt über die `fetch`-API.
* JSON als Austauschformat.
* Keine Authentifizierung erforderlich (für die oben genannten Funktionen).

---

## 🧰 Tools

* `Client`: Hauptschnittstelle zur API.
* `UserBuilder`: Hilfsklasse zum einfachen und validierten Aufbau eines Benutzers.

---

## ⚠️ Validierung

Der `UserBuilder` erzwingt folgende Pflichtfelder:

* `firstName`
* `lastName`
* `email`
* `password`

Optional können angegeben werden:

* `birthDate`
* `phoneNumber`
* `profilePicture`

Falls Pflichtfelder fehlen, wird beim Aufruf von `.build()` ein Fehler geworfen.

---

## 📝 Beispiel-Workflow

```js
const client = new Client();

try {
  const user = new UserBuilder()
    .setFirstName("Anna")
    .setLastName("Beispiel")
    .setEmail("anna@example.com")
    .setPassword("passwort123")
    .build();

  const response = await client.createUser(user);
  console.log("Benutzer erfolgreich erstellt:", response);
} catch (error) {
  console.error("Fehler beim Erstellen des Benutzers:", error);
}
```
