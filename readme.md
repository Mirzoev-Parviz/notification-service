# Notification Service

## Getting Started

### Step 1: Clone the repository
To get started, clone the repository and navigate into the project directory:

```bash
git clone https://github.com/Mirzoev-Parviz/notification_service.git
cd notification_service
```

### Step 2: Configure the Database
Create a PostgreSQL database and then configure the `config.yaml` file (you'll need to create this file based on the provided `config.example.yaml`):

```yaml
dbSettings:
  user: your_db_user
  password: your_db_password
  host: localhost
  port: 5432
  database: notification_db
  sslmode: disable
```

Ensure you fill in your actual database credentials.

### Step 3: Build the Project
Build the project by running the following command:

```bash
go build -o notification_service ./cmd/main.go
```

### Step 4: Run the Application
Once built, you can start the application by running:

```bash
./notification_service
```

The application should start and listen on `localhost:8080`.

### Step 5: Send Events
You can now send POST requests with card events to `http://localhost:8080/event`.

### Example Requests

#### Example 1: Purchase Event
```json
{
  "orderType": "Purchase",
  "sessionId": "29827525-06c9-4b1e-9d9b-7c4584e82f56",
  "card": "4433**1409",
  "eventDate": "2023-01-04 13:44:52.835626 +00:00",
  "websiteUrl": "https://amazon.com"
}
```

#### Example 2: Card Verification Event
```json
{
  "orderType": "CardVerify",
  "sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4",
  "card": "4433**1409",
  "eventDate": "2023-04-07 05:29:54.362216 +00:00",
  "websiteUrl": "https://adidas.com"
}
```

#### Example 3: Send OTP Event
```json
{
  "orderType": "SendOtp",
  "sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4",
  "card": "4433**1409",
  "eventDate": "2023-04-06 22:52:34.930150 +00:00",
  "websiteUrl": "https://somon.tj"
}
```

### Step 6: Notification result
Just check the notifications printed in the terminal.

### Step 7: Shut Down
Stop the server gracefully by pressing `Ctrl+C`.


#### And if all else fails... well, there's always more coffee ☕.