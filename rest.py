from flask import Flask, request, jsonify

app = Flask(__name__)
messages = []

@app.route('/messages', methods=['GET'])
def get_messages():
    return jsonify(messages)

@app.route('/messages', methods=['POST'])
def add_message():
    message = request.data.decode('utf-8')
    messages.append(message)
    return '', 204

if __name__ == '__main__':
    app.run(debug=True)
