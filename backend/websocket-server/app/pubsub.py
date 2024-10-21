import asyncio
import logging
from typing import Callable

import redis.asyncio
from redis.asyncio import from_url
from tenacity import retry, retry_if_exception_type, wait_fixed

RETRY_INTERVALL = 2


async def redis_reader(channel: redis.asyncio.client.PubSub, callback: Callable[[str], None]):
    while True:
        message = await channel.get_message(ignore_subscribe_messages=True)

        if message and message["data"]:
            data = message["data"].decode("utf-8")
            logging.info(f"Received data from Redis: {data}")
            callback(data)


@retry(wait=wait_fixed(RETRY_INTERVALL), retry=retry_if_exception_type(redis.exceptions.ConnectionError))
async def start_redis_subscription(url: str, callback: Callable[[str], None]):
    client = from_url(url)
    async with client.pubsub() as pubsub:
        try:
            await pubsub.psubscribe("events.*")
            await asyncio.create_task(redis_reader(pubsub, callback))
        except redis.exceptions.ConnectionError as e:
            logging.error(f"Connection to Redis failed, trying again in {RETRY_INTERVALL} seconds.")
            raise e
