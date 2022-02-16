#!/usr/bin/python3
import hmac
import hashlib
import time
import requests
import base64
from datetime import datetime, timezone
from typing import Callable
import json
import os


api = "https://rest.websupport.sk"
apiKey = "YOUR-WEBSUPPORT-API-KEY"
secret = "YOUR-WEBSUPPORT-SECRET"


def make_request(method: Callable, path: str, data: dict={}) -> any:
    timestamp = int(time.time())
    canonicalRequest = f"{method.__name__.upper()} {path} {timestamp}"
    signature = hmac.new(bytes(secret, 'UTF-8'), bytes(canonicalRequest, 'UTF-8'), hashlib.sha1).hexdigest()

    headers = {
        "Content-Type": "application/json",
        "Accept": "application/json",
        "Date": datetime.fromtimestamp(timestamp, timezone.utc).isoformat()
    }

    return method(f"{api}{path}", headers=headers, auth=(apiKey, signature), data=json.dumps(data))


def main():
    records = make_request(requests.get, "/v1/user/self/zone/streamer272.com/record")
    if records.status_code != 200 or not records.ok:
        raise Exception(f"Something went wrong, {records.content.decode()}")
    else:
        records = json.loads(records.content.decode())

    ip = os.getenv("IP")

    for record in records["items"]:
        if "requireipchange" not in record["note"]:
            continue

        print(f"Updating record {record['id']}, setting to {ip}")
        request = make_request(requests.put, f"/v1/user/self/zone/YOUR-DOMAIN/record/{record['id']}", {
            "content": ip,
        })
        print(f"Response: {request.content.decode()}")


if __name__ == "__main__":
    main()

