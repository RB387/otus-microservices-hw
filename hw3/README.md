## Запуск приложения:
```
helm install users ./manifests/users {-n namespace}
```

Для метрик ingress в minikube (опционально, если используется ingress как addon):
* `minikube addons disable ingress`
* Поставить в values.yaml `ingress.customIngress = true`
* `helm install users ./manifests/users {-n namespace}`


Настроенные дашборды:
* service-user: Метрики сервиса и NGINX Ingress, подов приложения
* PostgreSQL Database: Метрики базы
* NGINX Ingress controller: Подробные метрики NGINX

Benchmark проводился в 200 коннектов и максимум 2000 запросов в секунду
```
fasthttploader -c 200 -q 2000 -t 10s -d 300s http://arch.homework/service-user/user/someusername
```
Результаты в файлах:
* benchmark.png - Метрики Postgres из Prometheus
* benchmark_db.png - Метрики базы из Prometheus
* benchmark_report.html - Сгенерированный отчет

## Задание:
В этом ДЗ вы научитесь инструментировать сервис.  
Инструментировать сервис из прошлого задания метриками в формате Prometheus с помощью библиотеки для вашего фреймворка и ЯП.  
Сделать дашборд в Графане, в котором были бы метрики с разбивкой по API методам:  
* Latency (response time) с квантилями по 0.5, 0.95, 0.99, max  
* RPS
* Error Rate - количество 500ых ответов  

Добавить в дашборд графики с метрикам в целом по сервису, взятые с nginx-ingress-controller:  
* Latency (response time) с квантилями по 0.5, 0.95, 0.99, max
* RPS
* Error Rate - количество 500ых ответов
* Настроить алертинг в графане на Error Rate и Latency.

На выходе должно быть:
* скриншоты дашборды с графиками в момент стресс-тестирования сервиса. Например, после 5-10 минут нагрузки.
* json-дашборды.
   
Задание со звездочкой (+5 баллов)
 * Используя существующие системные метрики из кубернетеса, добавить на дашборд графики с метриками:
 * Потребление подами приложения памяти
 * Потребление подами приолжения CPU
 * Инструментировать базу данных с помощью экспортера для prometheus для этой БД.
 * Добавить в общий дашборд графики с метриками работы БД. 