---
title: 'Login'
---
Dieser Service ermöglicht es Benutzern, sich über E-Mail/Passwort oder passwortlos via WebAuthn (Passkey) anzumelden. Die Authentifizierung erzeugt ein Token, das für nachfolgende geschützte Anfragen benötigt wird.


## 📥 Klassische Anmeldung: E-Mail & Passwort

### 🔧 Methode

```js
await client.login(email, password);
```

### 🧾 Beschreibung

Authentifiziert einen Benutzer mit E-Mail und Passwort. Beim Erfolg wird automatisch ein Authentifizierungstoken gespeichert und eine WebSocket-Verbindung geöffnet.

### 📌 Parameter

| Name       | Typ    | Beschreibung     |
| ---------- | ------ | ---------------- |
| `email`    | string | Benutzer-E-Mail  |
| `password` | string | Benutzerpasswort |

### ✅ Beispiel

```js
await client.login("benutzer@example.com", "meinPasswort123");
```

### ⚠️ Fehlerbehandlung

Falls die Anmeldung fehlschlägt, wird eine aussagekräftige Fehlermeldung ausgelöst.

---

## 🔐 Passwortlos anmelden mit Passkey (WebAuthn)

### 🔧 Methode

```js
await client.loginPasskey(email);
```

### 🧾 Beschreibung

Führt eine passwortlose Anmeldung per Passkey (z. B. FaceID, TouchID, FIDO2-Sicherheitsschlüssel) durch. Diese Methode nutzt den WebAuthn-Standard zur sicheren Authentifizierung.

### 📌 Parameter

| Name    | Typ    | Beschreibung                     |
| ------- | ------ | -------------------------------- |
| `email` | string | Die registrierte Benutzer-E-Mail |

### ✅ Beispiel

```js
await client.loginPasskey("benutzer@example.com");
```

### 🔄 Ablauf

1. Holt Anmeldeoptionen vom Server (`/api/user/webauthn/login/options`).
2. Führt WebAuthn-Authentifizierung durch.
3. Sendet Authentifizierungsdaten an den Server zur Verifikation (`/api/user/webauthn/login`).
4. Speichert Token und verbindet WebSocket.

### ⚠️ Fehlerbehandlung

* Ungültige E-Mail: Es wird eine Exception geworfen.
* Abbruch oder Verweigerung der WebAuthn-Anfrage durch den Benutzer führt zu einem Fehler.

---

## 🌐 WebSocket-Verbindung

Nach erfolgreichem Login (egal ob klassisch oder per Passkey) wird eine WebSocket-Verbindung automatisch aufgebaut:

```js
client.connectWebSocket();
```

Dies ermöglicht z. B. serverseitige Push-Nachrichten in Echtzeit.

---

## 🧪 Token-Nutzung

Das erhaltene Token wird in `client.token` gespeichert und automatisch in geschützten API-Requests (z. B. Angebotserstellung) verwendet.

---

## 📤 Logout

### 🔧 Methode

```js
await client.logout();
```

### Beschreibung

* Setzt das Token zurück.
* Trennt die WebSocket-Verbindung.

---

## 🔒 Sicherheitshinweise

* Alle Passkey-bezogenen Anfragen verwenden `credentials: "include"` und CORS-Modus.
* Die Anmeldedaten werden **nicht im Klartext** übertragen.
* Die WebAuthn-Authentifizierung erfolgt vollständig im Browser.

---

## 📚 Weitere Methoden

Für Registrierung mit Passkey siehe:

```js
await client.registerPasskey();
```

Für Benutzer-Management siehe Methoden wie `getUsers()`, `createUser()`, etc.
