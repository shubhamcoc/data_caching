import requests
import time
import json
from flask import Flask, request, Response
from werkzeug import serving

APP = Flask(__name__)
@APP.route("/submit", methods=['GET', 'POST'])
def send_submit():
    url = "http://localhost:10000/api/submit"
    id = 123400
    num = 5
    for i in range(0, 99):
        data = {"key":str(id), "value": "test"+str(num)}
        response = requests.post(url, data=json.dumps(data))
        time.sleep(1)
        id = id+1
        num = num+1

def main():
    APP.run('0.0.0.0', 8082)

if __name__ == "__main__":
    main()
