# Zenoty
Zenoty is a notification service built with Golang that function currently to deliver a message from email or FCM. Using RabbitMQ Message Broker for executing a message from the background service to achieve a better performance and scability.

## Requirement
- Go 1.20 or later
- RabbitMQ installed on your machine (https://www.rabbitmq.com/download.html)
- Firebase Cloud Messaging

## Installation

1. Clone this repository to your local machine
    ```bash  
   git clone https://github.com/rafiseptian90/zenoty.git
   ``` 
2. Enter to the project path
   ```bash  
   cd zenoty
   ```
3. Configure the app by copy and rename .env.example to .env file and after that, you should fill out the required key on that .env file. The required key marked as (*).
4. Create your own FCM project in Firebase, and generate your service account json file and finally put that file to root config folder and rename it to "fcm_config.json"
5. Build the app
   ```go  
   go build -o zenoty cmd/main.go
   ```
6. Run the app <br/> On Windows OS
   ```bash  
   ./zenoty
   ```
   On Unix based OS
   ```bash  
   chmod +x zenoty
   ```
   ```bash
   zenoty
   ```
## Creator

>:boy: [Rafi Septian Hadi](https://github.com/rafiseptian90)

## License
This project is licensed under the Apache License 2.0