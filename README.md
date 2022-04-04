# :heavy_check_mark: wb-L0
**MailingService** - сервис получения модели, используя nats-streaming, добавления модели в бд, а также просмотра
информации модели по запросу

# Конфиг
Конфиг находиться по умолчанию configs/

![image](https://user-images.githubusercontent.com/80648065/161534453-da3fb8d4-4172-48f6-abb8-3c3b23a9d241.png)
# Флаги
![image](https://user-images.githubusercontent.com/80648065/161534720-45f962d7-3f3b-4026-ae58-10af32f7868f.png)
# Схема базы данных
![image](https://user-images.githubusercontent.com/80648065/161535043-5321106b-9bf8-49ad-bf32-72c4b1bed38c.png)
# Usage
По умолчанию собираются и запускаются все контейнеры

    make

Для того чтобы опубликовать в канал сообщение, файл должен находиться nats-streaming-publish/

    make pub FILE="filename.json"

    
## Other
**Author:**  
:vampire:*[Deeds Baron](https://github.com/DeedsBaron)*  
