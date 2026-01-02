import pika
import json
import time
from flask import Flask, request, jsonify

app = Flask(__name__)

def get_rabbitmq_connection():
    # Retry connection because RabbitMQ takes a few seconds to start
    retries = 5
    while retries > 0:
        try:
            connection = pika.BlockingConnection(
                pika.ConnectionParameters(host='rabbitmq'))
            return connection
        except Exception:
            retries -= 1
            time.sleep(5)
    return None

@app.route('/send', methods=['POST'])
def send_task():
    data = request.json
    message = data.get('message', 'Default Task')

    connection = get_rabbitmq_connection()
    if not connection:
        return jsonify({"error": "Could not connect to RabbitMQ"}), 500

    channel = connection.channel()
    # Ensure the queue exists
    channel.queue_declare(queue='task_queue', durable=True)

    # Publish message to the queue
    channel.basic_publish(
        exchange='',
        routing_key='task_queue',
        body=json.dumps({"task": message, "timestamp": time.time()}),
        properties=pika.BasicProperties(delivery_mode=2) # Make message persistent
    )

    connection.close()
    return jsonify({"status": "Task sent to queue", "task": message})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)