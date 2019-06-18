#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266WebServer.h>
#include <ESP8266mDNS.h>
#include <SimpleDHT.h>
#include <Thread.h>
#include <ESP8266HTTPClient.h>
#include <ArduinoJson.h>

#define MIN1 16
#define MIN2 5
#define CIN1 4
#define CIN2 0
const char* ssid = "snailpet";
const char* password = "wcg6t4pw";

ESP8266WebServer server(80);
HTTPClient http;
SimpleDHT11 dht11(2);
Thread thread_send = Thread();
Thread thread_receive = Thread();
void handleRoot() {
  server.send(200, "text/plain", "VAEMC IOT");
}

void handleNotFound() {
  String message = "File Not Found\n\n";
  message += "URI: ";
  message += server.uri();
  message += "\nMethod: ";
  message += (server.method() == HTTP_GET) ? "GET" : "POST";
  message += "\nArguments: ";
  message += server.args();
  message += "\n";
  for (uint8_t i = 0; i < server.args(); i++) {
    message += " " + server.argName(i) + ": " + server.arg(i) + "\n";
  }
  server.send(404, "text/plain", message);
}

void setup(void) {
  Serial.begin(115200);
  WiFi.begin(ssid, password);
  // Wait for connection
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected to ");
  Serial.println(ssid);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());
  if (MDNS.begin("esp8266")) {
    Serial.println("MDNS responder started");
  }
  // 当WiFi连接成功后，开启web服务
  server.on("/", handleRoot);
  server.on("/inline", []() {
    server.send(200, "text/plain", "this works as well");
  });
  server.onNotFound(handleNotFound);
  server.begin();
  Serial.println("HTTP server started");
  pinMode(MIN1, OUTPUT);
  pinMode(MIN2, OUTPUT);
  pinMode(CIN1, OUTPUT);
  pinMode(CIN2, OUTPUT);

  thread_send.onRun(niceCallback);
  thread_send.setInterval(10000);
  thread_receive.onRun(niceCallback2);
  thread_receive.setInterval(100);
}
void direction(String t)
{
  if (t == "上")
  {
    digitalWrite(MIN1, HIGH);
    digitalWrite(MIN2, LOW);
    digitalWrite(CIN1, LOW);
    digitalWrite(CIN2, HIGH);

  } else if (t == "下")
  {
    digitalWrite(MIN1, LOW);
    digitalWrite(MIN2, HIGH);
    digitalWrite(CIN1, HIGH);
    digitalWrite(CIN2, LOW);
  } else if (t == "左")
  {
    digitalWrite(MIN1, LOW);
    digitalWrite(MIN2, LOW);
    digitalWrite(CIN1, LOW);
    digitalWrite(CIN2, HIGH);
  } else if (t == "右")
  {
    digitalWrite(MIN1, HIGH);
    digitalWrite(MIN2, LOW);
    digitalWrite(CIN1, LOW);
    digitalWrite(CIN2, LOW);
  }
}

void Mstop()
{
  digitalWrite(MIN1, LOW);
  digitalWrite(MIN2, LOW);
  digitalWrite(CIN1, LOW);
  digitalWrite(CIN2, LOW);
}

void loop(void) {
  server.handleClient();
  if (thread_send.shouldRun())
    thread_send.run();
  if (thread_receive.shouldRun())
    thread_receive.run();
}

void niceCallback() {
  byte temperature = 0;
  byte humidity = 0;
  int err = SimpleDHTErrSuccess;
  if ((err = dht11.read(&temperature, &humidity, NULL)) != SimpleDHTErrSuccess) {
    Serial.print("Read DHT11 failed, err="); Serial.println(err); delay(1000);
    return;
  }
  Serial.print((int)temperature); Serial.print(" *C, ");
  Serial.print((int)humidity); Serial.println(" H");
  // 向服务器发送数据
  http.begin("http://www.gotit.top/api/car/upload_data");
  http.addHeader("Content-Type", "application/json");
  StaticJsonBuffer<64> jsonBuffer;
  JsonObject& root = jsonBuffer.createObject();
  root["car_id"] = 4;
  root["t"] = temperature;
  root["h"] = humidity;
  String output = "";
  root.printTo(output);
  int httpCode = http.POST(output);
  http.end();
}

void niceCallback2() {
  String command = server.arg("direction");
  Serial.println(command);
  if (command == "左") {
    direction(command);
    Serial.println(command);
  } 
  else if (command == "右")
  {
    direction(command);
    Serial.println(command);
  } 
  else if (command == "上")
  {
    direction(command);
    Serial.println(command);
  }
  else if (command == "下")
  {
    direction(command);
    Serial.println(command);
  } 
  else if (command == "中心")
  {
    Mstop();
    Serial.println(command);
  }
}
