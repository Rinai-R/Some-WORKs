CONSUMER_DIR = /home/rinai/PROJECTS/Some-WORKs/2025/03March/20250319/rocketmq/consu
PRODUCER_DIR = /home/rinai/PROJECTS/Some-WORKs/2025/03March/20250319/rocketmq/prod

start:
	cd $(PRODUCER_DIR) && go run main.go
	cd $(CONSUMER_DIR) && go run main.go

stop:
	cd $(CONSUMER_DIR) && kill main.go
	cd $(PRODUCER_DIR) && kill main.go