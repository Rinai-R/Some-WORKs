PRODUCER_DIR = /home/rinai/PROJECTS/Some-WORKs/2025/03March/20250318/nsq/prod
CONSUMER_DIR = /home/rinai/PROJECTS/Some-WORKs/2025/03March/20250318/nsq/consu

start:
	cd $(PRODUCER_DIR) && go run main.go &
	cd $(CONSUMER_DIR) && go run main.go &

stop:
	cd $(PRODUCER_DIR) && kill main.go &
	cd $(CONSUMER_DIR) && kill main.go &

