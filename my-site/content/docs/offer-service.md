---
title: 'Angebot Service'
---

Dieser Service ermöglicht das Erstellen, Abrufen und Filtern von Angeboten über eine einfache REST-API. Er unterstützt strukturierte Angebotsobjekte mithilfe von Builder-Pattern-Klassen und basiert auf einer tokenbasierten Authentifizierung.

## 🛠️ Installation

1. Binde die `Client`, `OfferBuilder`, `SpaceBuilder`, `ItemBuilder`, etc. Klassen in deinem Frontend ein.
2. Stelle sicher, dass ein Backend mit folgenden Routen verfügbar ist:

   * `POST /api/angebot`
   * `GET /api/angebot/:id`
   * `POST /api/angebot/filter`

---

## 🚀 Verwendung

### ✅ Angebot erstellen

```js
const client = new Client();
await client.login("user@example.com", "password");

const size = new SizeBuilder().setWidth(100).setHeight(50).setDepth(30);
const item = new ItemBuilder().setSize(size).setWeight(15);
const space = new SpaceBuilder().addItem(item).setSeats(2);

const locationFrom = new LocationBuilder().setLatitude(52.52).setLongitude(13.405);
const locationTo = new LocationBuilder().setLatitude(48.1351).setLongitude(11.582);

const offer = new OfferBuilder()
  .setTitle("Transport für Möbel")
  .setDescription("Ich kann dein Sofa mitnehmen.")
  .setPrice(50)
  .setLocationFrom(locationFrom.build())
  .setLocationTo(locationTo.build())
  .setStartDateTime(new Date().toISOString())
  .setEndDateTime(new Date(Date.now() + 3600000).toISOString())
  .setCanTransport(space.build())
  .build();

await client.createOffer(offer);
```

---

### 🔍 Angebote filtern

```js
const filter = new FilterBuilder()
  .setNameStartsWith("Transport")
  .setLocationFrom(locationFrom)
  .setLocationTo(locationTo)
  .build();

const results = await client.getOffersByFilter(filter);
console.log(results);
```

---

### 📄 Angebot abrufen

```js
const offer = await client.getOfferById("angebot-id-123");
console.log(offer);
```

---

## 🧱 Builder-Klassen

| Klasse            | Zweck                                  |
| ----------------- | -------------------------------------- |
| `OfferBuilder`    | Erstellen eines vollständigen Angebots |
| `SpaceBuilder`    | Beschreibt den verfügbaren Platz       |
| `ItemBuilder`     | Einzelnes zu transportierendes Objekt  |
| `SizeBuilder`     | Dimensionen eines Items                |
| `LocationBuilder` | Geografische Koordinaten               |
| `FilterBuilder`   | Filterkriterien für Angebotssuche      |

Jede `build()`-Methode validiert automatisch alle Pflichtfelder und wirft bei Fehlern eine aussagekräftige Exception.

---

## 🌐 API-Endpunkte

| Methode | Pfad                  | Beschreibung                         |
| ------- | --------------------- | ------------------------------------ |
| POST    | `/api/angebot`        | Erstellt ein neues Angebot           |
| GET     | `/api/angebot/:id`    | Holt ein Angebot per ID              |
| POST    | `/api/angebot/filter` | Holt eine Liste gefilterter Angebote |

---

## ⚠️ Fehlerbehandlung

Alle HTTP-Aufrufe werfen Exceptions bei Fehlschlägen. Fehlermeldungen werden als `Error`-Objekte mit Statuscodes und Text bereitgestellt:

```js
try {
  await client.createOffer(myOffer);
} catch (err) {
  console.error("Fehler beim Erstellen des Angebots:", err.message);
}
```

---

## 🔐 Authentifizierung

Der Nutzer muss vor dem Erstellen eines Angebots angemeldet sein. Ein gültiges Token wird intern vom `Client` verwaltet:

```js
await client.login("email", "password"); // Token wird automatisch gesetzt
```

---

## 💬 WebSocket Integration (optional)

Nach erfolgreicher Anmeldung wird automatisch ein WebSocket geöffnet, z.B. für Echtzeitkommunikation:

```js
client.registerOnMessage((msg) => {
  console.log("Neue Nachricht:", msg);
});
```
