import pika
import json
import time
from flask import Flask, request, jsonify
from flask_cors import CORS # Importa CORS

app = Flask(__name__)
CORS(app) # Abilita CORS per tutte le rotte

def get_rabbitmq_connection():
    retries = 10
    while retries > 0:
        try:
            # Connect to RabbitMQ container
            connection = pika.BlockingConnection(
                pika.ConnectionParameters(host='rabbitmq'))
            return connection
        except Exception:
            retries -= 1
            print(f"Connection failed, retries left: {retries}")
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
    channel.queue_declare(queue='task_queue', durable=True)

    channel.basic_publish(
        exchange='',
        routing_key='task_queue',
        body=json.dumps({"task": message}),
        properties=pika.BasicProperties(delivery_mode=2)
    )

    connection.close()
    return jsonify({"status": "Task sent", "message": message})

if __name__ == '__main__':
    # IMPORTANTE: host='0.0.0.0' per accettare connessioni esterne al container
    app.run(host='0.0.0.0', port=5000)