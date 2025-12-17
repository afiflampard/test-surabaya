# Kumparan Test API

Backend service berbasis Go (CloudWeGo Hertz) untuk mengelola data sensor dan perintah kontrol perangkat, dilengkapi integrasi PostgreSQL, MQTT (EMQX), dan dokumentasi Swagger.

## Fitur

- **Create Sensor Data**: Simpan data suhu/kelembapan dan publikasikan ke MQTT.
- **Create Device Control**: Simpan dan kirim perintah kontrol perangkat via MQTT.
- **Health Check**: Cek konektivitas Database dan MQTT.
- **Swagger**: Dokumentasi API tersedia di `/swagger/index.html`.

## Teknologi
- Go + CloudWeGo Hertz
- PostgreSQL (sqlx)
- MQTT (Eclipse Paho) dengan broker EMQX
- Elasticsearch + Kibana (opsional, tersedia pada compose)
- Docker & docker-compose

## Struktur Proyek (ringkas)
```
api/
  handler/handler.go        # HTTP handlers
  router/router.go          # Route definitions
  service/service.go        # Application service layer
config/                     # Load env, definisi Config
domain/
  infra/                    # Init Postgres, MQTT, health checks
  sensor/                   # Entity, repo, mutation, validation
main.go                     # App entrypoint
Dockerfile                  # Image build
docker-compose.yaml         # Dev stack (DB, EMQX, ES, Kibana)
```

## Konfigurasi Environment
Variabel env dibaca via `envconfig` dan `.env` (lihat `config/app_config.go`).

- `APP_NAME` (default: HertzApp)
- `PORT` (default: 8080)
- `DB_HOST` (required)
- `DB_PORT` (default: 5432)
- `DB_USER` (required)
- `DB_PASSWORD` (required)
- `DB_NAME` (required)
- `JWT_SECRET` (required)
- `DB_REPLICA_HOST` (required)
- `DB_REPLICA_PORT` (default: 5433)
- `ELASTIC_URL` (default: http://localhost:9200)
- `MQTT_BROKER` (default: tcp://localhost:1883)

Contoh nilai saat memakai docker-compose:
```
PORT=8080
DB_HOST=postgres-primary
DB_PORT=5432
DB_USER=hertz_user
DB_PASSWORD=hertz_pass
DB_NAME=hertz_db
JWT_SECRET=supersecret
DB_REPLICA_HOST=postgres-replica
DB_REPLICA_PORT=5432
ELASTIC_URL=http://elasticsearch:9200
MQTT_BROKER=tcp://emqx:1883
```

## Menjalankan Aplikasi

### 1) Lokal (tanpa Docker)
- Pastikan PostgreSQL dan EMQX berjalan lokal sesuai env.
- Buat file `.env` di root project (sesuai variabel di atas).
- Jalankan:
```
go run main.go
```
- Aplikasi akan berjalan pada `http://localhost:<PORT>` (default 8080).

### 2) Docker (image aplikasi saja)
```
docker build -t kumparan-test-api .
docker run --rm -p 8080:8080 --env-file .env kumparan-test-api
```

### 3) docker-compose (recommended untuk dev)
docker-compose sudah menyiapkan PostgreSQL primary/replica, Elasticsearch+Kibana, dan EMQX. Service app masih dikomentari agar fleksibel.

- Jalankan infra:
```
docker compose up -d postgres-primary postgres-replica elasticsearch kibana emqx
```
- Setelah semua siap, jalankan app lokal (pakai `.env` yang mengarah ke service compose), atau tambahkan service app di compose dengan env yang sama (uncomment dan sesuaikan bagian `hertz-app`).

## Endpoint API

- `GET /health`
  - Response 200 jika sehat, 503 jika salah satu dependency down.

- `POST /sensor/create-sensor`
  - Body JSON:
    ```json
    {
      "device_id": "device-123",
      "temperature": 25.5,
      "humidity": 60
    }
    ```
  - Mengembalikan `uuid` id data sensor yang dibuat, dan publish MQTT ke topik `sensor/data/<device_id>`.

- `POST /sensor/create-control`
  - Body JSON:
    ```json
    {
      "device_id": "device-123",
      "command": "FAN_ON"
    }
    ```
  - Mengembalikan `uuid` id perintah kontrol yang dibuat, dan publish MQTT ke topik `greenhouse/control/<device_id>`.

- `GET /swagger/*any`
  - Akses dokumentasi di `http://localhost:8080/swagger/index.html`.

## Validasi Input (ringkas)
- Sensor Data: `device_id` wajib, `temperature` >= -273.15, `humidity` 0..100.
- Device Control: `device_id` dan `command` wajib.

## Catatan Pengembangan
- Health check memeriksa koneksi MQTT dan Postgres (`/health`).
- MQTT client diinit dari `MQTT_BROKER` (contoh: `tcp://emqx:1883`).
- Pastikan schema tabel tersedia (lihat implementasi repository di `domain/sensor`).

## Lisensi
Untuk kebutuhan tes/penilaian internal.

