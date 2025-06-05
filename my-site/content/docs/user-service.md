---
date: '2025-06-05T21:59:11+02:00'
draft: true
title: 'User Service'
---

# UserService API Dokumentation

Der UserService stellt eine REST-API für Benutzerverwaltung, Login und WebAuthn-Authentifizierung bereit.

> **Hinweis:** Die API setzt JWT-Token für geschützte Endpunkte voraus.

---

## Endpunkte

| Methode | Pfad                                    | Beschreibung                               | Auth nötig? |
| ------- | --------------------------------------- | ------------------------------------------ | ----------- |
| GET     | `/`                                     | Alle Benutzer abrufen                      | Nein        |
| POST    | `/`                                     | Benutzer erstellen                         | Nein        |
| GET     | `/{id}`                                 | Benutzer nach ID abrufen                   | Nein        |
| PUT     | `/{id}`                                 | Benutzer aktualisieren                     | Ja          |
| DELETE  | `/{id}`                                 | Benutzer löschen                           | Ja          |
| GET     | `/email?email={email}`                  | Benutzer per E-Mail abrufen                | Nein        |
| POST    | `/login`                                | Login mit Email und Passwort, JWT erhalten | Nein        |
| GET     | `/webauthn/register/options`            | WebAuthn Registrierung starten             | Ja          |
| POST    | `/webauthn/register`                    | WebAuthn Registrierung abschließen         | Ja          |
| GET     | `/webauthn/login/options?email={email}` | WebAuthn Login starten                     | Nein        |
| POST    | `/webauthn/login`                       | WebAuthn Login abschließen                 | Nein        |

---

## Beispiel JavaScript Client

```js
const API_BASE = "https://localhost/api/user"; // Beispielbase-URL

// Hilfsfunktion: HTTP Request mit JSON
async function request(path, method = "GET", body = null, token = null) {
  const headers = { "Content-Type": "application/json" };
  if (token) headers["Authorization"] = `Bearer ${token}`;

  const res = await fetch(`${API_BASE}${path}`, {
    method,
    headers,
    body: body ? JSON.stringify(body) : null,
  });

  if (!res.ok) {
    const error = await res.text();
    throw new Error(`HTTP ${res.status}: ${error}`);
  }

  if (res.status === 204) return null; // No Content
  return await res.json();
}

// Beispiel: Alle Benutzer abrufen
async function getUsers() {
  return await request("/");
}

// Beispiel: Benutzer erstellen
async function createUser(user) {
  return await request("/", "POST", user);
}

// Beispiel: Benutzer mit JWT aktualisieren
async function updateUser(id, user, token) {
  return await request(`/${id}`, "PUT", user, token);
}

// Beispiel: Benutzer löschen mit JWT
async function deleteUser(id, token) {
  return await request(`/${id}`, "DELETE", null, token);
}

// Beispiel: Login um JWT zu erhalten
async function login(email, password) {
  const data = await request("/login", "POST", { email, password });
  return data.token;
}

// Beispielnutzung
(async () => {
  try {
    // Benutzer erstellen
    const newUser = { email: "max@example.com", name: "Max Mustermann", password: "secret123" };
    const created = await createUser(newUser);
    console.log("User created with ID:", created.id);

    // Login und Token holen
    const token = await login(newUser.email, newUser.password);
    console.log("JWT Token:", token);

    // Alle Benutzer abrufen (ohne Auth)
    const users = await getUsers();
    console.log("Users:", users);

    // Benutzer aktualisieren mit Token
    await updateUser(created.id, { name: "Max M." }, token);
    console.log("User updated");

    // Benutzer löschen mit Token
    await deleteUser(created.id, token);
    console.log("User deleted");
  } catch (e) {
    console.error("Error:", e.message);
  }
})();
```

---

## Hinweise

* Für Endpunkte, die JWT-Token erfordern, muss das Token im Header `Authorization: Bearer <token>` mitgesendet werden.
* WebAuthn-Endpunkte sind für die Passwortlose Authentifizierung gedacht und benötigen Client-seitige Unterstützung (nicht im JS Beispiel enthalten).
* Die API unterstützt CORS, daher können Browser-Clients direkt mit der API kommunizieren.

